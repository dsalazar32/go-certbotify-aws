package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/dsalazar32/go-gen-ssl/command/certbot"
	"os"
	"path/filepath"
	"strings"
)

// This just acts as a proxy to the certbot project.
// The resulting certificates will then be handed off to
// logic that will interact with AWS iam certificate manager.
// TODO: Have post processing of generated cert be flag driven.
// TODO: Nice to have would be a DNS host name validator
type SSLGenerator struct {
	Certbot certbot.Certbot

	Meta
}

func (s *SSLGenerator) Help() string {
	return "implement me"
}

func (s *SSLGenerator) Synopsis() string {
	return `This tool just acts as a proxy to the certbot project. The resulting artifacts (certificates) will be used to update AWS Certificate Manager.`
}

func (s *SSLGenerator) Run(args []string) int {

	var (
		domainsFlag certbot.Domains
		emailFlag   string
		s3Flag      bool
	)

	f := s.Meta.flagSet("SSLGenerator")
	f.Var(&domainsFlag, "d", "Comma-separated list of domains to obtain a certificate for")
	f.StringVar(&emailFlag, "email", "", "Email address for important account notifications")
	f.BoolVar(&s3Flag, "s3", false, "Target S3 bucket to upload generated certificates to")
	if err := f.Parse(args); err != nil {
		return 1
	}

	// These will be the default flags that will be proxied to the certbot cli.
	s.Certbot.SetCertbotFlag("--email", emailFlag)
	s.Certbot.SetCertbotFlag("-d", domainsFlag)

	if !s.Certbot.Test {
		if err := s.Certbot.GenerateCertificate(); err != nil {
			s.Ui.Error(fmt.Sprintf("execution error: %s, %v", s.Certbot.CommandString(), err))
			return 1
		}
	} else {
		s.Ui.Info(s.Certbot.CommandString())
	}

	// TODO: If any of these flags are set s3, ssl-manager
	// TODO: Create bucket for certificates to land in `accountno_certbot_certificates/domain/date/`
	if !s.Certbot.Test && s3Flag {
		sess := session.Must(session.NewSession())

		// Get callers aws account number to use for unique naming of resources. In this case
		// we'll assume that a bucket prefixed with the caller's account number will is unique.
		stsSvc := sts.New(sess)
		stsInput := &sts.GetCallerIdentityInput{}
		stsOut, err := stsSvc.GetCallerIdentity(stsInput)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					s.Ui.Error(aerr.Error())
				}
			} else {
				s.Ui.Error(err.Error())
			}
			return 1
		}
		awsAccntNo := *stsOut.Account

		s3BucketPttrn := "certbot-certificates-%s"
		s3Bucket := fmt.Sprintf(s3BucketPttrn, awsAccntNo)
		if s3svc, err := findOrCreateS3Bucket(sess, s3Bucket); err != nil {
			s.Ui.Error(err.Error())
		} else {
			if s.uploadCertificatesToS3(s3svc, s3Bucket); err != nil {
				s.Ui.Error(err.Error())
				return 1
			}
		}
	}

	return 0
}

func (s *SSLGenerator) uploadCertificatesToS3(s3svc *s3.S3, bucket string) error {
	p := certbot.OutfilePath
	uploader := s3manager.NewUploaderWithClient(s3svc)
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			upInput := &s3manager.UploadInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(strings.TrimPrefix(path, p)),
				Body:   f,
			}
			result, err := uploader.Upload(upInput, func(uploader *s3manager.Uploader) {})
			if err != nil {
				return err
			}
			s.Ui.Info(fmt.Sprintf("file uploaded to, %s", result.Location))
			f.Close()
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// AWS Helpers
func findOrCreateS3Bucket(sess *session.Session, bucket string) (*s3.S3, error) {
	s3svc := s3.New(sess)
	s3Input := &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}

	_, err := s3svc.HeadBucket(s3Input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			// Some shenanigans is afoot with the error codes returned from this sdk call.
			// Doesn't seem to map to any of the Error constants defined and just sticks to
			// the response type of the http request.
			case s3.ErrCodeNoSuchBucket, "NotFound":
				return createS3Bucket(s3svc, bucket)
			default:
				return nil, aerr
			}
		}
	}

	return s3svc, nil
}

func createS3Bucket(s3svc *s3.S3, bucket string) (*s3.S3, error) {
	s3Input := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}

	_, err := s3svc.CreateBucket(s3Input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				return nil, aerr
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				return s3svc, nil
			default:
				return nil, aerr
			}
		} else {
			return nil, err
		}
	}
	return s3svc, nil
}

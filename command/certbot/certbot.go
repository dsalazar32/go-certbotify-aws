package certbot

import (
	"fmt"
	"github.com/dsalazar32/go-gen-ssl/command/utils"
	"os"
	"strings"
)

const (
	CertbotCmdPrefix = "certbot certonly"
	// OutfilePath      = "/etc/letsencrypt/archive"
)

type Certbot struct {
	CertbotFlags CertbotFlags
	Test         bool
}

type certbotFlag struct {
	Flag string
	Val  string
}

type CertbotFlags []certbotFlag

func (cf CertbotFlags) String() string {
	var f []string
	var s string
	for _, i := range cf {
		s = i.Flag
		if i.Val != "" {
			s = fmt.Sprintf("%s %s", s, i.Val)
		}
		f = append(f, s)
	}
	return strings.Join(f, " ")
}

type Domains []string

func (d *Domains) Set(v string) error {
	*d = append(*d, v)
	return nil
}

func (d *Domains) String() string {
	return fmt.Sprint(*d)
}

func (c *Certbot) CommandString() string {
	return fmt.Sprintf("%s %s", CertbotCmdPrefix, c.CertbotFlags.String())
}

func (c *Certbot) GenerateCertificate() error {
	var awsCred []string
	awsEnvs := []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"}
	for _, e := range awsEnvs {
		v, ok := os.LookupEnv(e)
		if ok {
			awsCred = append(awsCred, e, v)
		}
	}

	cmd := utils.Commander(".", awsCred...)
	if !c.Test {
		_, err := cmd(fmt.Sprintf("type %s", strings.Split(CertbotCmdPrefix, " ")[0]), true)
		if err != nil {
			return err
		}

		_, err = cmd(c.CommandString(), true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Certbot) SetCertbotFlag(k string, v interface{}) {
	var cf certbotFlag
	switch v.(type) {
	case string:
		vt := v.(string)
		if vt == "" {
			return
		}
		cf = certbotFlag{k, vt}
	case bool:
		vt := v.(bool)
		if !vt {
			return
		}
		cf = certbotFlag{k, ""}
	case Domains:
		vt := v.(Domains)
		if len(vt) == 0 {
			return
		}
		var d []string
		for _, f := range vt {
			d = append(d, fmt.Sprintf("%s %s", k, f))
		}
		cf = certbotFlag{strings.Join(d, " "), ""}
	}
	c.CertbotFlags = append(c.CertbotFlags, cf)
}

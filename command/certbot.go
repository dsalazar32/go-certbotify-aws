package command

import (
	"fmt"
	"github.com/dsalazar32/go-certbotify-aws/command/utils"
	"strings"
)

// This just acts as a proxy to the certbot project.
// The resulting certificates will then be handed off to
// logic that will interact with AWS iam certificate manager.
// TODO: Build go package in docker container
// TODO: Setup proxy logic
// TODO: Upload resulting cert to s3 for backup (Optional)
// TODO: Nice to have would be a DNS host name validator
var certbotCommand = []string{"certbot", "certonly"}

type CertbotCommand struct {
	CertbotFlags certbotFlags

	Meta
}

// The list of commands and flags that are proxied to the certbot
// command is limited to those that support the route53 scenario.
type certbotFlags []certbotFlag
type certbotFlag struct {
	Flag string
	Val  string
}

func (cf certbotFlags) String() string {
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

type domains []string

func (d *domains) String() string {
	return fmt.Sprint(*d)
}

func (d *domains) Set(v string) error {
	*d = append(*d, v)
	return nil
}

func (c *CertbotCommand) Help() string {
	return "implement me"
}

func (c *CertbotCommand) Run(args []string) int {
	cmd := utils.Commander(".")

	// Check if Certbot is even installed before proceeding.
	out, err := cmd(fmt.Sprintf("type %s", certbotCommand[0]), true)
	if err != nil {
		c.Ui.Error("Certbot is not installed")
		return 1
	}
	c.Ui.Info(fmt.Sprintf("%s", out))

	var (
		domainsFlag domains
		emailFlag   string
		nFlag       bool
		tosFlag     bool
		r53Flag     bool
	)

	f := c.Meta.flagSet("CertbotCommand")
	f.Var(&domainsFlag, "d", "Comma-separated list of domains to obtain a certificate for")
	f.StringVar(&emailFlag, "email", "", "Email address for important account notifications")
	f.BoolVar(&nFlag, "n", false, "Run non-interactively")
	f.BoolVar(&tosFlag, "agree-tos", false, "Agree to the ACME server's Subscriber Agreement")
	f.BoolVar(&r53Flag, "dns-route53", false, "Use route53 for the challenge")
	if err := f.Parse(args); err != nil {
		return 1
	}

	// These will be the default flags that will be proxied to the certbot cli.
	c.setCertbotFlag("-n", nFlag)
	c.setCertbotFlag("--dns-route53", r53Flag)
	c.setCertbotFlag("--agree-tos", tosFlag)
	c.setCertbotFlag("--email", emailFlag)
	c.setCertbotFlag("-d", domainsFlag)

	return 0
}

func (c *CertbotCommand) Synopsis() string {
	return `This tool just acts as a proxy to the certbot project. The resulting artifacts (certificates) will be used to update AWS Certificate Manager.`
}

func (c *CertbotCommand) CommandString() string {
	return fmt.Sprintf("%s %s", strings.Join(certbotCommand, " "), c.CertbotFlags.String())
}

func (c *CertbotCommand) setCertbotFlag(k string, v interface{}) {
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
	case domains:
		vt := v.(domains)
		if len(vt) == 0 {
			return
		}
		var c []string
		for _, f := range vt {
			c = append(c, fmt.Sprintf("%s %s", k, f))
		}
		cf = certbotFlag{strings.Join(c, " "), ""}
	}
	c.CertbotFlags = append(c.CertbotFlags, cf)
}

package command

import (
	"fmt"
	"strings"
)

// This just acts as a proxy to the certbot project.
// The resulting certificates will then be handed off to
// logic that will interact with AWS iam certificate manager.
// TODO: Write test
// TODO: Detect certbot installed
// TODO: Setup proxy logic
// TODO: Upload resulting cert to s3 for backup (Optional)
// TODO: Nice to have would be a DNS host name validator
type CertbotCommand struct {
	CertbotFlags []certbotFlag

	Meta
}

// The list of commands and flags that are proxied to the certbot
// command is limited to those that support the route53 scenario.
type certbotFlag struct {
	Flag string
	Expl string
}

type domains []string

// var certbotCommand = "certonly"
var domainsFlag domains
var emailFlag string

func (c *CertbotCommand) Help() string {
	return "implement me"
}

func (d *domains) String() string {
	return fmt.Sprint(*d)
}

func (d *domains) Set(v string) error {
	*d = append(*d, v)
	return nil
}

func (c *CertbotCommand) Run(args []string) int {
	f := c.Meta.flagSet("CertbotCommand")
	f.Var(&domainsFlag, "d", "Comma-separated list of domains to obtain a certificate for")
	f.StringVar(&emailFlag, "email", "", "Email address for important account notifications")
	if err := f.Parse(args); err != nil {
		return 1
	}

	// These will be the default flags that will be proxied to the certbot cli.
	c.setCertbotFlag("-n", "Run non-interactively")
	c.setCertbotFlag("--agree-tos", "Agree to the ACME server's Subscriber Agreement")
	c.setCertbotFlag("--dns-route53", "Use route53 for the challenge")
	c.setCertbotFlag("-email", emailFlag)
	c.setCertbotFlag("-d", domainsFlag)

	fmt.Printf("%v", c.CertbotFlags)
	return 0
}

func (c *CertbotCommand) Synopsis() string {
	return `This tool just acts as a proxy to the certbot project. The resulting artifacts (certificates) will be used to update AWS Certificate Manager.`
}

func (c *CertbotCommand) CommandString(f []certbotFlag) string {
	return ""
}

func (c *CertbotCommand) setCertbotFlag(k string, v interface{}) {
	var cf certbotFlag
	switch v.(type) {
	case string:
		vt := v.(string)
		if vt == "" {
			return
		}
		cf = certbotFlag{fmt.Sprintf("%s %s", k, vt), ""}
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

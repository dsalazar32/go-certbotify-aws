package command

import (
	"fmt"
	"github.com/dsalazar32/go-gen-ssl/command/certbot"
)

// This just acts as a proxy to the certbot project.
// The resulting certificates will then be handed off to
// logic that will interact with AWS iam certificate manager.
// TODO: See if it's possible to get io.Writer out of UI
// TODO: Upload resulting cert to s3 for backup (Optional)
// TODO: Nice to have would be a DNS host name validator
type CertbotCommand struct {
	Certbot certbot.Certbot

	Meta
}

func (c *CertbotCommand) Help() string {
	return "implement me"
}

func (c *CertbotCommand) Run(args []string) int {

	var (
		domainsFlag certbot.Domains
		emailFlag   string
	)

	f := c.Meta.flagSet("CertbotCommand")
	f.Var(&domainsFlag, "d", "Comma-separated list of domains to obtain a certificate for")
	f.StringVar(&emailFlag, "email", "", "Email address for important account notifications")
	if err := f.Parse(args); err != nil {
		return 1
	}

	// These will be the default flags that will be proxied to the certbot cli.
	c.Certbot.SetCertbotFlag("--email", emailFlag)
	c.Certbot.SetCertbotFlag("-d", domainsFlag)

	if !c.Certbot.Test {
		if err := c.Certbot.GenerateCertificate(); err != nil {
			c.Ui.Error(fmt.Sprintf("execution error: %s, %v", c.Certbot.CommandString(), err))
			return 1
		}
	} else {
		c.Ui.Info(c.Certbot.CommandString())
	}

	return 0
}

func (c *CertbotCommand) Synopsis() string {
	return `This tool just acts as a proxy to the certbot project. The resulting artifacts (certificates) will be used to update AWS Certificate Manager.`
}

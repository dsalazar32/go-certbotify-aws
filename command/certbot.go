package command

import "fmt"

// This just acts as a proxy to the certbot project.
// The resulting certificates will then be handed off to
// logic that will interact with AWS iam certificate manager.
// TODO: Detect certbot installed
// TODO: Setup proxy logic
// TODO: Upload resulting cert to s3 for backup (Optional).
// TODO: Nice to have would be a DNS host name validator.
type Certbot struct {
	Meta
}

// The list of commands and flags that are proxied to the certbot
// command is limited to those that support the route53 scenario.
type certbotFlag struct {
	Flag string
	Expl string
}

type domains []string

var domainsFlag domains
var email string
var certbotCommand = "certonly"
var certbotFlags = []certbotFlag{
	{"-n", "Run non-interactively"},
	{"--agree-tos", "Agree to the ACME server's Subscriber Agreement"},
	{"--dns-route53", "Email address for important account information"},
}

func (c *Certbot) Help() string {
	return "implement me"
}

func (d *domains) String() string {
	return fmt.Sprint(*d)
}

func (d *domains) Set(v string) error {
	*d = append(*d, v)
	return nil
}

func (c *Certbot) Run(args []string) int {
	f := c.Meta.flagSet("Certbot")
	f.Var(&domainsFlag, "d", "Comma-separated list of domains to obtain a certificate for")
	f.StringVar(&email, "email", "", "Email address for important account notifications")
	if err := f.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *Certbot) Synopsis() string {
	return `This tool just acts as a proxy to the certbot project. The resulting artifacts (certificates) will be used to update AWS Certificate Manager.`
}

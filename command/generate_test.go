package command

import (
	"github.com/dsalazar32/go-gen-ssl/command/certbot"
	"github.com/mitchellh/cli"
	"testing"
)

type sampleInOut struct {
	In  []string
	Out string
}

func TestCertbotCommand_CommandString(t *testing.T) {
	tests := []sampleInOut{
		{
			[]string{
				"-email",
				"me@dsalazar.io",
				"-d",
				"dsalazar.io",
				"-d",
				"www.dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --server https://acme-v02.api.letsencrypt.org/directory --email me@dsalazar.io -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[]string{
				"-d",
				"dsalazar.io",
				"-d",
				"www.dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --server https://acme-v02.api.letsencrypt.org/directory -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[]string{
				"-email",
				"me@dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --server https://acme-v02.api.letsencrypt.org/directory --email me@dsalazar.io",
		},
		{
			[]string{
				"-email",
				"me@dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --server https://acme-v02.api.letsencrypt.org/directory --email me@dsalazar.io",
		},
	}

	for _, test := range tests {
		ui := &cli.MockUi{}
		c := &SSLGenerator{
			Certbot: *newCertbotClient(),
			Meta:    Meta{Ui: ui},
		}
		if code := c.Run(test.In); code != 0 {
			t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
		}

		expect, got := test.Out, c.Certbot.CommandString()
		if expect != got {
			t.Fatalf("assertion failed\n expected: %s\n got: %s\n", expect, got)
		}
	}
}

func newCertbotClient() *certbot.Certbot {
	return &certbot.Certbot{
		CertbotFlags: certbot.CertbotFlags{
			{"-n", ""},
			{"--dns-route53", ""},
			{"--agree-tos", ""},
		},
		Test: true,
	}
}

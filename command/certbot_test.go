package command

import (
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
				"-n",
				"-dns-route53",
				"-agree-tos",
				"-email",
				"me@dsalazar.io",
				"-d",
				"dsalazar.io",
				"-d",
				"www.dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --email me@dsalazar.io -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[]string{
				"-n",
				"-dns-route53",
				"-agree-tos",
				"-d",
				"dsalazar.io",
				"-d",
				"www.dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[]string{
				"-n",
				"-dns-route53",
				"-agree-tos",
				"-email",
				"me@dsalazar.io",
			},
			"certbot certonly -n --dns-route53 --agree-tos --email me@dsalazar.io",
		},
		{
			[]string{
				"-dns-route53",
				"-agree-tos",
				"-email",
				"me@dsalazar.io",
			},
			"certbot certonly --dns-route53 --agree-tos --email me@dsalazar.io",
		},
	}

	for _, test := range tests {
		ui := &cli.MockUi{}
		c := &CertbotCommand{
			Meta: Meta{
				Ui: ui,
			},
		}
		if code := c.Run(test.In); code != 0 {
			t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
		}

		expect, got := test.Out, c.CommandString()
		if expect != got {
			t.Fatalf("assertion failed\n expected: %s\n got: %s\n", expect, got)
		}
	}
}

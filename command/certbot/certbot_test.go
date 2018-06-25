package certbot

import (
	"testing"
)

type sampleInOut struct {
	In  [][]string
	Out string
}

func TestCertbot_CommandString(t *testing.T) {
	tests := []sampleInOut{
		{
			[][]string{
				{"--email", "me@dsalazar.io"},
				{"-d", "dsalazar.io"},
				{"-d", "www.dsalazar.io"},
			},
			"certbot certonly -n --dns-route53 --agree-tos --email me@dsalazar.io -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[][]string{
				{"-d", "dsalazar.io"},
				{"-d", "www.dsalazar.io"},
			},
			"certbot certonly -n --dns-route53 --agree-tos -d dsalazar.io -d www.dsalazar.io",
		},
		{
			[][]string{
				{"--email", "me@dsalazar.io"},
			},
			"certbot certonly -n --dns-route53 --agree-tos --email me@dsalazar.io",
		},
		{
			[][]string{
				{"--email", "me@dsalazar.io"},
			},
			"certbot certonly -n --dns-route53 --agree-tos --email me@dsalazar.io",
		},
	}

	for _, test := range tests {
		c := newCertbotClient()
		for _, f := range test.In {
			c.SetCertbotFlag(f[0], f[1])
		}

		expect, got := test.Out, c.CommandString()
		if expect != got {
			t.Fatalf("assertion failed\n expected: %s\n got: %s\n", expect, got)
		}
	}
}

func newCertbotClient() *Certbot {
	return &Certbot{
		CertbotFlags: CertbotFlags{
			{"-n", ""},
			{"--dns-route53", ""},
			{"--agree-tos", ""},
		},
		Test: true,
	}
}

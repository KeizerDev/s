package gmail

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("gmail", &GmailProvider{})
}

// GmailProvider adheres to the Provider interface.
type GmailProvider struct {
}

// BuildURI generates a search URL for gmail.
func (p *GmailProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://mail.google.com/mail/u/0/#search/%s", url.QueryEscape(q))
}

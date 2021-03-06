package amazon

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("amazon", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Amazon.
func (p *Provider) BuildURI(q string) string {
	switch providers.Region() {
	case "GB":
		return fmt.Sprintf("https://www.amazon.co.uk/s/?keywords=%s", url.QueryEscape(q))
	default:
		return fmt.Sprintf("https://www.amazon.com/s/?keywords=%s", url.QueryEscape(q))
	}
}

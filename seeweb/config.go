package seeweb

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/uwtrilogyseaward0m/go-seeweb/seeweb"
)

// Config defines the configuration options for the Seeweb client
type Config struct {
	mu sync.Mutex

	// The Seeweb API URL
	ApiUrl string

	// Override default Seeweb API URL
	ApiUrlOverride string

	// The Seeweb API token
	Token string

	// UserAgent for API Client
	UserAgent string

	client *seeweb.Client
}

const invalidCreds = `

No valid credentials found for Seeweb provider.
Please see https://www.terraform.io/docs/providers/seeweb/index.html
for more information on providing credentials for this provider.
`

// Client returns a Seeweb client, initializing when necessary.
func (c *Config) Client() (*seeweb.Client, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Return the previously-configured client if available.
	if c.client != nil {
		return c.client, nil
	}

	// Validate that the Seeweb token is set
	if c.Token == "" {
		return nil, fmt.Errorf(invalidCreds)
	}

	var httpClient *http.Client
	httpClient = http.DefaultClient
	httpClient.Transport = logging.NewTransport("Seeweb", http.DefaultTransport)

	var apiUrl = c.ApiUrl
	if c.ApiUrlOverride != "" {
		apiUrl = c.ApiUrlOverride
	}

	config := &seeweb.Config{
		BaseURL:    apiUrl,
		Debug:      logging.IsDebugOrHigher(),
		HTTPClient: httpClient,
		Token:      c.Token,
		UserAgent:  c.UserAgent,
	}

	client, err := seeweb.NewClient(config)
	if err != nil {
		return nil, err
	}

	c.client = client

	log.Printf("[INFO] Seeweb client configured")

	return c.client, nil
}

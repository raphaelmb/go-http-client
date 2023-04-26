package client

import "net/http"

var (
	DefautlAPIURL = "https://pokeapi.co"
)

type Client struct {
	apiURL     string
	httpClient *http.Client
}

type Option func(c *Client)

func NewClient(opts ...Option) *Client {
	client := &Client{
		apiURL:     DefautlAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, o := range opts {
		o(client)
	}

	return client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

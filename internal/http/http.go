package http

import (
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	config     *nacos.Config
	webClient  *resty.Client
	Middleware bool
}

func NewClient(config *nacos.Config) *Client {
	client := Client{config: config, webClient: resty.New().SetDebug(true)}
	return &client
}

func (client *Client) Get[T any](url string, params map[string]string) (*T, error) {
	if client.Middleware {
		applyMiddleware(client)
	}
	var result T
	req := client.webClient.R().
		SetResult(&result)

	if params != nil {
		req.SetQueryParams(params)
	}

	_, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func applyMiddleware(client *Client) {
	client.webClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		if client.config.Nacos.Auth {
			req.SetQueryParam("accessToken", client.config.Nacos.Token)
		}
		return nil
	})
	client.webClient.OnAfterResponse(func(c *resty.Client, res *resty.Response) error {
		return nil
	})
}

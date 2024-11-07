package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal/types"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	config    *types.Config
	webClient *resty.Client
}

func NewClient(config *types.Config) *Client {
	var webClient = resty.New().SetDebug(true)
	webClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		if config.Nacos.Auth {
			req.SetQueryParam("accessKey", config.Nacos.Token)
		}
		fmt.Println("BeforeRequest", req.URL)
		return nil
	})
	webClient.OnAfterResponse(func(c *resty.Client, res *resty.Response) error {
		fmt.Println("AfterResponse", string(res.Body()))
		var result map[string]any
		err := json.Unmarshal(res.Body(), &result)
		if err != nil {
			return errors.New(string(res.Body()))
		}
		if fmt.Sprintf("%v", result["code"]) != "0" {
			return errors.New(result["data"].(string))
		}
		return nil
	})
	client := &Client{
		config:    config,
		webClient: webClient,
	}
	return client
}

func Get[T any](client *Client, url string, params map[string]string) (*T, error) {
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

func Post[T any](client *Client, url string, params map[string]string) (*T, error) {
	var result T
	req := client.webClient.R().
		SetResult(&result)

	if params != nil {
		req.SetFormData(params)
	}

	_, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Put[T any](client *Client, url string, params map[string]string) (*T, error) {
	var result T
	req := client.webClient.R().
		SetResult(&result)

	if params != nil {
		req.SetFormData(params)
	}

	_, err := req.Put(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Delete[T any](client *Client, url string, params map[string]string) (*T, error) {
	var result T
	req := client.webClient.R().
		SetResult(&result)

	if params != nil {
		req.SetFormData(params)
	}

	_, err := req.Delete(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

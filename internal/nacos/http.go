package nacos

import (
	"encoding/json"
	"errors"
	"fmt"
	nurl "net/url"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/setting"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	config    *types.Config
	webClient *resty.Client
	owner     *Client
}

func NewHttpClient(config *types.Config, owner *Client) *HttpClient {
	var webClient = resty.New()
	client := &HttpClient{
		config:    config,
		webClient: webClient,
		owner:     owner,
	}
	webClient.OnBeforeRequest(
		func(c *resty.Client, req *resty.Request) error {
			if config.Nacos.Token != "" && !IsNoAuthApi(req.URL) {
				req.SetQueryParam("accessToken", config.Nacos.Token)
			}
			internal.VerboseLogReq(req)
			return nil
		},
	)
	webClient.OnAfterResponse(
		func(c *resty.Client, res *resty.Response) error {
			internal.VerboseLogRes(res)
			url := res.Request.URL

			if res.StatusCode() == 200 {
				var result map[string]any
				err := json.Unmarshal(res.Body(), &result)
				if err != nil {
					return errors.New("json unmarshal error : " + string(res.Body()))
				}
				if value, exists := result["code"]; exists {
					codeStr := fmt.Sprintf("%v", value)
					if codeStr != "0" && codeStr != "200" {
						return errors.New(result["data"].(string))
					}
				}
				return nil
			} else if res.StatusCode() == 403 && !IsLoginApi(url) {
				loginResp, err := client.owner.Login(
					config.Nacos.Addr, config.Nacos.Username, config.Nacos.Password,
				)
				internal.CheckErr(err)
				config.Nacos.Token = loginResp.AccessToken
				err = setting.SaveSetting(loginResp.AccessToken)
				internal.CheckErr(err)
				parse, err := nurl.Parse(url)
				internal.CheckErr(err)
				reUrl := fmt.Sprintf("%s://%s%s", parse.Scheme, parse.Host, parse.Path)
				internal.VerboseLog("re-url: %s", reUrl)
				res.Request.SetCookies(nil)
				_, _ = res.Request.Execute(res.Request.Method, reUrl)
				return nil
			} else {
				return fmt.Errorf("%s", res.Body())
			}
		},
	)
	return client
}

func Get[T any](client *HttpClient, url string, params map[string]string) (*T, error) {
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

func Post[T any](client *HttpClient, url string, params map[string]string) (*T, error) {
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

func Put[T any](client *HttpClient, url string, params map[string]string) (*T, error) {
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

func Delete[T any](client *HttpClient, url string, params map[string]string) (*T, error) {
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

func DeleteByQuery[T any](client *HttpClient, url string, params map[string]string) (*T, error) {
	var result T
	req := client.webClient.R().
		SetResult(&result)

	if params != nil {
		req.SetQueryParams(params)
	}

	_, err := req.Delete(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

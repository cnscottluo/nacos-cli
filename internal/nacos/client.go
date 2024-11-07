package nacos

import (
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/http"
)

// Client Nacos client
type Client struct {
	config *Config
	http   *http.Client
}

func NewClient(config *Config) *Client {
	client := Client{config: config, http: http.NewClient(config, true)}
	return &client
}

func (client *Client) GetNamespaces() ([]NamespaceResp, error) {
	http.ApplyMiddleware = true
	defer func() {
		http.ApplyMiddleware = false
	}()
	var result R[[]NamespaceResp]
	_, err := http.Client.R().
		SetResult(&result).
		Get(client.config.Nacos.Addr + getNamespaceListUrl)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (client *Client) GetNamespace(namespaceId string) (R[NamespaceResp], error) {
	http.ApplyMiddleware = true
	defer func() {
		http.ApplyMiddleware = false
	}()
	resp, _ := http.Client.R().
		SetResult(&R[NamespaceResp]{}).
		SetQueryParam("namespaceId", namespaceId).
		Get(client.config.Nacos.Addr + getNamespaceUrl)
	return *resp.Result().(*R[NamespaceResp]), nil
}

func (client *Client) Login(addr string, username string, password string) (*LoginResponse, error) {
	resp, err := http.Client.R().
		SetResult(&LoginResponse{}).
		SetFormData(internal.Struct2StringMap(LoginRequest{
			Username: username,
			Password: password,
		})).
		Post(addr + loginUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 200 {
		return resp.Result().(*LoginResponse), nil
	} else {
		internal.Log("Login failed: %s", string(resp.Body()))
		return nil, nil
	}
}

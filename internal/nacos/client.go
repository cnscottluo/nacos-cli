package nacos

import (
	"github.com/cnscottluo/nacos-cli/internal/http"
	"github.com/cnscottluo/nacos-cli/internal/types"
)

// Client Nacos client
type Client struct {
	config     *types.Config
	httpClient *http.Client
}

func NewClient(config *types.Config) *Client {
	client := &Client{
		config:     config,
		httpClient: http.NewClient(config),
	}
	return client
}

func (client *Client) GetNamespaces() (*[]NamespaceResp, error) {
	r, err := http.Get[R[[]NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceListUrl, nil)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) GetNamespace(namespaceId string) (*NamespaceResp, error) {
	r, err := http.Get[R[NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceUrl, map[string]string{"namespaceId": namespaceId})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) Login(addr string, username string, password string) (*LoginResp, error) {
	r, err := http.Post[LoginResp](client.httpClient, addr+loginUrl, map[string]string{"username": username, "password": password})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (client *Client) CreateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := http.Post[R[bool]](client.httpClient, client.config.Nacos.Addr+createNamespaceUrl, map[string]string{"namespaceId": namespaceId, "namespaceName": namespaceName, "namespaceDesc": namespaceDesc})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) UpdateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := http.Put[R[bool]](client.httpClient, client.config.Nacos.Addr+updateNamespaceUrl, map[string]string{"namespaceId": namespaceId, "namespaceName": namespaceName, "namespaceDesc": namespaceDesc})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) DeleteNamespace(namespaceId string) (bool, error) {
	r, err := http.Delete[R[bool]](client.httpClient, client.config.Nacos.Addr+deleteNamespaceUrl, map[string]string{"namespaceId": namespaceId})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

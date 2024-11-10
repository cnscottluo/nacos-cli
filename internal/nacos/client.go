package nacos

import (
	"github.com/cnscottluo/nacos-cli/internal/types"
	"strings"
)

// Client Nacos client
type Client struct {
	config     *types.Config
	httpClient *HttpClient
}

// NewClient new client
func NewClient(config *types.Config) *Client {
	client := new(Client)
	client.config = config
	client.httpClient = NewHttpClient(config, client)
	return client
}

// Login login
func (client *Client) Login(addr string, username string, password string) (*LoginResp, error) {
	r, err := Post[LoginResp](client.httpClient, addr+loginUrl, map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetNamespaces get namespaces
func (client *Client) GetNamespaces() (*[]NamespaceResp, error) {
	r, err := Get[R[[]NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceListUrl, nil)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// GetNamespace get namespace
func (client *Client) GetNamespace(namespaceId string) (*NamespaceResp, error) {
	r, err := Get[R[NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceUrl, map[string]string{
		"namespaceId": processNamespace(namespaceId),
	})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// CreateNamespace create namespace
func (client *Client) CreateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Post[R[bool]](client.httpClient, client.config.Nacos.Addr+createNamespaceUrl, map[string]string{
		"namespaceId":   processNamespace(namespaceId),
		"namespaceName": namespaceName,
		"namespaceDesc": namespaceDesc,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// UpdateNamespace update namespace
func (client *Client) UpdateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Put[R[bool]](client.httpClient, client.config.Nacos.Addr+updateNamespaceUrl, map[string]string{
		"namespaceId":   processNamespace(namespaceId),
		"namespaceName": namespaceName,
		"namespaceDesc": namespaceDesc,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// DeleteNamespace delete namespace
func (client *Client) DeleteNamespace(namespaceId string) (bool, error) {
	r, err := Delete[R[bool]](client.httpClient, client.config.Nacos.Addr+deleteNamespaceUrl, map[string]string{
		"namespaceId": processNamespace(namespaceId),
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// GetConfigs get configs
func (client *Client) GetConfigs(namespaceId string) (*[]ConfigResp, error) {
	r, err := Get[R[[]ConfigResp]](client.httpClient, client.config.Nacos.Addr+getConfigListUrl, map[string]string{
		"namespaceId": processNamespace(namespaceId),
	})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// GetConfig get config
func (client *Client) GetConfig(dataId string) (string, error) {
	r, err := Get[R[string]](client.httpClient, client.config.Nacos.Addr+getConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": processNamespace(client.config.Nacos.Namespace),
		"group":       client.config.Nacos.Group,
	})
	if err != nil {
		return "", err
	}
	return *r.Data, nil
}

// DeleteConfig delete config
func (client *Client) DeleteConfig(dataId string) (bool, error) {
	r, err := DeleteByQuery[R[bool]](client.httpClient, client.config.Nacos.Addr+deleteConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": processNamespace(client.config.Nacos.Namespace),
		"group":       client.config.Nacos.Group,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// PublishConfig publish config
func (client *Client) PublishConfig(dataId string, content string, configType string) (bool, error) {
	r, err := Post[R[bool]](client.httpClient, client.config.Nacos.Addr+publishConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": processNamespace(client.config.Nacos.Namespace),
		"group":       client.config.Nacos.Group,
		"content":     content,
		"type":        configType,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// processNamespace process namespace
func processNamespace(namespaceId string) string {
	if strings.ToLower(namespaceId) == "public" {
		return ""
	}
	return namespaceId
}

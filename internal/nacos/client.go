package nacos

import (
	"github.com/cnscottluo/nacos-cli/internal/types"
)

// Client Nacos client
type Client struct {
	config     *types.Config
	httpClient *HttpClient
}

func NewClient(config *types.Config) *Client {
	client := new(Client)
	client.config = config
	client.httpClient = NewHttpClient(config, client)
	return client
}

func (client *Client) Login(addr string, username string, password string) (*LoginResp, error) {
	r, err := Post[LoginResp](client.httpClient, addr+loginUrl, map[string]string{"username": username, "password": password})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (client *Client) GetNamespaces() (*[]NamespaceResp, error) {
	r, err := Get[R[[]NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceListUrl, nil)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) GetNamespace(namespaceId string) (*NamespaceResp, error) {
	r, err := Get[R[NamespaceResp]](client.httpClient, client.config.Nacos.Addr+getNamespaceUrl, map[string]string{"namespaceId": namespaceId})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) CreateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Post[R[bool]](client.httpClient, client.config.Nacos.Addr+createNamespaceUrl, map[string]string{"namespaceId": namespaceId, "namespaceName": namespaceName, "namespaceDesc": namespaceDesc})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) UpdateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Put[R[bool]](client.httpClient, client.config.Nacos.Addr+updateNamespaceUrl, map[string]string{"namespaceId": namespaceId, "namespaceName": namespaceName, "namespaceDesc": namespaceDesc})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) DeleteNamespace(namespaceId string) (bool, error) {
	r, err := Delete[R[bool]](client.httpClient, client.config.Nacos.Addr+deleteNamespaceUrl, map[string]string{"namespaceId": namespaceId})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) GetConfigs(namespaceId string) (*[]ConfigResp, error) {
	r, err := Get[R[[]ConfigResp]](client.httpClient, client.config.Nacos.Addr+getConfigListUrl, map[string]string{
		"namespaceId": namespaceId,
	})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) GetConfig(dataId string) (*string, error) {
	r, err := Get[R[string]](client.httpClient, client.config.Nacos.Addr+getConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": client.config.Nacos.Namespace,
		"group":       client.config.Nacos.Group,
	})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) DeleteConfig(dataId string) (bool, error) {
	r, err := DeleteByQuery[R[bool]](client.httpClient, client.config.Nacos.Addr+deleteConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": client.config.Nacos.Namespace,
		"group":       client.config.Nacos.Group,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) PublishConfig(dataId string, content string, configType string) (bool, error) {
	r, err := Post[R[bool]](client.httpClient, client.config.Nacos.Addr+publishConfigUrl, map[string]string{
		"dataId":      dataId,
		"namespaceId": client.config.Nacos.Namespace,
		"group":       client.config.Nacos.Group,
		"content":     content,
		"type":        configType,
	})
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

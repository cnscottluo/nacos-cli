package nacos

import (
	"strings"

	"github.com/cnscottluo/nacos-cli/internal/setting"
	"github.com/cnscottluo/nacos-cli/internal/types"
)

// Client Nacos client
type Client struct {
	config     *types.Config
	httpClient *HttpClient
}

// NewClient new client
func NewClient(config *types.Config) *Client {
	client := new(Client)
	setting.DecryptConfig(config)
	setting.InitConfig(config)
	client.config = config
	client.httpClient = NewHttpClient(config, client)
	return client
}

// Login login
func (client *Client) Login(addr string, username string, password string) (*LoginResp, error) {
	r, err := Post[LoginResp](
		client.httpClient, addr+LoginUrl, map[string]string{
			"username": username,
			"password": password,
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// UpdateUser update user
func (client *Client) UpdateUser(username string, password string) (string, error) {
	r, err := Put[R[string]](
		client.httpClient, client.config.Nacos.Addr+UserUrl, map[string]string{
			"username":    username,
			"newPassword": password,
		},
	)
	if err != nil {
		return "", err
	}
	return *r.Data, nil
}

// InitAdmin init admin
func (client *Client) InitAdmin(password string) (*InitAdminResp, error) {
	r, err := Post[InitAdminResp](
		client.httpClient, client.config.Nacos.Addr+AdminUrl, map[string]string{
			"password": password,
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetNamespaces get namespaces
func (client *Client) GetNamespaces() (*[]NamespaceResp, error) {
	r, err := Get[R[[]NamespaceResp]](client.httpClient, client.config.Nacos.Addr+GetNamespaceListUrl, nil)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// GetNamespace get namespace
func (client *Client) GetNamespace(namespaceId string) (*NamespaceResp, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	r, err := Get[R[NamespaceResp]](
		client.httpClient, client.config.Nacos.Addr+GetNamespaceUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
		},
	)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// CreateNamespace create namespace
func (client *Client) CreateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Post[R[bool]](
		client.httpClient, client.config.Nacos.Addr+CreateNamespaceUrl, map[string]string{
			"namespaceId":   processNamespace(namespaceId),
			"namespaceName": namespaceName,
			"namespaceDesc": namespaceDesc,
		},
	)
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// UpdateNamespace update namespace
func (client *Client) UpdateNamespace(namespaceId string, namespaceName string, namespaceDesc string) (bool, error) {
	r, err := Put[R[bool]](
		client.httpClient, client.config.Nacos.Addr+UpdateNamespaceUrl, map[string]string{
			"namespaceId":   processNamespace(namespaceId),
			"namespaceName": namespaceName,
			"namespaceDesc": namespaceDesc,
		},
	)
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// DeleteNamespace delete namespace
func (client *Client) DeleteNamespace(namespaceId string) (bool, error) {
	r, err := Delete[R[bool]](
		client.httpClient, client.config.Nacos.Addr+DeleteNamespaceUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
		},
	)
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// GetConfigs get configs
func (client *Client) GetConfigs(namespaceId string) (*[]ConfigResp, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	r, err := Get[R[[]ConfigResp]](
		client.httpClient, client.config.Nacos.Addr+GetConfigListUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
		},
	)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// GetNamespaceId get namespace id
func (client *Client) GetNamespaceId(namespaceId string) string {
	if namespaceId == "" {
		return client.config.Nacos.Namespace
	} else {
		return namespaceId
	}
}

// GetGroup get group
func (client *Client) GetGroup(group string) string {
	if group == "" {
		return client.config.Nacos.Group
	} else {
		return group
	}
}

// GetConfig get setting
func (client *Client) GetConfig(namespaceId string, group string, dataId string) (string, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := Get[R[string]](
		client.httpClient, client.config.Nacos.Addr+GetConfigUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"group":       group,
			"dataId":      dataId,
		},
	)
	if err != nil {
		return "", err
	}
	return *r.Data, nil
}

// DeleteConfig delete setting
func (client *Client) DeleteConfig(namespaceId string, group string, dataId string) (bool, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := DeleteByQuery[R[bool]](
		client.httpClient, client.config.Nacos.Addr+DeleteConfigUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"group":       group,
			"dataId":      dataId,
		},
	)
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

// ApplyConfig publish setting
func (client *Client) ApplyConfig(
	namespaceId string, group string, dataId string, content string, configType string,
) (bool, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := Post[R[bool]](
		client.httpClient, client.config.Nacos.Addr+PublishConfigUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"group":       group,
			"dataId":      dataId,
			"content":     content,
			"type":        configType,
		},
	)
	if err != nil {
		return false, err
	}
	return *r.Data, nil
}

func (client *Client) GetServices(namespaceId, group string) (*ServiceResp, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := Get[R[ServiceResp]](
		client.httpClient, client.config.Nacos.Addr+GetServiceListUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"groupName":   group,
			"pageNo":      "1",
			"pageSize":    "100",
		},
	)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) GetService(namespaceId string, group string, serviceName string) (*ServiceDetailResp, error) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := Get[R[ServiceDetailResp]](
		client.httpClient, client.config.Nacos.Addr+GetServiceUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"serviceName": serviceName,
			"groupName":   group,
		},
	)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (client *Client) GetServiceInstances(namespaceId string, group string, serviceName string) (
	*InstanceResp, error,
) {
	if namespaceId == "" {
		namespaceId = client.config.Nacos.Namespace
	}
	if group == "" {
		group = client.config.Nacos.Group
	}
	r, err := Get[R[InstanceResp]](
		client.httpClient, client.config.Nacos.Addr+GetInstanceListUrl, map[string]string{
			"namespaceId": processNamespace(namespaceId),
			"serviceName": serviceName,
			"groupName":   group,
		},
	)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// processNamespace process namespace
func processNamespace(namespaceId string) string {
	if strings.ToLower(namespaceId) == "public" {
		return ""
	}
	return namespaceId
}

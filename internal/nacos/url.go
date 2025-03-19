package nacos

// LoginUrl login url
const LoginUrl = "/v1/auth/login"

// UserUrl user url
const UserUrl = "/v1/auth/users"

// AdminUrl admin url
const AdminUrl = "/v1/auth/admin"

var NoAuthUrl = []string{LoginUrl, AdminUrl}

const (
	// GetConfigListUrl get setting list url
	GetConfigListUrl = "/v2/cs/history/configs"

	// GetConfigUrl get setting url
	GetConfigUrl = "/v2/cs/setting"

	// PublishConfigUrl publish setting url
	PublishConfigUrl = "/v2/cs/setting"

	// DeleteConfigUrl delete setting url
	DeleteConfigUrl = "/v2/cs/setting"
)

const (
	// GetServiceListUrl get service list url
	GetServiceListUrl = "/v2/ns/service/list"

	// GetServiceUrl get service url
	GetServiceUrl = "/v2/ns/service"

	// GetInstanceListUrl get service instance list url
	GetInstanceListUrl = "/v2/ns/instance/list"
)

const (
	// GetNamespaceListUrl get namespace list url
	GetNamespaceListUrl = "/v2/console/namespace/list"

	// GetNamespaceUrl get namespace url
	GetNamespaceUrl = "/v2/console/namespace"

	// CreateNamespaceUrl create namespace url
	CreateNamespaceUrl = "/v2/console/namespace"

	// UpdateNamespaceUrl update namespace url
	UpdateNamespaceUrl = "/v2/console/namespace"

	// DeleteNamespaceUrl delete namespace url
	DeleteNamespaceUrl = "/v2/console/namespace"
)

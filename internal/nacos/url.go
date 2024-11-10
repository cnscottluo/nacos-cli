package nacos

import "strings"

// 登录 POST
const loginUrl = "/v1/auth/login"

const getConfigListUrl = "/v2/cs/history/configs"

// 获取配置 GET
const getConfigUrl = "/v2/cs/config"

// 发布配置 POST
const publishConfigUrl = "/v2/cs/config"

// 删除配置 DELETE
const deleteConfigUrl = "/v2/cs/config"

// 查询服务列表 GET
const getServiceListUrl = "/v2/ns/service/list"

// 查询命名空间列表 GET
const getNamespaceListUrl = "/v2/console/namespace/list"

// 查询具体命名空间 GET
const getNamespaceUrl = "/v2/console/namespace"

// 创建命名空间 POST
const createNamespaceUrl = "/v2/console/namespace"

// 编辑命名空间 PUT
const updateNamespaceUrl = "/v2/console/namespace"

// 删除命名空间 DELETE
const deleteNamespaceUrl = "/v2/console/namespace"

func IsLogin(url string) bool {
	return strings.Contains(url, loginUrl)
}

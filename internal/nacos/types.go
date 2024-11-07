package nacos

type EnvConfig struct {
	Addr      string `json:"addr" yaml:"addr"`
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Group     string `json:"group" yaml:"group"`
	Auth      bool   `json:"auth" yaml:"auth"`
	Token     string `json:"token" yaml:"token"`
}

type Config struct {
	Nacos EnvConfig `json:"nacos" yaml:"nacos"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    uint64 `json:"tokenTtl"`
	GlobalAdmin bool   `json:"globalAdmin"`
	Username    string `json:"username"`
}

type NamespaceResp struct {
	Namespace         string `json:"namespace"`
	NamespaceShowName string `json:"namespaceShowName"`
	NamespaceDesc     string `json:"namespaceDesc"`
	Quota             int    `json:"quota"`
	ConfigCount       int    `json:"configCount"`
	Type              int    `json:"type"`
}

type R[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

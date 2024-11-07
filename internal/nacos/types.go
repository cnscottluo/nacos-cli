package nacos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
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

type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

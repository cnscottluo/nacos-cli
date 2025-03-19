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

type InitAdminResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NamespaceResp struct {
	Namespace         string `json:"namespace"`
	NamespaceShowName string `json:"namespaceShowName"`
	NamespaceDesc     string `json:"namespaceDesc"`
	Quota             int    `json:"quota"`
	ConfigCount       int    `json:"configCount"`
	Type              int    `json:"type"`
}

type ConfigResp struct {
	Id               string `json:"id"`
	DataId           string `json:"dataId"`
	Group            string `json:"group"`
	Content          string `json:"content"`
	Md5              string `json:"md5"`
	EncryptedDataKey string `json:"encryptedDataKey"`
	Tenant           string `json:"tenant"`
	AppName          string `json:"appName"`
	Type             string `json:"type"`
	LastModified     int64  `json:"lastModified"`
}

type ServiceResp struct {
	Count    int      `json:"count"`
	Services []string `json:"services"`
}

type ServiceDetailResp struct {
	Namespace        string         `json:"namespace"`
	GroupName        string         `json:"groupName"`
	ServiceName      string         `json:"serviceName"`
	ClusterMap       map[string]any `json:"clusterMap"`
	Metadata         map[string]any `json:"metadata"`
	ProtectThreshold float64        `json:"protectThreshold"`
	Selector         any            `json:"selector"`
	Ephemeral        bool           `json:"ephemeral"`
}

type InstanceResp struct {
	Name                     string `json:"name"`        // 分组名@@服务名
	GroupName                string `json:"groupName"`   // 分组名
	Cluster                  string `json:"cluster"`     // 集群名
	CacheMillis              int    `json:"cacheMillis"` // 缓存时间
	Hosts                    []Host `json:"hosts"`       // 实例列表
	LastRefTime              int    `json:"lastRefTime"` // 上次刷新时间
	Checksum                 string `json:"checksum"`    // 校验码
	AllIPs                   bool   `json:"allIPs"`
	ReachProtectionThreshold bool   `json:"reachProtectionThreshold"` // 是否到达保护阈值
	Valid                    bool   `json:"valid"`                    // 是否有效
}

type Host struct {
	Ip                        string         `json:"ip"`                        // 实例IP
	Port                      int            `json:"port"`                      // 实例端口号
	Weight                    float64        `json:"weight"`                    // 实例权重
	Healthy                   bool           `json:"healthy"`                   // 实例是否健康
	Enabled                   bool           `json:"enabled"`                   // 实例是否可用
	Ephemeral                 bool           `json:"ephemeral"`                 // 是否为临时实例
	ClusterName               string         `json:"clusterName"`               // 实例所在的集群名称
	ServiceName               string         `json:"serviceName"`               // 服务名
	Metadata                  map[string]any `json:"metadata"`                  // 实例元数据
	InstanceHeartBeatTimeOut  int            `json:"instanceHeartBeatTimeOut"`  // 实例心跳超时时间
	IpDeleteTimeout           int            `json:"ipDeleteTimeout"`           // 实例删除超时时间
	InstanceHeartBeatInterval int            `json:"instanceHeartBeatInterval"` // 实例心跳间隔
}

type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

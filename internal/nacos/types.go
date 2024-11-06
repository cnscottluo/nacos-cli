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

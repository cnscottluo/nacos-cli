package types

type NacosConfig struct {
	Addr      string `json:"addr" yaml:"addr"`
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Group     string `json:"group" yaml:"group"`
	Token     string `json:"token" yaml:"token"`
}

type Config struct {
	Nacos NacosConfig `json:"nacos" yaml:"nacos"`
}

# Viper

## 优先级

- explicit call to Set
- flag
- env
- config
- key/value store
- default

## 设置值

### 默认值

```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

### 读取配置文件

```go
viper.SetConfigName("config") // name of config file (without extension)
viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
viper.AddConfigPath("/etc/appname/") // path to look for the config file in
viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
viper.AddConfigPath(".") // optionally look for config in the working directory
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
panic(fmt.Errorf("fatal error config file: %w", err))
}
```

### 写入配置文件

```go
viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

### 监控配置文件变更

```go
viper.OnConfigChange(func (e fsnotify.Event) {
fmt.Println("Config file changed:", e.Name)
})
viper.WatchConfig()
```

### 读取流

```go
viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

// any approach to require this configuration into your program.
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") // this would be "steve"
```

### 注册别名

```go
viper.RegisterAlias("loud", "Verbose")
```

### 环境变量

### 标志




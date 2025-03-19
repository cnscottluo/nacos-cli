package setting

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/viper"
)

const (
	defaultKey = "83bEkaEAzg1kqJaoYWCisy0QI4XYX73w"
	keyEnv     = "NACOS_CLI_KEY"
)

const (
	DefaultAddr      = "http://127.0.0.1:8848/nacos"
	DefaultUsername  = "nacos"
	DefaultPassword  = "nacos"
	DefaultNamespace = "public"
	DefaultGroup     = "DEFAULT_GROUP"
)

const (
	NacosAddrKey      = "nacos.addr"
	NacosUsernameKey  = "nacos.username"
	NacosPasswordKey  = "nacos.password"
	NacosNamespaceKey = "nacos.namespace"
	NacosGroupKey     = "nacos.group"
	NacosTokenKey     = "nacos.token"
)

var configFileName = ".nacos.toml"

// CreateIfNotExistConfigFile creates the config file if it does not exist
func CreateIfNotExistConfigFile() {
	home, err := os.UserHomeDir()
	internal.CheckErr(err)

	if _, err := os.Stat(path.Join(home, configFileName)); os.IsNotExist(err) {
		_, err = os.Create(path.Join(home, configFileName))
		internal.CheckErr(err)
	}
}

// DeleteConfigFile deletes the config file
func DeleteConfigFile() {
	home, err := os.UserHomeDir()
	internal.CheckErr(err)

	if _, err := os.Stat(path.Join(home, configFileName)); os.IsNotExist(err) {
		return
	}
	err = os.Remove(path.Join(home, configFileName))
	internal.CheckErr(err)
}

// SaveSetting saves the Nacos setting to the config file
func SaveSetting(accessToken string) error {
	viper.Set(NacosTokenKey, accessToken)
	internal.Info("%s login success", viper.GetString(NacosUsernameKey))
	EncryptViper()
	return viper.WriteConfig()
}

// EncryptViper encrypts the Nacos password and token in the viper config
func EncryptViper() {
	viper.Set(NacosPasswordKey, encryptStr(viper.GetString(NacosPasswordKey)))
	viper.Set(NacosTokenKey, encryptStr(viper.GetString(NacosTokenKey)))
}

// DecryptConfig  decrypts the Nacos password and token in the config
func DecryptConfig(config *types.Config) {
	config.Nacos.Password = decrypt(config.Nacos.Password)
	config.Nacos.Token = decrypt(config.Nacos.Token)
}

// InitConfig initializes the Nacos config with default values
func InitConfig(config *types.Config) {
	if config.Nacos.Addr == "" {
		config.Nacos.Addr = DefaultAddr
	}
	if config.Nacos.Username == "" {
		config.Nacos.Username = DefaultUsername
	}
	if config.Nacos.Password == "" {
		config.Nacos.Password = DefaultPassword
	}
	if config.Nacos.Namespace == "" {
		config.Nacos.Namespace = DefaultNamespace
	}
	if config.Nacos.Group == "" {
		config.Nacos.Group = DefaultGroup
	}
}

// GetKey returns the defaultKey from the environment variable or the default defaultKey
func getKey() []byte {
	env := os.Getenv(keyEnv)
	if env != "" {
		return []byte(env)
	}
	return []byte(defaultKey)
}

// EncryptStr encrypts the given plaintext
func encryptStr(plaintext string) string {
	if plaintext == "" {
		return ""
	}
	ciphertext, err := aesEncryptWithGCM(getKey(), []byte(plaintext))
	internal.CheckErr(err)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Decrypt decrypts the given ciphertext
func decrypt(ciphertext string) string {
	if ciphertext == "" {
		return ""
	}
	decodeCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	internal.CheckErr(err)
	plaintext, err := aesDecryptWithGCM(getKey(), decodeCiphertext)
	internal.CheckErr(err)
	return string(plaintext)
}

// AesEncryptWithGCM encrypt the given plaintext using AES-GCM with the given defaultKey
func aesEncryptWithGCM(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AesDecryptWithGCM decrypts the given ciphertext using AES-GCM with the given defaultKey
func aesDecryptWithGCM(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

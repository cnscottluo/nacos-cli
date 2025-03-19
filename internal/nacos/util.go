package nacos

import (
	"path/filepath"
	"strings"
)

// DetermineConfigType determines the type of the configuration file based on the file extension.
func DetermineConfigType(filePath string) string {
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filePath)), ".")
	return StandardizeConfigType(ext, "text")
}

// StandardizeConfigType standardizes the configType to a known type.
func StandardizeConfigType(configType string, defaultType string) string {
	switch strings.ToLower(configType) {
	case "properties":
		return "properties"
	case "xml":
		return "xml"
	case "json":
		return "json"
	case "html", "htm":
		return "html"
	case "yaml", "yml":
		return "yaml"
	case "text":
		return "text"
	case "toml":
		return "toml"
	default:
		return defaultType
	}
}

// DetermineDataId determines the dataId of the configuration file based on the file name.
func DetermineDataId(file string) string {
	return filepath.Base(file)
}

// IsValidConfigType checks if the configType is valid.
func IsValidConfigType(configType string) bool {
	standardizeConfigType := StandardizeConfigType(configType, "")
	if standardizeConfigType == "" {
		return false
	} else {
		return true
	}
}

// IsLoginApi checks if the url is a login url.
func IsLoginApi(url string) bool {
	return strings.Contains(url, LoginUrl)
}

// IsNoAuthApi checks if the url is a no auth url.
func IsNoAuthApi(url string) bool {
	for _, u := range NoAuthUrl {
		if strings.Contains(url, u) {
			return true
		}
	}
	return false
}

// IsV1Api check if the url is a v1 api
func IsV1Api(url string) bool {
	return strings.Contains(url, "/v1/")
}

// IsV2Api check if the url is a v2 api
func IsV2Api(url string) bool {
	return strings.Contains(url, "/v2/")
}

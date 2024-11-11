package nacos

import (
	"path/filepath"
	"strings"
)

// DetermineConfigType determines the type of the configuration file based on the file extension.
func DetermineConfigType(filePath string) string {
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filePath)), ".")
	switch ext {
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
	default:
		return "text"
	}
}

// DetermineDataId determines the dataId of the configuration file based on the file name.
func DetermineDataId(file string) string {
	return filepath.Base(file)
}

// IsValidConfigType checks if the configType is valid.
func IsValidConfigType(configType string) bool {
	switch strings.ToLower(configType) {
	case "properties", "xml", "json", "html", "yaml", "text":
		return true
	default:
		return false
	}
}

// IsLogin checks if the url is a login url.
func IsLogin(url string) bool {
	return strings.Contains(url, LoginUrl)
}

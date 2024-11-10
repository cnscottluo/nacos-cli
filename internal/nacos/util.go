package nacos

import (
	"path"
	"strings"
)

// DetermineConfigType determines the type of the configuration file based on the file extension.
func DetermineConfigType(filePath string) string {
	ext := strings.TrimPrefix(strings.ToLower(path.Ext(filePath)), ".")
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
func DetermineDataId(filePath string) string {
	return path.Base(filePath)
}

// IsValidConfigType checks if the configType is valid.
func IsValidConfigType(configType string) bool {
	switch configType {
	case "properties", "xml", "json", "html", "yaml", "text":
		return true
	default:
		return false
	}
}

// IsLogin checks if the url is a login url.
func IsLogin(url string) bool {
	return strings.Contains(url, loginUrl)
}

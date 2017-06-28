package conf

import (
	"os"
	"strings"
)

func IsTest() bool { return GetEnv() == "test" }
func IsDev() bool  { return GetEnv() == "dev" }

func Get(key, defaultVal string) string {
	if key != strings.ToUpper(key) {
		panic("conf.Get() called with non-uppercase key: " + key)
	}
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func Set(key, value string) {
	os.Setenv(key, value)
}

func GetEnv() string {
	return Get("ENV", "dev")
}

// URLs
///////

func GetSiteUrl(path string) string {
	switch GetEnv() {
	case "test":
		return "http://localhost:9080" + path
	case "dev":
		return "http://dev.example.com" + path
		// return "https://example.ngrok.io" + path
	default:
		panic("TODO: GetSiteUrl")
	}
}

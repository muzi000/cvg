package security

import (
	"net/url"
	"strings"

	"github.com/muzi000/vuln-go/helper"
)

// CheckUrl   if sec, return true
func CheckUrl(u string) bool {
	if u == "" {
		return false
	}
	if len(u) == 1 {
		return true
	}
	u = strings.ReplaceAll(u, "\\", "/")
	if u[0] == '/' && u[1] == '/' {
		return false
	}
	whileList := []string{"localhost", "127.0.0.1"}
	uri, _ := url.Parse(u)
	if uri == nil {
		return false
	}
	if uri.Hostname() == "" && !strings.Contains(uri.Path, "//") {
		return true
	}
	// 支持一级域名
	if helper.ListContain(uri.Hostname(), whileList) {
		return true
	}
	// 支持多级域名
	if helper.ListSubContain(uri.Hostname(), whileList) {
		return true
	}
	return false
}

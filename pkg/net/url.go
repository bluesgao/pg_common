package net

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

// 解析 URL 字符串为 host 和 port
func ParseHostPort(raw string) (host string, port string, err error) {
	// 添加 scheme，如果没有则默认 http
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "http://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", "", fmt.Errorf("URL 解析失败: %v", err)
	}

	hostInURL := u.Host

	// 如果 host 中没有端口，添加默认端口
	if !strings.Contains(hostInURL, ":") {
		switch u.Scheme {
		case "https":
			hostInURL = net.JoinHostPort(hostInURL, "443")
		default:
			hostInURL = net.JoinHostPort(hostInURL, "80")
		}
	}

	// 分离 host 和 port
	host, port, err = net.SplitHostPort(hostInURL)
	if err != nil {
		return "", "", fmt.Errorf("分离 host 和 port 失败: %v", err)
	}

	return host, port, nil
}

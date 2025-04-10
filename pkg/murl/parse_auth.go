package murl

import (
	"fmt"
	stdUrl "net/url"
)

// ParseAuth 解析 URL 中的用户名和密码, 去除并返回 URL, Username, Password
func ParseAuth(aURL string) (url string, username string, password string, err error) {
	if aURL == "" {
		err = fmt.Errorf("URL is empty")
		return
	}

	urlParse, err := stdUrl.Parse(aURL)
	if err != nil {
		err = fmt.Errorf("URL parse error: %v", err)
		return
	}

	url = fmt.Sprintf("%s://%s%s%s", urlParse.Scheme, urlParse.Host, urlParse.Path, urlParse.RawQuery)
	username = urlParse.User.Username()
	password, _ = urlParse.User.Password()
	return
}

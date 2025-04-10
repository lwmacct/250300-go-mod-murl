package murl

import (
	"fmt"
	"net/url"
)

// ParseAuth 解析 URL 中的用户名和密码, 去除并返回 URL, Username, Password
func ParseAuth(aURL string) (URL string, Username string, Password string, err error) {
	if aURL == "" {
		err = fmt.Errorf("URL is empty")
		return
	}

	urlParse, err := url.Parse(aURL)
	if err != nil {
		err = fmt.Errorf("URL parse error: %v", err)
		return
	}

	URL = fmt.Sprintf("%s://%s%s%s", urlParse.Scheme, urlParse.Host, urlParse.Path, urlParse.RawQuery)
	Username = urlParse.User.Username()
	Password, _ = urlParse.User.Password()
	return
}

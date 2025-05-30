package murl

import (
	"fmt"
	"net/url"
)

type Parse struct {
	Raw *url.URL

	Scheme   string
	Host     string
	Port     string
	Path     string
	RawQuery string
	Username string
	Password string
}

// ParseAuth 解析 URL 中的用户名和密码, 去除并返回 URL, Username, Password
func NewParse(URL string) (*Parse, error) {
	if URL == "" {
		return nil, fmt.Errorf("URL is empty")
	}

	p, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("URL parse error: %v", err)
	}

	info := &Parse{
		Scheme:   p.Scheme,
		Host:     p.Host,
		Port:     p.Port(),
		Path:     p.Path,
		RawQuery: p.RawQuery,
	}

	info.Username = p.User.Username()
	info.Password, _ = p.User.Password()

	return info, nil
}

func (p *Parse) GetBaseURL() string {
	return fmt.Sprintf("%s://%s%s", p.Scheme, p.Host, p.Path)
}

func (p *Parse) GetNotAuthURL() string {
	return fmt.Sprintf("%s://%s%s%s", p.Scheme, p.Host, p.Path, p.RawQuery)
}

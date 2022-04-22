package utils

import (
	"net/url"
  	"strings"
)

// return `http` instead of `http://`
func GetProtocol(url string) string {
	return url[0:strings.Index(url, `://`)]
}

func GetDomain(rawURL string) (string, error) {
	u, err := url.Parse(rawURL )
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")
	return parts[len(parts)-2] + "." + parts[len(parts)-1], nil
}

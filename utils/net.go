package utils

import (
  	"strings"
)

// return `http` instead of `http://`
func GetProtocol(url string) string {
	return url[0:strings.Index(url, `://`)]
}


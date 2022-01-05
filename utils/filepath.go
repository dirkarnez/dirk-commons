package utils

import (
  	"strings"
)

func GetFileName(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}

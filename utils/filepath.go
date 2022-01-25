package utils

import (
  	"strings"
)

func GetFileName(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}

// valid: `P:\testing`, `P:\testing\`
// returning `testing`
func GetFolderName(input string) string {
	lastIndex := strings.LastIndex(input, `\`)
	length := len(input)
	if lastIndex+1 == length {
		lastIndex = strings.LastIndex(input[0:length-1], `\`)
		return input[lastIndex+1 : length-1]
	} else {
		return input[lastIndex+1 : length]
	}
}

package utils

import (
  "strings"
	"time"
)

// LocalDateStringForFileName
func LocalDateStringForFileName() {
  return strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "-")
}

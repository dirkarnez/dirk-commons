package utils

import (
	"encoding/csv"
	"io"
	"strings"
)

// ReadFileAsCSV read file as CSV
func ReadFileAsCSV(path string, onEachRecord func(map[string]int, []string)) error {
	content, err := ReadFileAsString(path)
	if err != nil {
		return err
	}
	var titleMap = make(map[string]int)
	r := csv.NewReader(strings.NewReader(content))
	for i := 0; ; i++ {
		recordArr, err := r.Read()
		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}
		if i > 0 {
			onEachRecord(titleMap, recordArr)
		} else {
			for j := 0; j < len(recordArr); j++ {
				titleMap[recordArr[j]] = j
			}
		}
	}
}

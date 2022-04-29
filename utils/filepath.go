package utils

import (
  	"strings"
)

func GetFileName(path string) string {
	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
}


// fmt.Println(GetFolderName(`P`))           //P
// fmt.Println(GetFolderName(`P:`))          //P
// fmt.Println(GetFolderName(`P:\`))         //P
// fmt.Println(GetFolderName(`P:\testing`))  //testing
// fmt.Println(GetFolderName(`P:\testing\`)) //testing
func GetFolderName(input string) string {
	var folderName string
	lastIndex := strings.LastIndex(input, `\`)
	length := len(input)
	if lastIndex+1 == length {
		lastIndex = strings.LastIndex(input[0:length-1], `\`)
		folderName = input[lastIndex+1 : length-1]
	} else {
		folderName = input[lastIndex+1 : length]
	}

	if folderName[len(folderName)-1:] == ":" {
		return folderName[0 : len(folderName)-1]
	} else {
		return folderName
	}
}

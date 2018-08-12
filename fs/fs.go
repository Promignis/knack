package fs

import (
	"io/ioutil"

	"github.com/promignis/knack/utils"
)

var FileState map[string]*File

// get file data if it is there
func GetFileData(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	utils.CheckErr(err)
	return data
}

// default permission of read and write
func WriteFileData(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0644)
}

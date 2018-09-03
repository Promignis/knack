package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"

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

func AddFileToState(path string, info os.FileInfo, err error) error {
	_, fileName := filepath.Split(path)

	if !utils.IsBlackListed(fileName) {

		data := GetFileData(path)

		// remove null check and make them globally
		// for now
		if FileState == nil {
			FileState = make(map[string]*File)
		}

		FileState[fileName] = NewFile(fileName, path, data)
	}

	return nil
}

func GetFileStat(filePath string) *FileStat {
	file, err := os.Open(filePath)
	utils.CheckErr(err)
	stat, err := file.Stat()
	utils.CheckErr(err)

	defer file.Close()

	return NewFileStat(stat)
}

func GetFileList(filePath string) FileList {
	var fileList FileList

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		fileList = append(fileList, NewFileStat(info))
		return nil
	})

	utils.CheckErr(err)

	return fileList
}

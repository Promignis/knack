package utils

import (
	"os"
	"path/filepath"

	"github.com/promignis/knack/constants"
)

// get path of the binary when it is running
// for getting relative paths to other folders
func GetRootPath() string {
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	CheckErr(err)
	return root
}

// black listed names,
// walk is being called for the
// root path as well
func IsBlackListed(fileName string) bool {
	blackListed := [4]string{constants.ViewFolder, constants.RuntimeJsPath, constants.JsFolder, constants.CssFolder}

	for i := 0; i < len(blackListed); i++ {
		if fileName == blackListed[i] {
			return true
		}
	}
	return false
}

// tmp error function
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

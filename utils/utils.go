package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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
	blackListed := [5]string{constants.ViewFolder, constants.RuntimeJsPath, constants.JsFolder, constants.CssFolder, constants.ImageFoler}

	for _, blackListedFile := range blackListed {
		if fileName == blackListedFile {
			return true
		}
	}
	return false
}

// tmp error function
func CheckErr(err error) {
	if err != nil {
		// print the error
		fmt.Print(err.Error())
		panic(err)
	}
}

func getOS() string {
	return runtime.GOOS
}

func IsUnixBased() bool {

	// darwin freebsd linux
	if getOS() != "windows" {
		return true
	}

	return false
}

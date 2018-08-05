package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// get file data if it is there
func getFileData(file_path string) []byte {
	data, err := ioutil.ReadFile(file_path)
	checkErr(err)
	return data
}

// get path of the binary when it is running
// for getting relative paths to other folders
func getRootPath() string {
	if root == "" {
		cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		checkErr(err)
		root = cwd
	}

	return root
}

// black listed names,
// walk is being called for the
// root path as well
func isBlackListed(file_name string) bool {
	black_listed := [3]string{"views", "runtime", "js"}

	for i := 0; i < len(black_listed); i++ {
		if file_name == black_listed[i] {
			return true
		}
	}
	return false
}

// tmp error function
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

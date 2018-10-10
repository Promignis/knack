package persistance

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/promignis/knack/app"
	"github.com/promignis/knack/fs"
)

var lock sync.Mutex
var userdataFilePath = filepath.Join(app.GetUserDataPath(), "/", app.GetAppName())

func init() {

	// In case the directory does not exist , create it before we start creating files inside it
	if _, err := os.Stat(userdataFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(userdataFilePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func Set(filename string, stringifiedJson string) {
	// Putting this here so that read and write on the same file can not happen. Can be clenaed up in future.
	lock.Lock()
	defer lock.Unlock()
	name := filename + ".json"
	path := filepath.Join(userdataFilePath, name)
	fs.WriteFileData(path, []byte(stringifiedJson))
}

func Get(filename string) string {
	lock.Lock()
	defer lock.Unlock()
	name := filename + ".json"
	path := filepath.Join(userdataFilePath, name)
	stringifiedJson := fs.GetFileData(path)
	return string(stringifiedJson)
}

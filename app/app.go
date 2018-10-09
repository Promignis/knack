package app

import (
	"encoding/json"
	"path/filepath"

	"github.com/promignis/knack/constants"
	"github.com/promignis/knack/fs"
	"github.com/promignis/knack/utils"
)

type Manifest struct {
	AppName string
}

var manifest Manifest

func parseManifest() Manifest {
	if manifest == (Manifest{}) {
		root := utils.GetRootPath()
		manifestData := fs.GetFileData(filepath.Join(root, constants.Manifest))
		json.Unmarshal(manifestData, &manifest)
	}
	return manifest
}

func GetAppName() string {
	return parseManifest().AppName
}

func GetUserDataPath() string {
	return globalSettingFolder
}

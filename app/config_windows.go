package app

import "os"

var systemSettingFolders = []string{os.Getenv("PROGRAMDATA")}
var globalSettingFolder = os.Getenv("APPDATA")

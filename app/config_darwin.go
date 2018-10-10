package app

import "os"

var systemSettingFolders = []string{"/Library/Application Support"}
var globalSettingFolder = os.Getenv("HOME") + "/Library/Application Support"

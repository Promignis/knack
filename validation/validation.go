package validation

import (
	"os"
	"strings"
)

// TODO:
// if empty path i.e no file there but still valid
// if file is there and operation is allowed like appending to file for save
// if file is there for open case
func IsValidPath(path string) bool {
	var validationResult bool
	if strings.Contains(path, string(os.PathSeparator)) {
		validationResult = true
	} else {
		validationResult = false
	}
	return validationResult
}

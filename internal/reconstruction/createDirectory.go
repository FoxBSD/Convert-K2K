package reconstruction

import (
	"convert-k2k/internal/debug"
	"os"
	"path/filepath"
)

// This function receives the old Path and the newPath, and create the directory of the new path. At same time calls the function to copy
func CreateDirectory(oP string, nP string) {
	dir := filepath.Dir(nP)

	os.MkdirAll(dir, 0777)
	CopyFile(oP, nP)

	debug.CompletedOperation("Create the directory with sucessful")
}

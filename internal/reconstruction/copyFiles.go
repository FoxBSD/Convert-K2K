package reconstruction

import (
	"convert-k2k/internal/debug"
	"io"
	"os"
)

// This function is used to copy the file from old repository in reconstructed repository
func CopyFile(src string, dst string) {
	debug.CopyFileMessage(src, dst)
	sourceFile, err := os.Open(src)
	if err != nil {
		debug.PrintErrorMessage("open files", err)
		return
	}

	destFile, err := os.Create(dst)
	if err != nil {
		debug.PrintErrorMessage("create files", err)
		return
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		debug.PrintErrorMessage("copying files", err)
		return
	}

	err = destFile.Sync()
	if err != nil {
		debug.PrintErrorMessage("syncing files", err)
		return
	}

	debug.CopyCompletedOperation("copied file with success", src, dst)
}

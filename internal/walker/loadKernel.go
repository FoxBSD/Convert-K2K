package walker

import (
	"convert-k2k/internal/debug"
	"path/filepath"
)

// This function return the directory of work and add a sys (form of kernel in BSD files)
func LoadKernel(d string, notIsBSD bool) string {
	if !notIsBSD {
		debug.CompletedOperation("Load the kernel path with sucess")
		return filepath.Join(d, "sys")
	}

	return d
}

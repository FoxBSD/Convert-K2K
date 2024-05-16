package reconstruction

import (
	"convert-k2k/internal/debug"
	"convert-k2k/internal/structures"
	"path/filepath"
	"strings"
)

// This function parses the old path in a new path
func CreateNewPath(f structures.FileData, n string) {
	// TODO: Add a flag to flattern teh entire repository files.
	words := strings.Split(f.FilePath, "/")
	exts := strings.Split(f.GetExt(), ".")

	reconstructed := "reconstructed-" + n

	if words[0] != ".." {
		words[0] = "./"
	}

	// BUG: The script directory may have binnary files as well.
	if f.GetExt() == "" {
		words[1] = filepath.Join(reconstructed, "script")
	}

	if f.GetExt() != "" {
		words[1] = filepath.Join(reconstructed, strings.ToLower(exts[1]))
	}

	word := strings.Join(words, "/")

	debug.CompletedOperation("Create New Path with sucess")
	CreateDirectory(f.GetPath(), word)
}

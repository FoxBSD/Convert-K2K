package reconstruction

import (
	"convert-k2k/internal/debug"
	"convert-k2k/internal/structures"
	"fmt"
	"path/filepath"
	"strings"
)

// This function parses the old path in a new path
func CreateNewPath(f structures.FileData, n string) {
	words := strings.Split(f.FilePath, "/")
	exts := strings.Split(f.GetExt(), ".")

	reconstructed := fmt.Sprintf("reconstructed-%s", n)

	if words[0] != ".." {
		words[0] = "./"
	}

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

package reconstruction

import (
	"convert-k2k/internal/debug"
	"convert-k2k/internal/structures"
	"strings"
)

// This function parses the old path in a new path
func CreateNewPath(f structures.FileData) {
	words := strings.Split(f.FilePath, "/")
	exts := strings.Split(f.GetExt(), ".")

	if f.GetExt() == "" {
		words[1] = "reconstructed/script"
	}

	if f.GetExt() != "" {
		words[1] = "reconstructed/" + strings.ToLower(exts[1])
	}

	word := strings.Join(words, "/")

	debug.CompletedOperation("Create New Path with sucess")
	CreateDirectory(f.GetPath(), word)
}

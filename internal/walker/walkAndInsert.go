package walker

import (
	"convert-k2k/internal/db"
	"convert-k2k/internal/debug"
	"convert-k2k/internal/structures"
	"database/sql"
	"io/fs"
	"path/filepath"
)

// This private function get the kernel directory and return a list of file metadata
func walk(dir string) ([]structures.FileData, error) {
	var files []structures.FileData

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)

			files = append(files, structures.FileData{
				FileName:  info.Name(),
				FilePath:  path,
				Extension: ext,
			})
		}

		return nil
	})
	if err != nil {
		debug.PrintErrorMessage("Walking in Repo", err)
		return nil, err
	}

	debug.CompletedOperation("The path reachs to the end")
	return files, nil
}

// This function load the directory and the database to insert a file
func WalkAndInsert(dir string, database *sql.DB) {
	kdir := LoadKernel(dir)

	files, _ := walk(kdir)

	for _, file := range files {
		debug.InsertMessage(file.FileName)
		db.InsertFile(file.FilePath, file.FileName, file.Extension, database)
	}

	debug.CompletedOperation("sqlite is populated")
}

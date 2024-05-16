package db

import (
	"convert-k2k/internal/debug"
	"convert-k2k/internal/structures"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// This private function is used in backend of createDB and is used to create the principal table of sqlite3
func populateTable(db *sql.DB) {
	createTable := `CREATE TABLE IF NOT EXISTS repos (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "fullPath" TEXT UNIQUE,
  "filename" TEXT,
  "extension" TEXT
  )`

	_, err := db.Exec(createTable)
	if err != nil {
		debug.PrintErrorMessage("populate a empty table", err)
		return
	}
}

// This function creates and popule a instance of sqlite3
func CreateDB(N string) *sql.DB {
	dbName := fmt.Sprintf("repos_%s.db", N)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		debug.PrintErrorMessage("Open database", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		debug.PrintErrorMessage("Ping database", err)
	}

	debug.CompletedOperation("Sucessful connect to the sqlite3 database!")

	populateTable(db)
	return db
}

// This function insert some file metadata in sqlite3
func InsertFile(fullPath string, fileName string, extension string, database *sql.DB) {
	insertFileString := `INSERT INTO repos (fullPath, filename, extension) VALUES (?, ?, ?)`

	_, err := database.Exec(insertFileString, fullPath, fileName, extension)
	if err != nil {
		debug.PrintErrorMessage("insert a new file", err)
		return
	}
}

// This function returns a list of filemetadata in sqlite3
func GetRepos(database *sql.DB) []structures.FileData {
	var files []structures.FileData

	rows, err := database.Query("SELECT fullPath, filename, extension from repos")
	if err != nil {
		debug.PrintErrorMessage("get repos from database", err)
		return nil
	}

	for rows.Next() {
		var file structures.FileData

		err := rows.Scan(&file.FilePath, &file.FileName, &file.Extension)
		if err != nil {
			debug.PrintErrorMessage("Scan the repos", err)
			return nil
		}

		files = append(files, file)
	}

	return files
}

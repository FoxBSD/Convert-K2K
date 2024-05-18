package main

import (
	"convert-k2k/internal/db"
	"convert-k2k/internal/debug"
	"convert-k2k/internal/files"
	"convert-k2k/internal/parser"
	"convert-k2k/internal/reconstruction"
	"convert-k2k/internal/walker"
	"database/sql"
	"flag"
	"os"
)

var database *sql.DB

func main() {
	argv, err := parser.Parse(flag.CommandLine)
	if err != nil {
		debug.PrintErrorMessage(argv.SourceDir, err)

		os.Exit(1)
	}

	file := files.GetFileName(argv.SourceDir)
	database = db.CreateDB(file)

	defer database.Close()

	walker.WalkAndInsert(argv.SourceDir, argv.IsNotBSD, database)

	// NOTE: Since the file information was already in memory, the previous
	// function could just return that information to the next steps.
	repos := db.GetRepos(database)

	for _, repo := range repos {
		reconstruction.CreateNewPath(repo, file)
	}
}

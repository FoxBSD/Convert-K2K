package main

import (
	"convert-k2k/internal/db"
	"convert-k2k/internal/debug"
	"convert-k2k/internal/files"
	"convert-k2k/internal/reconstruction"
	"convert-k2k/internal/walker"
	"database/sql"
	"errors"
	"flag"
)

var database *sql.DB

// NOTE: Put tests in it's own package.

func main() {
	// TODO: Separate this logic into a parser package, with tests included.
	dir := flag.String("dir", "", "this flag is used to pass the BSD OS dir")
	notIsBSD := flag.Bool("nbsd", true, "this flag is used to verify if is a bsd project")
	flag.Parse()

	if *dir != "" {
		file := files.GetFileName(*dir)
		database = db.CreateDB(file)

		walker.WalkAndInsert(*dir, *notIsBSD, database)

		// NOTE: Since the file information was already in memory, the previous
		// function could just return that information to the next steps.
		repos := db.GetRepos(database)

		for _, repo := range repos {
			reconstruction.CreateNewPath(repo, file)
		}

		defer database.Close()
		return
	}

	debug.PrintErrorMessage(*dir, errors.New("we need the flag -dir with BSD OS path to convert"))
}

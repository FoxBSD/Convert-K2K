package main

import (
	"convert-k2k/internal/db"
	"convert-k2k/internal/debug"
	"convert-k2k/internal/reconstruction"
	"convert-k2k/internal/walker"
	"database/sql"
	"errors"
	"flag"
)

var database *sql.DB

func main() {
	dir := flag.String("dir", "", "this flag is used to pass the BSD OS dir")
	flag.Parse()

	if *dir != "" {
		database = db.CreateDB()

		walker.WalkAndInsert(*dir, database)

		repos := db.GetRepos(database)

		for _, repo := range repos {
			reconstruction.CreateNewPath(repo)
		}

		defer database.Close()
		return
	}

	debug.PrintErrorMessage(*dir, errors.New("we need the flag -dir with BSD OS path to convert"))
}

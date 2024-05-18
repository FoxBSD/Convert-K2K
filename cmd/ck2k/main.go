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
	"os"
)

// NOTE: Put tests in it's own package.

var database *sql.DB

func main() {
	argv, err := parser(flag.CommandLine)
	if err != nil {
		debug.PrintErrorMessage(argv.sourceDir, err)

		os.Exit(1)
	}

	file := files.GetFileName(argv.sourceDir)
	database = db.CreateDB(file)

	defer database.Close()

	walker.WalkAndInsert(argv.sourceDir, argv.isNotBSD, database)

	// NOTE: Since the file information was already in memory, the previous
	// function could just return that information to the next steps.
	repos := db.GetRepos(database)

	for _, repo := range repos {
		reconstruction.CreateNewPath(repo, file)
	}
}

// Used in the `parser()` function (on the main package) as a return type,
// mainly to store the user CLI arguments in a easy access way.
type args struct {
	sourceDir string
	isNotBSD  bool
}

// Parse the user arguments given a `flag.FlagSet`, normally `flag.CommandLine`
// to have access to the shell's command line arguments. That's because it
// becames easier to test without the need to execute system commands, only
// updating the OS'es flags in a custom `flag.FlagSet` for unit testing.
func parser(fs *flag.FlagSet) (args, error) {
	argv := args{}

	fs.StringVar(&argv.sourceDir, "dir", "", "this flag is used to pass the BSD OS dir")
	fs.BoolVar(&argv.isNotBSD, "nbsd", true, "this flag is used to verify if is a bsd project")

	fs.Parse(os.Args[1:])

	return argv, validateArgumentValues(argv)
}

// Function helper fro the `parser(*flag.FlagSet)`, used to check the flag
// custom/default values and return an error in case something is wrong.
func validateArgumentValues(argv args) error {
	if argv.sourceDir == "" {
		return errors.New("we need the flag -dir with BSD OS path to convert")
	}

	return nil
}

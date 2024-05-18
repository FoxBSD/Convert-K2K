package parser

import (
	"errors"
	"flag"
	"os"
)

// Parse the user arguments given a `flag.FlagSet`, normally `flag.CommandLine`
// to have access to the shell's command line arguments. That's because it
// becames easier to test without the need to execute system commands, only
// updating the OS'es flags in a custom `flag.FlagSet` for unit testing.
func Parse(fs *flag.FlagSet) (Args, error) {
	argv := Args{}

	fs.StringVar(&argv.SourceDir, "dir", "", "this flag is used to pass the BSD OS dir")
	fs.BoolVar(&argv.IsNotBSD, "nbsd", true, "this flag is used to verify if is a bsd project")

	fs.Parse(os.Args[1:])

	return argv, validateArgumentValues(argv)
}

// Function helper fro the `parser(*flag.FlagSet)`, used to check the flag
// custom/default values and return an error in case something is wrong.
func validateArgumentValues(argv Args) error {
	if argv.SourceDir == "" {
		return errors.New("we need the flag -dir with BSD OS path to convert")
	}

	return nil
}

package parser_test

import (
	"convert-k2k/internal/parser"
	"errors"
	"flag"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Should throw a -dir flag required error", func(t *testing.T) {
		// Create a new FlagSet for testing
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		// Parse the flags
		_, err := parser.Parse(fs)

		// Check if the error is the expected one
		expectedError := errors.New("we need the flag -dir with BSD OS path to convert")
		if err == nil || err.Error() != expectedError.Error() {
			t.Fatalf("expected error '%v', got '%v'", expectedError, err)
		}
	})

	t.Run("Should return the struct the default isNotBSD", func(t *testing.T) {
		// Create a new FlagSet for testing
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		// Simulate the -dir argument
		os.Args = []string{"cmd", "-dir", "some/path"}

		// Parse the flags
		argv, err := parser.Parse(fs)

		// Check if no error is returned
		if err != nil {
			t.Fatalf("unexpected error '%v'", err)
		}

		// Check if SourceDir is correctly set
		if argv.SourceDir != "some/path" {
			t.Errorf("expected SourceDir to be 'some/path', got '%v'", argv.SourceDir)
		}

		// Check if the default value for isNotBSD is true
		if !argv.IsNotBSD {
			t.Errorf("expected isNotBSD to be true, got '%v'", argv.IsNotBSD)
		}
	})

	t.Run("Should return the struct with isNotBSD setted to false", func(t *testing.T) {
		// Create a new FlagSet for testing
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		// Simulate the -dir and -nbsd=false arguments
		os.Args = []string{"cmd", "-dir", "some/path", "-nbsd=false"}

		// Parse the flags
		argv, err := parser.Parse(fs)

		// Check if no error is returned
		if err != nil {
			t.Fatalf("unexpected error '%v'", err)
		}

		// Check if SourceDir is correctly set
		if argv.SourceDir != "some/path" {
			t.Errorf("expected SourceDir to be 'some/path', got '%v'", argv.SourceDir)
		}

		// Check if isNotBSD is set to false
		if argv.IsNotBSD {
			t.Errorf("expected isNotBSD to be false, got '%v'", argv.IsNotBSD)
		}
	})
}

package parser

// Used in the `parser()` function (on the main package) as a return type,
// mainly to store the user CLI arguments in a easy access way.
type Args struct {
	SourceDir string
	IsNotBSD  bool
}

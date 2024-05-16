package files

import "os"

// This code receives a string and return the file name
func GetFileName(s string) string {
	file, _ := os.Stat(s)

	return file.Name()
}

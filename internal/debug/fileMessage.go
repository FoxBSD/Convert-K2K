package debug

import "fmt"

// This function is used in copy function to synalize the copy between old path and new path
func CopyFileMessage(oP string, nP string) {
	fmt.Printf("🚛 %s -> %s\n", oP, nP)
}

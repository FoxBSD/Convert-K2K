package debug

import "fmt"

// This function is used in copy operation to sinalize is everything OK
func CopyCompletedOperation(m string, oP string, dP string) {
	nMessage := fmt.Sprintf("%s (%s -> %s)", m, oP, dP)
	fmt.Printf("✅ the operation %s is completed\n", nMessage)
}

// This function is used to sinalize is everything OK
func CompletedOperation(m string) {
	fmt.Printf("✅ the operation %s is completed\n", m)
}

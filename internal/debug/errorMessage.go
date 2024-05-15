package debug

import "fmt"

// This function prints a error message formatted with "Can't perform x action because y, being x the action and y the error"
func PrintErrorMessage(s string, err error) {
	message := fmt.Sprintf("⚠️  Can't perform %s action because %v", s, err)
	fmt.Println(message)
}

package debug

import "fmt"

// This function is used in InsertFile to synalize the file inserted in sqlite3
func InsertMessage(fileName string) {
	insertMessage := fmt.Sprintf("☑️  Insert a file: %s in database", fileName)

	fmt.Println(insertMessage)
}

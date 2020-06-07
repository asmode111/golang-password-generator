package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetLength get password length from the command line
func GetLength() int64 {
	for {
		length := ask("Enter password length: ")
		if _, err := strconv.Atoi(length); err == nil {
			iLength, err := strconv.ParseInt(length, 10, 64)
			if err != nil {
				fmt.Println(err)
			} else {
				return iLength
			}
		}
	}
}

// GetBooleanValue get boolean value from the command line
func GetBooleanValue(description string) bool {
	for {
		input := ask(description)
		if input == "y" {
			return true
		}

		if input == "n" {
			return false
		}
	}
}

func ask(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(question)
	scanner.Scan()

	return scanner.Text()
}
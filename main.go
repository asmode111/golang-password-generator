package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	length := getLength()
	fmt.Println(length)

	isAlphaNumeric := getBooleanValue("Do you need alpha numeric character? y(Yes) or n(No): ")
	fmt.Println(isAlphaNumeric)

	isNumeric := getBooleanValue("Do you need numeric character? y(Yes) or n(No): ")
	fmt.Println(isNumeric)
}

func getLength() int64 {
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

func getBooleanValue(description string) bool {
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

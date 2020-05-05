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

	isAlphaNumeric := getIsAlpaNumeric()
	fmt.Println(isAlphaNumeric)

	isNumeric := getIsNumeric()
	fmt.Println(isNumeric)
}

func getLength() int64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter password length: ")
		scanner.Scan()
		length := scanner.Text()
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

func getIsAlpaNumeric() bool {
	scanner := bufio.NewScanner(os.Stdin)
	isAlphaNumeric := true
	for {
		fmt.Print("Do you need alpha numeric character? y(Yes) or n(No): ")
		scanner.Scan()
		isAlphaNumericGiven := scanner.Text()
		if isAlphaNumericGiven == "y" {
			isAlphaNumeric = true
			break
		} else if isAlphaNumericGiven == "n" {
			isAlphaNumeric = false
			break
		}
	}

	return isAlphaNumeric
}

func getIsNumeric() bool {
	scanner := bufio.NewScanner(os.Stdin)
	isNumeric := true
	for {
		fmt.Print("Do you need numeric character? y(Yes) or n(No): ")
		scanner.Scan()
		isNumericGiven := scanner.Text()
		if isNumericGiven == "y" {
			isNumeric = true
			break
		} else if isNumericGiven == "n" {
			isNumeric = false
			break
		}
	}

	return isNumeric
}

package main

import (
	"fmt"

	helper "github.com/onurdegerli/golang-password-generator/helpers"
	model "github.com/onurdegerli/golang-password-generator/models"
	service "github.com/onurdegerli/golang-password-generator/services"
)

func main() {
	passwordLength := helper.GetLength()
	containsAlphanumeric := helper.GetBooleanValue("Do you want any alpha numeric characters? y(Yes) or n(No): ")
	containsNumeric := helper.GetBooleanValue("Do you want any numeric characters? y(Yes) or n(No): ")
	containsLowercase := helper.GetBooleanValue("Do you want any lowercase characters? y(Yes) or n(No): ")
	containsUppercase := helper.GetBooleanValue("Do you want any uppercase characters? y(Yes) or n(No): ")

	option := model.Option{
		PasswordLength:       passwordLength,
		ContainsAlphanumeric: containsAlphanumeric,
		ContainsNumeric:      containsNumeric,
		ContainsLowercase:    containsLowercase,
		ContainsUppercase:    containsUppercase,
	}

	password, err := service.Generate(option)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your unique password is %s\n", password)
	}
}

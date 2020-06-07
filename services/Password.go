package services

import (
	"errors"
	"math/rand"
	"time"

	model "github.com/onurdegerli/golang-password-generator/models"
)

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowercase = "abcdefghijklmnoprstuvwyz"
const _charsetUppercase = "ABCDEFGHIJKLMOPRSTUWYZ"
const _charsetNumerics = "0123456789"
const _charsetAlphanumeric = "=-_[.~]!#+&"

// Generate generates password
func Generate(option model.Option) (string, error) {
	x := make([]byte, option.PasswordLength)
	charset := "."

	if option.ContainsLowercase {
		charset += _charsetLowercase
	}

	if option.ContainsUppercase {
		charset += _charsetUppercase
	}

	if option.ContainsNumeric {
		charset += _charsetNumerics
	}

	if option.ContainsAlphanumeric {
		charset += _charsetAlphanumeric
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}

	if charset == "." {
		return "NON", errors.New("Please enter a character")
	}

	return string(x), nil
}

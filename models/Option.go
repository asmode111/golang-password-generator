package models

// Option model
type Option struct {
	PasswordLength       int64
	ContainsAlphanumeric bool
	ContainsNumeric      bool
	ContainsLowercase    bool
	ContainsUppercase    bool
}

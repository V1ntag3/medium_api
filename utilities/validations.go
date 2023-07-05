package utilities

import (
	"regexp"
	"unicode"
)

// verify if string is blank
func OnlyEmptySpaces(str string) bool {

	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// verify if string contains number
func ContainsNumber(str string) bool {

	for _, char := range str {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

// verify email
func ValidateEmail(email string) bool {

	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)

	return match
}

// verify passwork
// pre-requires
// min 8 characteres
// min 1 special character
// min 1 uppercase
// min 1 number
func ValidatePassword(password string) bool {
	reLen := `.{8,}$`
	reUpperCase := `.*([A-Z]+).*$`
	reSpecialChar := `(.*[!#@$%&]+).*`
	reNumber := `(.*[0-9]+).*`

	matchLen, _ := regexp.MatchString(reLen, password)
	matchUpp, _ := regexp.MatchString(reUpperCase, password)
	matchSpe, _ := regexp.MatchString(reSpecialChar, password)
	matchNum, _ := regexp.MatchString(reNumber, password)

	if matchLen && matchUpp && matchSpe && matchNum {
		return true
	} else {
		return false
	}
}

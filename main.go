package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasUpper, hasDigit := false, false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	if hasUpper && hasDigit {
		return true
	} else {
		return false
	}
}

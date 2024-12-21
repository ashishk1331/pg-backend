package util

import "regexp"

func ValidateEmail(email string) bool {
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func ValidateHeaderToken(token string) (string, bool) {
	if token == "" {
		return "", false
	}
	const bearerPrefix = "Bearer "
	if len(token) > len(bearerPrefix) && token[:len(bearerPrefix)] == bearerPrefix {
		token = token[len(bearerPrefix):]
		return token, true
	}
	return "", false
}

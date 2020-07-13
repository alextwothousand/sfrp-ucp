package password

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	matchLower *regexp.Regexp = regexp.MustCompile(`[a-z]`)
	matchUpper *regexp.Regexp = regexp.MustCompile(`[A-Z]`)

	matchNumber  *regexp.Regexp = regexp.MustCompile(`[0-9]`)
	matchSpecial *regexp.Regexp = regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\\\)\-_\=\+\,\.\?\/\:\;\{\}\[\]~]`)

	cost int = 12
)

// Hash encrypts your password using the bcrypt encryption method.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

// Match asserts whether the hash and the raw password entered match each other.
func Match(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Integrity checks whether a password has a good enough score for it to be used.
func Integrity(password string) bool {
	if len(password) < 6 || len(password) > 255 {
		return false
	}

	score := 0

	if matchLower.MatchString(password) {
		score++
	}

	if matchUpper.MatchString(password) {
		score++
	}

	if matchNumber.MatchString(password) {
		score++
	}

	if matchSpecial.MatchString(password) {
		score++
	}

	if score >= 3 {
		return true
	}

	return false
}

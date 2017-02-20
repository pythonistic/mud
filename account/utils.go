package account

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const PASSWORD_COST = 12

func buildPassword(email, rawPassword string) []byte {
	return []byte(strings.TrimSpace(rawPassword) + strings.ToLower(strings.TrimSpace(email)))
}

func hashPassword(email, rawPassword string) ([]byte, error) {
	password := buildPassword(email, rawPassword)
	return bcrypt.GenerateFromPassword(password, PASSWORD_COST)
}

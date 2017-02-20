package account

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	email string
	password []byte
}

func (acct *Account) String() string {
	return fmt.Sprintf("Account{email=%s,password=%s}", acct.email, acct.password)
}

func (acct *Account) ComparePassword(candidate string) bool {
	candidatePassword := buildPassword(acct.email, candidate)
	err := bcrypt.CompareHashAndPassword(acct.password, candidatePassword)
	if err != nil {
		fmt.Printf("INFO: failed login for %s: %v\n", acct.email, err.Error())
		return false
	} else {
		fmt.Printf("DEBUG: password match for %s\n", acct.email)
		return true
	}
}

func New(email, rawPassword string) (acct *Account) {
	hash, err := hashPassword(email, rawPassword)
	if err != nil {
		fmt.Printf("WARN: failed to create account for email %s: %v\n", email, err.Error())
		return
	}
	acct = &Account{
		email: email,
		password: hash,
	}
	return
}
package account

import (
	"testing"
)

func TestNew(t *testing.T) {
	email := "nobody@dev.null"
	password := "f41c87b5eb36c425a1165a419099131c"

	acct := New(email, password)
	if acct == nil {
		t.Error("Account was nil")
		t.FailNow()
	}
	if acct.email != email {
		t.Errorf("Account email should be %s, got %s", email, acct.email)
		t.FailNow()
	}
	if len(acct.password) != 60 {
		t.Errorf("Account password should be 60 bytes, got %d bytes", len(acct.password))
		t.FailNow()
	}
}

func TestAccount_ComparePassword(t *testing.T) {
	email := "nobody@dev.null"
	password := "f41c87b5eb36c425a1165a419099131c"
	failingPw := "Hello, World!"

	acct := New(email, password)

	if !acct.ComparePassword(password) {
		t.Error("Identical password failed validation")
		t.FailNow()
	}
	if acct.ComparePassword(failingPw) {
		t.Error("Failing password was erroneously allowed")
		t.FailNow()
	}
}

func TestGetLoginMessage(t *testing.T) {
	ltp := loginTextPath
	loginTextPath = "../" + loginTextPath
	defer func() {
		loginTextPath = ltp
	}()
	msg := GetLoginMessage()
	if len(msg.Content) == 0 {
		t.Error("Login message ewas empty")
		t.FailNow()
	}
}

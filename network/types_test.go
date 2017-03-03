package network

import (
	"testing"
	"fmt"
	"time"
)

func TestMessage_RoundTrip(t *testing.T) {
	createTime := time.Now()
	m := Message{
		Created: createTime,
		Content: []byte("Burning all the bridges now"),
		Kind: MT_COMBAT,
	}
	b := m.ToBytes()
	if len(b) < 45 {
		t.Errorf("Message should have been 90 bytes, was %d", len(b))
		t.FailNow()
	}
	m1 := FromBytes(&Client{}, b)
	expectedContent := fmt.Sprintf("%s:%d:%s", MT_COMBAT, createTime.Unix(), string(m.Content))
	if string(m1.Content) != expectedContent {
		t.Errorf("Decoded content mismatch; expected=%s decoded=%s", expectedContent, m1.Content)
		t.FailNow()
	}
	if m1.Kind != MT_FROM_CLIENT {
		t.Errorf("Decoded content type wrong; expected=%s got=%s", MT_FROM_CLIENT, m1.Kind)
		t.FailNow()
	}
}

func TestNewAccount(t *testing.T) {
	email := "nobody@dev.null"
	password := "f41c87b5eb36c425a1165a419099131c"

	acct := NewAccount(email, password)
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

	acct := NewAccount(email, password)

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

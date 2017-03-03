package network

import (
	"io/ioutil"
	"fmt"
	"time"
	"sync"
)

var maxLoginMessageAge = time.Duration(30) * time.Minute
var loginMessage *Message
var loginMessageExpires time.Time
var loginMessageMutex = sync.Mutex{}
var loginTextPath = "text/login.txt"

func GetLoginMessage() *Message {
	loginMessageMutex.Lock()
	defer loginMessageMutex.Unlock()
	if loginMessage == nil || time.Now().After(loginMessageExpires) {
		loginMessage = NewMessage(loginBanner(), MT_SYSTEM)
		loginMessageExpires = time.Now().Add(maxLoginMessageAge)
	}
	return loginMessage
}

func loginBanner() []byte {
	login, err := ioutil.ReadFile(loginTextPath)
	if err != nil {
		fmt.Printf("Error reading login.txt: %v\n", err.Error())
	}
	return login
}

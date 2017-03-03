package network

import (
	"mud/parser"
)

func HandleClient(client *Client) {
	// handle the client state
	if !client.LoggedIn {
		// give the client the login screen
		client.Write(GetLoginMessage())
	}
}

func HandleMessage(msg *Message) {
	// parse the command from the message
	command := parser.TokenizeMessage(string(msg.Content))

	// get the client (for convenience)
	client := msg.Client

	// get the client Location
	location := client.Location

	// pass the command to the location
	errMsg := location.ProcessCommand(client, command)

	if errMsg != "" {
		// create an error message for the client
		errorMessage := NewMessage([]byte(errMsg), MT_ERROR)
		client.Write(errorMessage)
	}
}
package network

import (
	"net"
	"time"
	"io"
	"fmt"
)

const WRITE_SLEEP = time.Duration(10) * time.Millisecond

type Client struct {
	connection net.Conn
	connected bool
	created time.Time
	messageAdapter Adapter
	removeClientChan chan<- *Client
	fromClientChan chan<- *Message
	toClientChan chan *Message

	// game state fields
	LoggedIn bool
	Account *Account
	Location *Location
}

func (client *Client) IoLoop() {
	// create the writer goroutine
	go client.writeMessages()

	for client.connected {
		// sleep to allow other goroutines to run
		time.Sleep(WRITE_SLEEP)

		// read from the client
		var b [MAX_MESSAGE_SIZE]byte
		n, err := client.connection.Read(b[0:])
		if err != nil {
			if err != io.EOF {
				fmt.Printf("WARN: error reading from client: %v\n", err.Error())
			} else {
				// close the connection
				client.connected = false
				fmt.Printf("INFO: client closed connection: %v\n", err.Error())
			}
		}

		if n > 0 {
			// push the message to the message handler channel
			message := FromBytes(client, b[0:n - 1])
			client.fromClientChan <- message
		}
	}

	// dump any existing messages before closing
	client.connection.Close()

	// remove the client from the connection pool
	println("Telling channel to remove this connection")
	client.removeClientChan <- client
}

func (client *Client) Write(msg *Message) {
	// only push messages if the client is connected
	if client.connected {
		client.toClientChan <- msg
	}

	println("Discarded client message. Connection closed.")
}

func (client *Client) writeMessages() {
	fmt.Printf("Started write messages for client %s\n", client)
	for client.connected {
		select {
		case msg := <-client.toClientChan:
			fmt.Printf("Client got msg: %s\n", msg)
			// get the next message to send

			// tcp connections make the logic to send partial messages unnecessary
			// but if we ever send UDP, we'll be happy for this
			start := 0
			b := client.messageAdapter(msg)
			for start < len(b) - 1 {
				n, err := client.connection.Write(b[start:])
				start += n
				if err != nil {
					fmt.Printf("WARN: error writing to client: %v\n", err.Error())
					break
				}
			}
		}
	}

	println("Done writing to client")
}

func (client *Client) String() string {
	return fmt.Sprintf("Client{connection=%s,connected=%t,loggedIn=%t,created=%s}",
		client.connection.RemoteAddr().String(), client.connected, client.LoggedIn, client.created.String())
}

// Disconnect the client.
func (client *Client) Disconnect() {
	client.connected = false
}


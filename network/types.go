package network

import (
	"net"
	"time"
	"fmt"
	"sync"
	"io"
)

const MAX_MESSAGE_SIZE = 4096
const WRITE_SLEEP = time.Duration(10) * time.Millisecond

type Client struct {
	connection net.Conn
	connected bool
	loggedIn bool
	created time.Time
	incomingMessage []byte
	outgoingMessages [][]byte
	outgoingMessagesMutex sync.Mutex
}

func (client *Client) Handle(disconnectHandler func(*Client)) {
	// create the writer goroutine
	go client.writeMessages()

	for client.connected {
		// sleep to allow other goroutines to run
		time.Sleep(WRITE_SLEEP)
		print(".")

		// read from the client
		var b [MAX_MESSAGE_SIZE]byte
		n, err := client.connection.Read(b[0:])
		print(n)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("WARN: error reading from client: %v\n", err.Error())
			} else {
				// close the connection
				client.connected = false
				fmt.Printf("INFO: client closd connection: %v\n", err.Error())
			}
		}

		if n > 0 {
			// push the message to the message handler channel
			fmt.Printf("message: %s\n", b)
		}
	}

	// force the client to close
	client.connection.Close()

	// remove the client from the connection pool
	disconnectHandler(client)
}

func (client *Client) Write(msg string) {
	b := []byte(msg)
	client.outgoingMessagesMutex.Lock()
	defer client.outgoingMessagesMutex.Unlock()
	client.outgoingMessages = append(client.outgoingMessages, b)
}

func (client *Client) writeMessages() {
	for client.connected {
		// we won't send more than one message per WRITE_SLEEP
		time.Sleep(WRITE_SLEEP)

		if len(client.outgoingMessages) > 0 {
			// get the next message to send
			client.outgoingMessagesMutex.Lock()
			b := client.outgoingMessages[0]
			client.outgoingMessages = append(client.outgoingMessages[:0], client.outgoingMessages[1:]...)
			client.outgoingMessagesMutex.Unlock()

			// tcp connections make the logic to send partial messages unnecessary
			// but if we ever send UDP, we'll be happy for this
			start := 0
			n := 0
			var err error
			for start < len(b) - 1 {
				n, err = client.connection.Write(b[start:])
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

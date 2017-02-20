package network

import (
	"net"
	"time"
	"fmt"
	"io"
	"bytes"
)

const MAX_MESSAGE_SIZE = 4096
const WRITE_SLEEP = time.Duration(10) * time.Millisecond

type Client struct {
	connection net.Conn
	connected bool
	loggedIn bool
	created time.Time
	removeClientChan chan<- *Client
	fromClientChan chan<- Message
	toClientChan chan Message
}

func (client *Client) Handle() {
	// create the writer goroutine
	go client.writeMessages()

	for client.connected {
		// sleep to allow other goroutines to run
		time.Sleep(WRITE_SLEEP)
		print(".")

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
			raw_message := b[0:n - 1]
			fmt.Printf("message: %s - ", raw_message)
			message := Message{
				content: string(raw_message) + "\n",
				created: time.Now(),
				kind: MT_FROM_CLIENT,
			}
			print("pushing to chan - ")
			client.fromClientChan <- message
			println("pushed")
		}
	}

	// force the client to close
	println("Forcing client to close")
	client.connection.Close()

	// remove the client from the connection pool
	println("Telling channel to remove this connection")
	client.removeClientChan <- client
}

func (client *Client) Write(msg Message) {
	client.toClientChan <- msg
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
			n := 0
			var err error
			b := msg.ToBytes()
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

func (client *Client) String() string {
	return fmt.Sprintf("Client{connection=%s,connected=%t,loggedIn=%t,created=%s}",
		client.connection.RemoteAddr().String(), client.connected, client.loggedIn, client.created.String())
}

type MessageType uint8
const (
	MT_UNKNOWN MessageType = iota
	MT_FROM_CLIENT
	MT_SYSTEM
	MT_SAY
	MT_EMOTE
	MT_TELL
	MT_DESCRIPTION
	MT_COMBAT
	MT_DISCONNECT
	MT_CONNECT
	MT_OTHER
)

func (mt MessageType) String() string {
	switch mt {
	case MT_UNKNOWN:
		return "UNKNOWN"
	case MT_FROM_CLIENT:
		return "FROM_CLIENT"
	case MT_SYSTEM:
		return "SYSTEM"
	case MT_SAY:
		return "SAY"
	case MT_EMOTE:
		return "EMOTE"
	case MT_TELL:
		return "TELL"
	case MT_DESCRIPTION:
		return "DESCRIPTION"
	case MT_COMBAT:
		return "COMBAT"
	case MT_OTHER:
		return "OTHER"
	}

	return "UNDEFINED"
}

type Message struct {
	kind MessageType
	created time.Time
	content string
}

func (m *Message) ToBytes() []byte {
	// convert the content to double-byte runes
	return []byte(fmt.Sprintf("%s:%d:%s", m.kind, m.created.Unix(), string([]rune(m.content))))
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{kind=%s,created=%s,content=%s}", m.kind, m.created, m.content)
}

func MessageFromBytes(b []byte) Message {
	r := bytes.Runes(b)
	m := Message{
		kind: MT_FROM_CLIENT,
		created: time.Now(),
		content: string(r),
	}
	return m
}
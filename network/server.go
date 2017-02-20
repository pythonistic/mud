package network

import (
	"net"
	"fmt"
	"os"
	"time"
	"sync"
)

var listening bool
var connectionPool []*Client = make([]*Client, 0)
var connectionPoolMtx sync.Mutex = sync.Mutex{}
var incomingMessages = make(chan Message)
var removeClients = make(chan *Client)
var newClients = make(chan net.Conn)

func Listen(host string) {
	listener, err := net.Listen("tcp", host)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open server %s: %v\n", host, err.Error())
		os.Exit(2)
	}

	listening = true

	go handler()

	for listening {
		// block until next connection
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("ERROR: can't accept connection:  %v", err.Error())
		} else {
			newClients<- conn
		}
	}
}

func handler() {
	for {
		select {
		case client := <-removeClients:
			removeClient(client)
		case message := <-incomingMessages:
			processMessage(message)
		case conn := <-newClients:
			createClient(conn)
		}
	}
}

func removeClient(client *Client) {
	fmt.Printf("Asked to remove client: %s\n", client)
	connectionPoolMtx.Lock()
	for idx, clientToConsider := range connectionPool {
		if clientToConsider == client {
			connectionPool = append(connectionPool[0:idx], connectionPool[idx + 1:]...)
			println("DEBUG: removed client from connection pool")
			break
		}
	}
	connectionPoolMtx.Unlock()
}

func processMessage(message Message) {
	println("started handleIncomingMessages")
	fmt.Printf("got message: %s\n", message)
	// stub out handling messages
	// push the message to all the clients
	connectionPoolMtx.Lock()
	for _, client := range connectionPool {
		client.Write(message)
	}
	connectionPoolMtx.Unlock()
}

func createClient(conn net.Conn) {
	// do something with the connection
	client := &Client{
		connection: conn,
		connected: true,
		loggedIn: false,
		created: time.Now(),
		toClientChan: make (chan Message),
		removeClientChan: removeClients,
		fromClientChan: incomingMessages,
	}

	connectionPoolMtx.Lock()

	// notify all about the connection
	connectMessage := Message{
		kind: MT_CONNECT,
		created: time.Now(),
		content: client.String() + " connected",
	}
	for _, otherClient := range connectionPool {
		otherClient.Write(connectMessage)
	}

	connectionPool = append(connectionPool, client)
	connectionPoolMtx.Unlock()

	// start the connection goroutine
	go client.Handle()
}

func shutdownConnections() {
	// forcibly disconnect remaining connections
	connectionPoolMtx.Lock()
	for _, connection := range connectionPool {
		err := connection.connection.Close()
		if err != nil {
			fmt.Printf("WARN: error closing connection: %v", err.Error())
		}
	}
	// empty the connection pool
	connectionPool = make([]*Client, 0)
	connectionPoolMtx.Unlock()
}
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
var sleep = time.Duration(10) * time.Millisecond

func Listen(host string) {
	listener, err := net.Listen("tcp", host)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open server %s: %v\n", host, err.Error())
		os.Exit(2)
	}

	listening = true

	for listening {
		// block until next connection
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("ERROR: can't accept connection:  %v", err.Error())
		} else {
			// do something with the connection
			client := &Client{
				connection: conn,
				connected: true,
				loggedIn: false,
				created: time.Now(),
			}

			connectionPoolMtx.Lock()
			connectionPool = append(connectionPool, client)
			connectionPoolMtx.Unlock()

			// start the connection goroutine
			go client.Handle(removeConnection)
		}
	}
}

func removeConnection(client *Client) {
	connectionPoolMtx.Lock()
	defer connectionPoolMtx.Unlock()
	for idx, clientToConsider := range connectionPool {
		if clientToConsider == client {
			connectionPool = append(connectionPool[0:idx], connectionPool[idx + 1:]...)
			println("DEBUG: removed client from connection pool")
			break
		}
	}
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
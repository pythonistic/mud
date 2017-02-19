package main

import (
	"mud/network"
	"os"
)

func main() {
	network.Listen("localhost:6666")
	os.Exit(0)
}

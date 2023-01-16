package main

import (
	"fmt"

	"github.com/devlongs/cli-app/internal/network"
)

func main() {
    fmt.Println("Network CLI")

    network.Ping("Hello")
}
package main

import (
	"fmt"
	"gsync/internal/server"
)

func main() {
	err := server.Run(":8090")
	if err == nil {
		fmt.Println("Bisey oldu")
	}
}

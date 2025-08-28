package main

import (
	"log"
	"main/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal("fatal error: ", err)
	}
}

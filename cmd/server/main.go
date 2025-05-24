package main

import (
	"os"
	"scrapeit/server"
)

func main() {
	topicName := os.Args[1]
	s := server.New(topicName)
	s.Run()
}

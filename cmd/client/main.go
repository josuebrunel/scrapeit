package main

import (
	"os"
	"scrapeit/client"
)

func main() {
	c := client.New(os.Args[1])
	c.Run()
}

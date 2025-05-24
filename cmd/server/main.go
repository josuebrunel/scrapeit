package main

import (
	"scrapeit/server"
)

func main() {
	a := server.NewApi()
	a.Start()
}

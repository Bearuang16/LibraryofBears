package main

import (
	config "BearLibrary/Config"
	"BearLibrary/Route"
)

func main() {
	config.InitConnection()
	e := Route.New()
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	"BearLibrary/Config"
	"BearLibrary/Route"
)

func main() {
	Config.InitConnection()
	e := Route.New()
	e.Logger.Fatal(e.Start(":8000"))
}

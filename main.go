package main

import (
	"goapi-nunu/app"
	"goapi-nunu/logs"
)

func main() {
	logs.Info("Starting application...")
	app.Start()
}

package main

import (
	"github.com/stock-market-simulator/Go/app"
)

func main() {
	e := app.MakeHandler()

	e.Logger.Fatal(e.Start(":5000"))
}

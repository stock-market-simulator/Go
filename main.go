package main

import app "github.com/stock-market-simulator/Go/controller"

func main() {
	e := app.MakeHandler()

	e.Logger.Fatal(e.Start(":5000"))
}

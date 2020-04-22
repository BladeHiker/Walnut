package main

import (
	_ "LiveAssistant/backend"
	"os"

	"github.com/go-qamel/qamel"
)

func main() {
	//Create Application
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Live Assistant")
	app.SetWindowIcon(":/res/Walnuts.ico")

	engine := qamel.NewEngine()
	engine.Load("qrc:/res/main.qml")

	// Exec app
	app.Exec()
}

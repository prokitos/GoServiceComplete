package main

import (
	"modular/internal/app"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"

	_ "modular/docs"
)

// @title User API
// @version 1.0
// @description This is a sample service for managing users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	log.SetLevel(log.DebugLevel) // показывать логи debug уровня
	//log.SetFormatter(&log.JSONFormatter{})	// выводить логи в Json формате
	//enableLogToFile()							// выводить логи в файл

	log.Info("the server is starting")

	PauseDisable() // отключение заморозки приложения при выделении текста в консоли

	app.MainServer()

}

func PauseDisable() {
	winConsole := windows.Handle(os.Stdin.Fd())
	var mode uint32
	err := windows.GetConsoleMode(winConsole, &mode)
	if err != nil {
		log.Println(err)
	}
	mode &^= windows.ENABLE_QUICK_EDIT_MODE
	err = windows.SetConsoleMode(winConsole, mode)
	if err != nil {
		log.Println(err)
	}
}

// Записывать логи в файл
func enableLogToFile() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}

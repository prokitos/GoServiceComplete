package main

import (
	"modular/internal/app"
	"os"

	log "github.com/sirupsen/logrus"
)

// Запуск сервера при старте приложения
func main() {

	log.SetLevel(log.DebugLevel) // показывать логи debug уровня
	//log.SetFormatter(&log.JSONFormatter{})	// выводить логи в Json формате
	//enableLogToFile()							// выводить логи в файл

	log.Info("the server is starting")
	app.MainServer()
}

// Записывать логи в файл
func enableLogToFile() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}

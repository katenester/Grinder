package main

import (
	"Grinder/server/internal/repository"
	"Grinder/server/internal/service"
	"Grinder/server/internal/transport"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Загружаем конфиги сервера
	if err := initConfig(); err != nil {
		log.Fatalf("error initalization config %s", err.Error())
		return
	}
	// Dependency injection for architecture application
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := transport.NewHandler(services)
	srv := transport.NewServer(handlers)
	// Запускаем сервер
	srv.Run()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

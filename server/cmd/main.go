package main

import (
	"Grinder/server/internal/transport"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Загружаем конфиги сервера
	if err := initConfig(); err != nil {
		log.Fatalf("error initalization config %s", err.Error())
		return
	}
	srv := transport.NewServer()
	fmt.Println("Server listening on port 80...")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

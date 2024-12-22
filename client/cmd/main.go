package main

import (
	"Grinder/client/internal/controller"
	"Grinder/client/internal/models"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
)

func main() {
	// Загружаем конфиги сервера
	if err := initConfig(); err != nil {
		log.Fatalf("error initalization config %s", err.Error())
		return
	}
	// Подключение к серверу через сокет tcp
	conn, err := net.Dial("tcp", viper.GetString("ip_server")+":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("error initalization config %s", err.Error())
		return
	}
	// Закрытие соединения
	defer conn.Close()
	// Создание клиента
	client := models.NewClient(conn)
	// Создание консольного обработчика
	handle := controller.NewHandler(client)
	ch := make(chan struct{})
	// Запуск блока контраля игры
	go controller.StartGame(ch, client)
	<-ch
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func Auth() string {
	for {
		fmt.Println("Please enter your username:")
		var username string
		fmt.Scanln(&username)
	}
}

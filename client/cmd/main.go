package main

import (
	"Grinder/client/internal/models"
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
	// Подключение к серверу через сокет tcpr
	conn, err := net.Dial("tcp", viper.GetString("ip_server")+":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("error initalization config %s", err.Error())
		return
	}
	// Закрытие соединения
	//defer conn.Close()

	// Создание клиента который работает с протоколом Kat
	client := models.NewClientKat(conn)
	client.ChooseUsername()
	//fmt.Println("client:", client)
	client.GetTopScores()
	//// Создание консольного обработчика
	//handle := controller.NewHandler(client)
	//ch := make(chan struct{})
	//// Запуск блока контраля игры
	//go controller.StartGame(ch, handle)
	//<-ch
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

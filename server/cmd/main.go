package main

import (
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
	ln, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server listening on port 80...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleClient(conn)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

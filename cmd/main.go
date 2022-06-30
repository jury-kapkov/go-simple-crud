package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка считывания конфигов: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Возникла ошибка запуска https сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("../configs")
	return viper.ReadInConfig()
}

package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/messages"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("%s: %s", messages.ReadConfigError, err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("%s: %s", messages.ReadEnvError, err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("%s: %s", messages.DBConnectionError, err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("%s: %s", messages.StartHTTPServerError, err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}

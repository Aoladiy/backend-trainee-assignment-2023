package main

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/handler"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("config initialization error %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env variables loading error %v", err)
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbName"),
	})
	if err != nil {
		log.Fatalf("DB initialization error %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(backendTraineeAssignment2023.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Http server error %v", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

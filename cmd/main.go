package main

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/handler"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("config initialization error %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env variables loading error %v", err)
	}

	var db *sqlx.DB
	var err error
	maxRetries := 60

	for retry := 0; retry < maxRetries; retry++ {
		db, err = repository.NewMysqlDB(repository.Config{
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			DBName:   viper.GetString("db.dbName"),
		})
		if err == nil {
			break
		}

		logrus.Printf("DB connection attempt %d failed: %v", retry+1, err)
		time.Sleep(time.Second)
	}

	if err != nil {
		logrus.Fatalf("DB initialization error after %d attempts: %v", maxRetries, err)
	}

	if err := db.Ping(); err != nil {
		logrus.Fatalf("DB connection test failed: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(backendTraineeAssignment2023.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("HTTP server error %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

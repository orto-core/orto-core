package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/orto-core/server/tenant-service/config"
	"github.com/orto-core/server/tenant-service/internal/models"
	"github.com/orto-core/server/tenant-service/internal/router"
	"github.com/orto-core/server/tenant-service/internal/store"

	"github.com/spf13/viper"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	dsn := viper.GetString("database.dsn")
	if err := store.InitDB(dsn, models.Models...); err != nil {
		log.Fatalf("Failed to initialize database ,%v", err)
	}

	r := router.RegisterRouter()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := r.Run(port); err != nil {
		panic(err)
	}
}

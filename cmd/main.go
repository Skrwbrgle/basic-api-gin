package main

import (
	"database/sql"
	"log"

	"restfull-api/m/v2/config"
	"restfull-api/m/v2/handler"
	"restfull-api/m/v2/repository"
	"restfull-api/m/v2/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := sql.Open(cfg.DbDriver, cfg.DbConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userUseCase := &usecase.UserUseCase{UserRepo: userRepository}

	router := gin.Default()
	handler.NewUserHandler(router, userUseCase)

	addr := cfg.ApiHost + ":" + cfg.ApiPort
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

package main

import (
	"restfull-api/m/v2/api/routes"
	"restfull-api/m/v2/config"
	"restfull-api/m/v2/managers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	timeout := time.Duration(2) * time.Second
	cfg, _ := config.NewConfig()
	db, _ := managers.Application(cfg)

	gin := gin.Default()

	// r.GET("/api", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Server is Online",
	// 	})

	//   if condition {

	//   }
	// })

	routes.Setup(cfg, timeout, db, gin)

	gin.Run(cfg.API.ApiPort)
}

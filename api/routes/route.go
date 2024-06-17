package routes

import (
	"restfull-api/m/v2/config"
	"restfull-api/m/v2/managers"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
}

func Setup(cfg *config.Config, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs

	app_manager, _ := managers.Application(cfg)
	// NewSignupRouter(env, timeout, db, publicRouter)
	// NewLoginRouter(env, timeout, db, publicRouter)
	// NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// // All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}

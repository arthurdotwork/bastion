package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func useCors(allowedOrigins []string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowedOrigins
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * 3600

	return cors.New(corsConfig)
}

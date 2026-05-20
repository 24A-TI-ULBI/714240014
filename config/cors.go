package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CorsConfig returns the CORS configuration for the app
func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "http://localhost:8080, http://127.0.0.1:8080",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
		AllowCredentials: true,
	}
}

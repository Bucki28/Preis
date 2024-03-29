package main

import (
	stdhttp "net/http"

	"github.com/Edw-Castro/Preis/internal/infrastructure/server/http"
	"github.com/Edw-Castro/Preis/internal/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Reemplaza con la URL de tu aplicación React
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == stdhttp.MethodOptions {
			c.AbortWithStatus(stdhttp.StatusOK)
		} else {
			c.Next()
		}
	}
}

/********************* AUTH **********************/
func main() {
	engine := gin.Default()

	// almacenamiento de sesiones
	cookieSession := cookie.NewStore([]byte(utils.SecretKey))

	// Usa el middleware de CORS personalizado
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Agrega esta línea

		if c.Request.Method == stdhttp.MethodOptions {
			c.AbortWithStatus(stdhttp.StatusOK)
			return
		}

		c.Next()
	})

	engine.Use(sessions.Sessions("session", cookieSession))

	// Registra tus rutas aquí
	http.RegisterRoutes(engine)

	engine.Run(":8081")
}

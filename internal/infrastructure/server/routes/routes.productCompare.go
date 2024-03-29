package routes

import (
	"log"

	"github.com/Bucki28/Preis/internal/infrastructure/repositories/sql"
	"github.com/Edw-Castro/Preis/internal/core/services"
	database "github.com/Edw-Castro/Preis/internal/infrastructure/database/mysql"
	server "github.com/Edw-Castro/Preis/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func ComparePreisProduct() func(c *gin.Context) {
	db1, err := database.SetupDatabaseArticleConnection()
	db2, err := database.SetupDatabaseUsersConnection()
	if err != nil {
		log.Fatalf("Error cannot connect to db: %v", err)
	}
	repoProduct := sql.NewArticleRepository(db1, db2)
	servProduct := services.NewProductService(repoProduct)
	getProductEndpoint := server.GetDetailProductEndpoint(servProduct)
	return getProductEndpoint
}

package Server

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Config"
	"SORA/Project/Go_Back_End_SEGA_Project/GraphQLModel"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartGraphQLServer() {
	// ==============================================
	CORS := cors.New(cors.Config{
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	})
	// --------------------------------------
	router := gin.Default()
	router.Use(CORS)
	// ==============================================
	router.GET("/", gin.WrapF(handler.Playground("GraphQL playground", "/query")))
	router.POST("/query", gin.WrapF(handler.GraphQL(
		GraphQLModel.NewExecutableSchema(
			GraphQLModel.Config{
				Resolvers: &GraphQLModel.Resolver{},
			},
		),
	)))
	// ==============================================

	router.GET("/loaderio-434253d2ac58483eba54001e1f0f0d69.txt", CertificationFunction)

	// ==============================================
	log.Fatal(
		router.RunTLS(
			":"+Config.GraphQLDefaultPort,
			"./SSL/server.crt",
			"./SSL/server.key",
		),
	)
	// ==============================================
}

func CertificationFunction(c *gin.Context) {
	CertificationData := "loaderio-434253d2ac58483eba54001e1f0f0d69"
	c.String(http.StatusOK, CertificationData)
}

package Server

import (
	"SORA/Base"
	"SORA/Config"
	"SORA/GraphQLModel"
	"SORA/Socket"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartGraphQLServer() {
	// ==============================================
	CORS := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowWebSockets:  true,
	})
	// -----------------------------------------------[Set User]
	router := gin.Default()
	router.Use(CORS)
	// -----------------------------------------------[Log]
	logFile, err := os.Create("./log/restful_server.log")
	if err != nil {
		Base.Warning.Println(err)
	} else {
		router.Use(gin.LoggerWithWriter(logFile))
	}
	// ==============================================[GraphQL]
	router.GET("/", gin.WrapF(handler.Playground("GraphQL playground", "/query")))
	router.POST("/query", gin.WrapF(handler.GraphQL(
		GraphQLModel.NewExecutableSchema(
			GraphQLModel.Config{
				Resolvers: &GraphQLModel.Resolver{},
			},
		),
	)))
	// ==============================================[Socket]
	router.GET("/socket", gin.WrapF(Socket.LinkSocket))
	// ==============================================[Test]
	router.GET("/loaderio-434253d2ac58483eba54001e1f0f0d69.txt", CertificationFunction)
	// ==============================================[RunTLS]
	Base.Error.Panicln(
		router.RunTLS(
			":"+Config.GraphQLDefaultPort,
			"./SSL/server.crt",
			"./SSL/server.key",
		))
	// ==============================================
}

func CertificationFunction(c *gin.Context) {
	CertificationData := "loaderio-434253d2ac58483eba54001e1f0f0d69"
	c.String(http.StatusOK, CertificationData)
}

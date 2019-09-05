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
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func StartGraphQLServer() {
	// ==============================================
	// -----------------------------------------------[Set User]
	router := gin.Default()
	CORS := cors.DefaultConfig()
	CORS.AllowAllOrigins = true
	CORS.AllowCredentials = true
	CORS.AllowWebSockets = true
	router.Use(cors.New(CORS))
	pprof.Register(router)
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
	router.GET("/socket/number", gin.WrapF(Socket.NumberSocket))
	// ==============================================[Test]
	router.GET("/loaderio-434253d2ac58483eba54001e1f0f0d69.txt", CertificationFunction)
	// ==============================================[STOP]
	router.GET("/SHUTDOWN", ServerStopFunc)
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

func ServerStopFunc(c *gin.Context) {
	if token := c.DefaultQuery("Token", ""); token == "SORAServerShutdown" {
		c.String(http.StatusOK, "SHUTDOWN SERVER")
		Base.ServerStop <- 1
	} else {
		c.String(404, "FAIL")
	}
}

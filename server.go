package main

import (
	log "log"
	http "net/http"
	os "os"
	GraphQLModel "sega/GraphQLModel"

	handler "github.com/99designs/gqlgen/handler"
	chi "github.com/go-chi/chi"
	cors "github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// ==============================================
	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	// --------------------------------------
	router := chi.NewRouter()
	router.Use(CORS.Handler)
	// ==============================================
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(GraphQLModel.NewExecutableSchema(GraphQLModel.Config{Resolvers: &GraphQLModel.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Fatal(http.ListenAndServeTLS(":"+port, "./SSL/server.crt", "./SSL/server.key", router))
}

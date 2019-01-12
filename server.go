package main

import (
	fmt "fmt"
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
	router.Get("/loaderio-434253d2ac58483eba54001e1f0f0d69.txt", CertificationFunction)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Fatal(http.ListenAndServeTLS(":"+port, "./SSL/server.crt", "./SSL/server.key", router))
}

func CertificationFunction(w http.ResponseWriter, r *http.Request) {
	CertificationData := "loaderio-434253d2ac58483eba54001e1f0f0d69"
	fmt.Fprint(w, CertificationData)
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GaelLnz/goya/api/graphql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := log.Default()
	resolver := graphql.NewResolver(logger)
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}

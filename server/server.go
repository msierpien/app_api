package server

import (
	"api/graph"
	icApi "api/libs/ic/auth"

	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "4000"

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("Błąd wczytywania pliku .env")
    }
}


func Server() {

	token, err := icApi.GetToken()
	if err != nil {
        fmt.Println(err)
        return
    }

    // Wyświetl token
    fmt.Println(token)


	port := os.Getenv("API_GRAPHQL_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

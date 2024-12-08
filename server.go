package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MW-7892/mini-grader-be/database"
	"github.com/MW-7892/mini-grader-be/graph/middleware"
	"github.com/MW-7892/mini-grader-be/graph/generated"
	"github.com/MW-7892/mini-grader-be/graph/resolver"
	"github.com/MW-7892/mini-grader-be/utils"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
  // Init database
  err := database.ConnectToMySql()
  if err != nil {
    panic(err)
  }

  // GraphQL
	port := utils.GetEnvVar("PORT")
	if port == "" {
		port = defaultPort
	}

  router := chi.NewRouter()
  router.Use(middleware.Middleware())
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

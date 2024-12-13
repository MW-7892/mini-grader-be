package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MW-7892/mini-grader-be/database"
	directive "github.com/MW-7892/mini-grader-be/graph/directives"
	"github.com/MW-7892/mini-grader-be/graph/generated"
	"github.com/MW-7892/mini-grader-be/graph/middleware"
	"github.com/MW-7892/mini-grader-be/graph/resolver"
	"github.com/MW-7892/mini-grader-be/utils"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
  err := godotenv.Load(".env")

  if err != nil {
    log.Print("No .env file found, using system env...")
  }
}


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

  // GQL directives
  config := generated.Config{Resolvers: &resolver.Resolver{}}
  directive.Init(&config)

  router := chi.NewRouter()
  router.Use(middleware.Middleware())
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

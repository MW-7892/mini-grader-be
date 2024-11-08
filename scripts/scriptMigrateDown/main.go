package main

import (
	"log"

	"github.com/MW-7892/mini-grader-be/database"
	"github.com/pressly/goose/v3"
)

func main() {
  err := database.ConnectToMySql()
  if err != nil {
      log.Fatal(err)
  }

  if err := goose.SetDialect("mysql"); err != nil {
      log.Fatal(err)
  }
  if err := goose.Down(database.SqlDB, "database/migrations"); err != nil {
      log.Fatal(err)
  }
}

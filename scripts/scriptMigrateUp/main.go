package main

import (
	"github.com/MW-7892/mini-grader-be/database"
	"github.com/pressly/goose/v3"
)

func main() {
  err := database.ConnectToMySql()
  if err != nil {
      panic(err)
  }

  if err := goose.SetDialect("mysql"); err != nil {
      panic(err)
  }
  if err := goose.Up(database.SqlDB, "database/migrations"); err != nil {
      panic(err)
  }
}

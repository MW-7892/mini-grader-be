// This is custom goose binary with sqlite3 support only.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/MW-7892/mini-grader-be/database/migrations"
	"github.com/MW-7892/mini-grader-be/utils"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
  err := godotenv.Load(".env")

  if err != nil {
    log.Print("No .env file found, using system env...")
  }
}

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = "database/migrations"
)

func main() {
  var (
    user      = utils.GetEnvVar("DB_USER")
    password  = utils.GetEnvVar("DB_PASSWORD")
    host      = utils.GetEnvVar("DB_HOST")
    port      = utils.GetEnvVar("DB_PORT")
    name      = utils.GetEnvVar("DB_NAME")
  )

  database_string := fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    user, password, host, port, name,
  )

  err := flags.Parse(os.Args[1:])
  if err != nil {
    panic(err)
  }

	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

  command := args[0]
  gormDB, err := gorm.Open(mysql.New(mysql.Config{
    DSN: database_string,
    DefaultStringSize: 256, // default size for string fields
    DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
    DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
    DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
    SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
  }), &gorm.Config{})

	if err != nil {
    panic(err)
	}

  // Somehow I cannot use gorm DB for goose migration
  db, err := goose.OpenDBWithDriver("mysql", database_string)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	if err != nil {
    panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := make([]string, 0)
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

  ctx := context.Background()
  ctx = context.WithValue(ctx, "database", gormDB)
  if err := goose.RunContext(ctx, command, db, dir, arguments...); err != nil {
      log.Fatalf("goose %v: %v", command, err)
  }
}

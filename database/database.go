package database

import (
	"database/sql"
	"fmt"

	"github.com/MW-7892/mini-grader-be/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SqlDB *sql.DB
var Migrator gorm.Migrator

func ConnectToMySql() (error) {
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

  db, err := gorm.Open(mysql.New(mysql.Config{
    DSN: database_string,
    DefaultStringSize: 256, // default size for string fields
    DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
    DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
    DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
    SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
  }), &gorm.Config{})

  if err != nil {
    return err
  }

  DB = db
  SqlDB, _ = DB.DB()
  Migrator = DB.Migrator()
  return nil
}

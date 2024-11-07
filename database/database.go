package database

import (
	"database/sql"

	"github.com/MW-7892/mini-grader-be/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SqlDB *sql.DB
var Migrator gorm.Migrator

func ConnectToMySql() (error) {
  database_string := utils.GetEnvVar("DB_USER") + ":" +
    utils.GetEnvVar("DB_PASSWORD") + "@tcp(" + 
    utils.GetEnvVar("DB_HOST") + ":" +
    utils.GetEnvVar("DB_PORT") + ")/" +
    utils.GetEnvVar("DB_NAME") +
    "?charset=utf8&parseTime=True&loc=Local"

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

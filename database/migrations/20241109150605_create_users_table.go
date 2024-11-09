package migrations

import (
	"context"
	"database/sql"

	"github.com/MW-7892/mini-grader-be/internal/models"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
  db := ctx.Value("database").(*gorm.DB)
	return db.Migrator().CreateTable(&models.User{})
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
  db := ctx.Value("database").(*gorm.DB)
	return db.Migrator().DropTable(&models.User{})
}

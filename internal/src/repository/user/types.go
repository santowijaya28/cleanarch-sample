package user

import "database/sql"

const (
	insertBySystem = -1
)

// Repository stores all dependencies needed by user repo package
type Repository struct {
	userdb *sql.DB
}

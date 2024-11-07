package user

import "database/sql"

func Init(userDB *sql.DB) *Repository {
	return &Repository{
		userdb: userDB,
	}
}

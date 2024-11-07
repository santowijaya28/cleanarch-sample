package postgres

import "database/sql"

type PQConn struct {
	userDB *sql.DB
}

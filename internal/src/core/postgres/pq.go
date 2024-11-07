package postgres

import "database/sql"

func (pq *PQConn) GetUserDB() *sql.DB {
	return pq.userDB
}

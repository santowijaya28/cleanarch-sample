package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/santowijaya28/cleanarch-sample/internal/src/config"
)

// InitConn will initialize postgres database
func InitConn(cfg *config.Config) (*PQConn, error) {
	if cfg == nil {
		return nil, errors.New("missing config")
	}

	// init wallet database
	userAppConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.DbUserName,
		cfg.DatabaseConfig.DbName,
		cfg.DatabaseConfig.DbPwd,
	)

	userDB, err := sql.Open("postgres", userAppConnStr)
	if err != nil {
		return nil, err
	}

	return &PQConn{
		userDB: userDB,
	}, nil
}

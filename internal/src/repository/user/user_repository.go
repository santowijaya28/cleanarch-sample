package user

import (
	"context"
	"errors"

	"github.com/santowijaya28/cleanarch-sample/internal/src/model"
	"github.com/santowijaya28/cleanarch-sample/pkg/hash"
)

// Insert will insert user data to database
func (r *Repository) Insert(ctx context.Context, userData model.UserData) (err error) {
	// query
	query := "INSERT INTO \"user-db\" (username, pwd, created_time, created_by)" +
		"VALUES($1,$2,now(),$3)"

	// generate hashed pwd before storing to our database
	pwdHash, err := hash.GenerateHash(userData.Pwd)
	if err != nil {
		return err
	}

	// exec the query
	res, err := r.userdb.ExecContext(
		ctx,
		query,
		userData.UserName,
		pwdHash,
		insertBySystem)

	if err != nil {
		return err
	}

	if totalData, _ := res.RowsAffected(); totalData == 0 {
		return errors.New("error when inserting data")
	}

	return nil
}

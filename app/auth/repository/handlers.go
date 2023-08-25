package authrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

func (repo *Repository) CreateUser(createdUser utils.User) (err error) {
	createUserQuery := `INSERT INTO users (ID) VALUES ($1);`

	_, err = repo.pool.Exec(context.Background(), createUserQuery, createdUser.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case literals.CodeUniqConflict:
				err = fmt.Errorf(utils.LogUserIdConflict)
			}
		}
	}

	return
}

func (repo *Repository) GetUser(userId string) (amount float64, err error) {
	getUserQuery := `SELECT balance FROM users WHERE ID = $1;`

	err = repo.pool.QueryRow(context.Background(), getUserQuery, userId).Scan(&amount)
	if err != nil && err == pgx.ErrNoRows {
		err = fmt.Errorf(utils.LogUserNotFoundError)
	}

	return
}

func (repo *Repository) UpdateBalance(userId string, newAmount float64) (err error) {
	updateBalanceQuery := `UPDATE users SET balance = $1 WHERE ID = $2;`

	_, err = repo.pool.Exec(context.Background(), updateBalanceQuery, newAmount, userId)

	return
}

func (repo *Repository) MinusBalance(userId string, minus float64) (curr float64, err error) {
	updateBalanceQuery := `UPDATE users SET balance = balance - $1 WHERE ID = $2 RETURNING balance;`

	err = repo.pool.QueryRow(context.Background(), updateBalanceQuery, minus, userId).Scan(&curr)

	return
}

func (repo *Repository) PlusBalance(userId string, plus float64) (curr float64, err error) {
	updateBalanceQuery := `UPDATE users SET balance = balance + $1 WHERE ID = $2 RETURNING balance;`

	err = repo.pool.QueryRow(context.Background(), updateBalanceQuery, plus, userId).Scan(&curr)

	return
}

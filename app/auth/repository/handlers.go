package authrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

func (repo *Repository) CreateUser(createdUser authUtils.User) (err error) {
	createUserQuery := `INSERT INTO users (ID) VALUES ($1);`

	_, err = repo.pool.Exec(context.Background(), createUserQuery, createdUser.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case literals.CodeUniqConflict:
				err = fmt.Errorf(authUtils.LogUserIdConflict)
			}
		}
	}

	return err
}

func (repo *Repository) MinusBalance(userId string, minus float64) (float64, error) {
	var (
		curr float64
		err  error
	)

	getBalanceQuery := `SELECT balance FROM users WHERE ID = $1 FOR UPDATE;`
	minusBalanceQuery := `UPDATE users SET balance = balance - $1 WHERE ID = $2 RETURNING balance;`
	ctx := context.Background()

	var tx pgx.Tx
	tx, err = repo.pool.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке создать транзакцию: %w", err)
		return curr, err
	}
	rollbackFunc := func() {
		err1 := tx.Rollback(ctx)
		if err1 != nil {
			if errors.Is(err1, pgx.ErrTxClosed) {
				return
			}

			if err != nil {
				err = fmt.Errorf("Произошла ошибка при откате транзакции: %w, до транзации ошибка: %w", err1, err)
			} else {
				err = fmt.Errorf("Произошла ошибка при откате транзакции: %w", err1)
			}
		}

		return
	}

	var balance float64
	err = tx.QueryRow(ctx, getBalanceQuery, userId).Scan(&balance)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке считать баланс: %w", err)
		rollbackFunc()
		return curr, err
	}

	if balance < minus {
		err = fmt.Errorf(utils.LogUnderfundedError)
		rollbackFunc()
		return curr, err
	}

	err = tx.QueryRow(ctx, minusBalanceQuery, minus, userId).Scan(&curr)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке обновить баланс: %w", err)
		rollbackFunc()
		return curr, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("Ошибка при коммите: %w", err)
		rollbackFunc()
		return curr, err
	}

	return curr, err
}

func (repo *Repository) PlusBalance(userId string, plus float64) (curr float64, err error) {
	getBalanceQuery := `SELECT balance FROM users WHERE ID = $1 FOR UPDATE;`
	plusBalanceQuery := `UPDATE users SET balance = balance + $1 WHERE ID = $2 RETURNING balance;`
	ctx := context.Background()

	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке создать транзакцию: %w", err)
		return curr, err
	}
	rollbackFunc := func() {
		err1 := tx.Rollback(ctx)
		if err1 != nil {
			if errors.Is(err1, pgx.ErrTxClosed) {
				return
			}

			if err != nil {
				err = fmt.Errorf("Произошла ошибка при откате транзакции: %w, до транзации ошибка: %w", err1, err)
			} else {
				err = fmt.Errorf("Произошла ошибка при откате транзакции: %w", err1)
			}
		}

		return
	}

	if _, err = tx.Exec(context.Background(), getBalanceQuery, userId); err != nil {
		rollbackFunc()
		return curr, err
	}

	err = tx.QueryRow(ctx, plusBalanceQuery, plus, userId).Scan(&curr)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке обновить баланс: %w", err)
		rollbackFunc()
		return curr, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("Ошибка при коммите: %w", err)
		rollbackFunc()
		return curr, err
	}

	return curr, err
}

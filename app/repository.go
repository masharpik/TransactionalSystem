package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func loadConfigUrl() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
}

func getConnectionDB() (conn *pgxpool.Pool, err error) {
	url := loadConfigUrl()

	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(2 * time.Minute)

	for {
		select {
		case <-timer.C:
			ticker.Stop()
			err = fmt.Errorf(literals.LogConnDBTimeout)
			return
		case <-ticker.C:
			conn, err = pgxpool.New(context.Background(), url)

			if err == nil {
				err = conn.Ping(context.Background())
				if err == nil {
					ticker.Stop()
					timer.Stop()
					logger.LogOperationSuccess(fmt.Sprintf(literals.LogConnDBSuccess, url))
					return
				}
			}
		}
	}
}

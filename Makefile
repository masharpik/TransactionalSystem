.PHONY: up down clean

up:
	docker compose up

down:
	docker compose down

clean:
	docker rm bwg-db-1
	docker volume rm bwg_db_data

test:
	go test -cover ./app/auth/delivery
	go test -cover ./app/transaction/delivery

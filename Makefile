.PHONY: up down dbclean serverclean reup

up:
	docker-compose build --no-cache
	docker compose up

down:
	docker compose down

run:
	docker compose up -d
	go run main.go run

migrate:
	go run main.go migrate

docker:
	docker compose up -d
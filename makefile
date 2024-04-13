migrate:
	migrate -path src/db/migrations -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up

create-migration:
	if [ -z $(NAME) ]; then \
  			echo "Необхрдимо указать NAME=Название миграции"; \
		else \
			migrate create -ext sql -dir ./src/db/migrations -seq $(NAME); \
		fi

up:
	cp src/.env .env
	@docker-compose up -d --build

down:
	@docker-compose down

restart:
	@make down
	@make up
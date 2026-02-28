DB_URL=postgres://postgres:postgres@localhost:5432/users_db?sslmode=disable

DOCKER_MIGRATE=docker run --rm \
	-v $(PWD)/migrations:/migrations \
	--network host \
	migrate/migrate

# Создать новую миграцию
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Применить все миграции
migrate-up:
	migrate \
	-path migrations \
	-database "$(DB_URL)" \
	up

# Откатить 1 миграцию
migrate-down:
	migrate \
	-path migrations \
	-database "$(DB_URL)" \
	down 1

# Откатить все миграции
migrate-down-all:
	migrate \
	-path migrations \
	-database "$(DB_URL)" \
	down

# Показать текущую версию
migrate-version:
	migrate \
	-path migrations \
	-database "$(DB_URL)" \
	version

#migrate-force:
#	migrate \
#	-path migrations \
#	-database "$(DB_URL)" \
#	force
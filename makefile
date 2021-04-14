start:
	bash scripts/build.sh

stop:
	docker-compose down -v

migrate:
	bash scripts/migrate.sh
run:
	go run ./cmd/app/app.go

migrate:
	docker cp ./etc/migration/migration.sql db_clean_architecture_example:/migration.sql -q
	docker exec db_clean_architecture_example psql -d db_clean_architecture_example -U postgres -f ./migration.sql

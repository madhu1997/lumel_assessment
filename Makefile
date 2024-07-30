DB_URL=postgresql://postgres:root@localhost:5433/lumel_sales_data?sslmode=disable
createdb:
	sudo -u postgres psql -c 'create database lumel_sales_data;'
dropdb:
	sudo -u postgres psql -c 'drop database lumel_sales_data;'
new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)
migrateup:
	migrate --path db/migrations/ --database "$(DB_URL)" --verbose up
migratedown:
	migrate --path db/migrations/ --database "$(DB_URL)" --verbose down
server:
	go run main.go data_loader.go
.PHONY: createdb dropdb migrateup migratedown server
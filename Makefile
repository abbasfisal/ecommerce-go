serve:
	@ go run cmd/http_server/main.go

#apply production migration
up:
	@sql-migrate up -env=production -config=internal/database/mysql/migration_config.yml

#drop production migration
down:
	@sql-migrate down -env=production -config=internal/database/mysql/migration_config.yml

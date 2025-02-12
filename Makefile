migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations up

migrate-up-force:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations -verbose force 0000001
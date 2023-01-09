db_host = zip-pg
db_port = 5432
db_user = root
db_pass = root
db_name = zip

conn = "host=${db_host} port=${db_port} user=${db_user} password=${db_pass} dbname=${db_name} sslmode=disable"

# Run from:

## zip-go:

migrate.up:
	goose -dir migrations postgres ${conn} up

migrate.down:
	goose -dir migrations postgres ${conn} down

migrate.status:
	goose -dir migrations postgres ${conn} status

migrate.new:
	goose -dir migrations postgres ${conn} create ${migration_name} sql

run:
	go build -o bin/zip_from_docker github.com/kaz-as/zip/cmd/app && \
    ./bin/zip_from_docker

## zip-gen:

gen:
	swagger generate server -A Zip --spec docs/swagger.yml --exclude-main && \
    rm restapi/configure_zip.go restapi/server.go

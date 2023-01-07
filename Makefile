host = zip-pg
port = 5432
user = root
pass = root
name = zip

conn = "host=${host} port=${port} user=${user} dbname=${name} sslmode=disable"

migrate.up:
	goose -dir migrations postgres ${conn} up

migrate.down:
	goose -dir migrations postgres ${conn} down

migrate.status:
	goose -dir migrations postgres ${conn} status

gen:
	swagger generate server -A Zip --spec docs/swagger.yml --exclude-main && \
    rm restapi/configure_zip.go restapi/server.go

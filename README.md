### CS BIRD - Monitoring tool

## Running locally
```sh
docker-compose up
```

## Migrations
1. using custom tool:
```sh
go run cmd/dbmigrate/main.go \
-migrate=up \
-dbname=csbird \
-dbhost=localhost \
-dbport=5432
```
2. with cli:
```sh
migrate -source file://migrations -database postgres://{user}:{password}@{host}:{port}/{dbname}
```

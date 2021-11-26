### CS BIRD - Monitoring tool

## Running locally
```sh
docker-compose up

go run cmd/dbmigrate/main.go -migrate=up -dbhost=localhost
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

3. sql snippet:
```
SELECT distinct on (balance) balance, coin, address, time
FROM checkpoints
WHERE time > NOW() - INTERVAL '20 hours';

select
    count(distinct balance) - 1 as balance_change_events,
    coin,
    address
from checkpoints
WHERE time > NOW() - INTERVAL '4 hours'
group by coin, address;
```

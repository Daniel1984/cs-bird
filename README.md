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

3. sql snippet:
```
SELECT distinct on (balance) balance, coin, address, time
FROM checkpoints
WHERE time > NOW() - INTERVAL '20 hours';

select count(distinct balance) balance_changes_over_24h, coin, address
from checkpoints
WHERE time > NOW() - INTERVAL '24 hours'
group by coin, address;
```

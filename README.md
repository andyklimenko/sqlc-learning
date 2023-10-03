### Run postgres container first
```shell
docker run -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
```

### Install sqlc
```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest 
```

### Generate db client
```shell
sqlc generate
```
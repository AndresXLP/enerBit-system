``` sh
Clone with ssh recommended
$ git clone git@github.com:AndresXLP/enerBit-system.git

Clone with https
$ git clone https://github.com/AndresXLP/enerBit-system.git
```

# Requirements

* go v1.20
* go modules

# Technology Stack

- [echo](https://echo.labstack.com/)
- [validator](https://github.com/go-playground/validator)
- [GORM](https://gorm.io/)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)

# Architecture

- [Hexagonal Architecture](https://www.happycoders.eu/software-craftsmanship/hexagonal-architecture/)

# Build

* Install dependencies:

```sh
$ go mod download
```

* [Migrations](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 
```sh
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
$ migrate -path ./internal/infra/resource/postgres/migrations -database postgresql://${POSTGRESQL_DB_USER}:${POSTGRESQL_DB_PASSWORD}@${POSTGRESQL_DB_HOST}:${POSTGRESQL_DB_PORT}/${POSTGRESQL_DB_NAME}?sslmode=disable up
```

* Run local
```sh
$ go run cmd/main.go
```

* Run with Docker:

```sh 
$ make compose-up 
```

# Environments

#### Required environment variables

* `SERVER_HOST`: host for the server
* `SERVER_PORT`: port for the server
* `DB_HOST`: host database
* `DB_USER`: user database
* `DB_PASSWORD`: password database
* `DB_NAME`: name database
* `DB_PORT`: port database
* `REDIS_HOST`: host redis
* `REDIS_PORT`: port redis
* `GRPC_PROTOCOL`: protocol for gRPC server
* `GRPC_HOST`: host for gRPC server
* `GRPC_PORT`: port for gRPC server


# Contributors

* Andres Puello


# Floral

## TODO

1. Add `make` rule for codegen: https://github.com/deepmap/oapi-codegen#contributing

## Development

### Dependencies

1. `sqlc`
2. [OAPI Codegen](https://github.com/deepmap/oapi-codegen#overview)

### Generate Server Side API Code

```sh
mkdir -p ./generated/api/ && \
$GOPATH/bin/oapi-codegen \
  -config ./server.api.yml \
  ./api/swagger.yaml
```

### Generate database code based on schemas & queries

For now to keep it simple only 1 package is generated

Run:

```sh
sqlc generate
```

### Start database

```sh
docker-compose up -d db
```

### Start Go application

```sh
go run *.go
```

## Build

## Run tests

### `auth` package

go test -v floral/auth

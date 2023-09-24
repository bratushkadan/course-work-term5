# Floral

## Develop

## Generate database code based on schemas & queries
For now to keep it simple only 1 package is generated

Run:
```sh
sqlc generate
```

## Build

### Production

1. Specify registry and version:

```sh
export REGISTRY=
```

```sh
export VERSION=0.0.1
```

2. Build the image

```sh
docker build --platform linux/amd64 -t "$REGISTRY/floral-server:$VERSION" .
```

3. Push the image

```sh
docker push $REGISTRY/floral-server:$VERSION
```

4. Deploy

on vm:
```sh
export BUILD_VERSION=0.0.4; ./deploy.sh
```

arbitrarily:
```sh
docker run \
  -d \
  --name floral-server \
  -e TOKEN_SECRET=<token secret> \
  -e POSTGRES_PASSWORD=<password> \
  -p 8080:8080 \
    <image_name:image_tag>
```

## Run tests

### `auth` package

go test -v floral/auth

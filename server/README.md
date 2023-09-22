# Build

## Production

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

```sh
docker run -d --name floral-server -p 8080:8080 <image_name:image_tag>
```
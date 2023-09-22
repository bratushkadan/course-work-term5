# Containerization

## Production

### Build

```sh
docker build --platform linux/amd64 -t <registry>/floral-client:0.0.2 .
```
### Deploy

```sh
docker run -p 3000:3000 --name floral-client <registry>/floral-client:0.0.2
```
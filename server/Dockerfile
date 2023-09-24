FROM golang:1.21.1-alpine3.18 AS build

RUN mkdir /src
COPY ./*.go ./go.mod ./go.sum /src/
COPY auth /src/auth
COPY database /src/database
COPY db /src/db
WORKDIR /src
RUN go get -d -v -t
RUN go build -v -o floral
RUN chmod +x floral

RUN mkdir -p /.postgresql && \
  wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" \
    --output-document /.postgresql/root.crt && \
chmod 0600 /.postgresql/root.crt

FROM scratch
COPY --from=build /src/floral /usr/local/bin/floral
COPY --from=build /.postgresql /.postgresql

ENV PGSSLROOTCERT="/.postgresql/root.crt"

EXPOSE 8080
CMD ["/usr/local/bin/floral"]

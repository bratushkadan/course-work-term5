FROM golang:1.21.1-alpine3.18 AS build

RUN mkdir /src
COPY ./*.go ./go.mod ./go.sum /src/
WORKDIR /src
RUN go get -d -v -t
RUN go build -v -o floral
RUN chmod +x floral

FROM scratch
COPY --from=build /src/floral /usr/local/bin/floral
EXPOSE 8080
CMD ["/usr/local/bin/floral"]

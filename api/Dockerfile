FROM golang:1.18 AS build

WORKDIR /app

COPY . ./
RUN go mod download

ENV GOARCH=amd64

RUN go build \
    -o /go/bin/app

## Deploy
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app

ENV ADDR=:8080
EXPOSE 8080

USER nonroot:nonroot

CMD ["/app"]

FROM golang:1.27-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/maxtasks-api ./cmd/server

FROM alpine:3.22

RUN addgroup -S maxtasks && adduser -S -G maxtasks maxtasks
COPY --from=build /out/maxtasks-api /usr/local/bin/maxtasks-api

USER maxtasks
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/maxtasks-api"]

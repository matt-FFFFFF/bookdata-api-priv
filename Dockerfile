FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/matt-FFFFFF/bookdata-api
COPY . .
# Build the binary.
RUN mkdir /app && go build -o /app/api

####

FROM busybox

COPY --from=builder /app/api /api
COPY assets/books.csv /assets/books.csv

EXPOSE 8080

ENTRYPOINT ["/api"]
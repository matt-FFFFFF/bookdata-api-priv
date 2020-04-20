FROM golang:1.14.2-alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/bookdata-api
COPY . .

# Get dependencies
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/bookdata-api .

####

FROM scratch

COPY --from=builder /go/bin/bookdata-api /
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
#COPY assets/books.csv /assets/books.csv

EXPOSE 8080

USER appuser:appuser

ENTRYPOINT ["/bookdata-api"]
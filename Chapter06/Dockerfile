FROM golang:1.17 AS builder

WORKDIR /usr/src/app

# Build go-sqlite3 package
RUN go get github.com/mattn/go-sqlite3@v1.14.11

# Install sqlite
RUN apt-get update
RUN apt-get install sqlite3
RUN sqlite3 htmlT.db "CREATE TABLE data (number INTEGER PRIMARY KEY, double INTEGER, square INTEGER);"

# Build program
COPY htmlT.go html.gohtml .
RUN go build htmlT.go

FROM golang:1.17

WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/htmlT /usr/src/app/html.gohtml /usr/src/app/htmlT.db .

EXPOSE 8080
ENTRYPOINT ["./htmlT"]

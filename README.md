# Hours API server

Example JSONApi server for the (also example) Hours project.

## Build

Just run `make`. It will generate a command `api-server`.

## Handle database

Get database migration handler:

```
go get github.com/mattes/migrate
```

Then, set up schema to database:

```
migrate -url postgres:///hours\?sslmode=disable -path=db up
```

## Run

Syntax:

```
api-server [-addr <bind-ip:port>] -db <database bind string>
```

Where

- addr: listening address. Default: 0.0.0.0:4000
- database bind string: postgres URL

## API

TBD

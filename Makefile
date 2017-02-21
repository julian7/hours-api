default: api-server

api-server: */*.go
	go build -o api-server ./server

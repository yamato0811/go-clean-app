run:
	DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 DB_NAME=gbd go run .
test:
	go test ./...
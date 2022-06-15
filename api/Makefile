run:
	go run main.go
build:
	go build main.go
image:
	docker build -t goapi .
container:
	docker run -e PORT=8081 -p 8081:8081 goapi
APP_NAME = kursho

all: server client

server:
	@echo "Building Kursho..."
	@go build -o bin/$(APP_NAME) main.go

client:
	@echo "Building Klient (Kursho example)..."
	@go build -o bin/klient klient.go

run: server
	@echo "Running..."
	@./bin/$(APP_NAME)

test: run client
	@echo "Testing..."
	@./bin/klient

clean:
	@echo "Cleaning..."
	@rm -rf bin/*

.PHONY: all server client run test clean

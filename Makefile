# Run main.go
run:
	go run main.go

# Format all Go code
fmt:
	gofmt -w .

# Run static checks (like CI)
check:
	./ci_check.sh

# Build all packages
build:
	go build -v ./...

# Shortcut for go vet
vet:
	go vet ./...

# Build and run using Docker Compose
docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-up-detach:
	docker-compose up -d

docker-down:
	docker-compose down

docker-reload:
	docker-compose down && docker-compose up -d
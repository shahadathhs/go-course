# Makefile for go-course

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

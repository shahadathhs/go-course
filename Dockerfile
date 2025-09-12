FROM golang:latest

WORKDIR /app

# Install git and bash
RUN apt-get update && apt-get install -y git bash

# Set Go bin path
ENV PATH=$PATH:/go/bin

# Install Air and Lefthook
RUN go install github.com/air-verse/air@v1.62.0 && \
    go install github.com/evilmartians/lefthook@latest

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Make scripts executable
RUN chmod +x ./ci_check.sh

# Install Lefthook git hooks
RUN lefthook install

# Default command
CMD ["air", "-c", ".air.toml"]

FROM golang:latest

WORKDIR /app

# Install git, bash, and air
RUN apk add --no-cache git bash && \
    go install github.com/cosmtrek/air@latest

# Install Lefthook
RUN go install github.com/evilmartians/lefthook@latest
ENV PATH=$PATH:/go/bin

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy project files
COPY . .

# Make scripts executable
RUN chmod +x ./ci_check.sh

# Install lefthook git hooks
RUN lefthook install

# Default command
CMD ["air", "-c", ".air.toml"]

#!/bin/bash

set -e

echo
echo "🔧 Building all packages..."
go build -v ./...
echo "✅ Build successful."

echo
echo "🧹 Running go vet..."
go vet ./...
echo "✅ go vet passed."

echo
echo "🧽 Checking code formatting with gofmt..."
UNFORMATTED=$(gofmt -l .)
if [ -n "$UNFORMATTED" ]; then
  echo "❌ The following files are not properly formatted:"
  echo "$UNFORMATTED"
  echo "Run: gofmt -w ."
  exit 1
else
  echo "✅ Code is properly formatted."
fi

echo
echo "🎉 All checks passed!"

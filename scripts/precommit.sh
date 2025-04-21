#!/bin/bash
set -e 

echo "⚙️ Running golangci-lint..."

if ! command -v golangci-lint &> /dev/null; then
  echo "❌ golangci-lint is not installed."
  exit 1
fi

if ! golangci-lint run ./...; then
  echo "❌ Linter has found problems."
  exit 1
fi

echo "✅ The code is clear."

echo "⚙️ Running gofmt..."

if ! command -v gofmt &> /dev/null; then
  echo "❌ gofmt is not installed."
  exit 1
fi

if ! gofmt -w .; then
  echo "❌ gofmt encountered an error."
  exit 1
fi

echo "✅ Code formatted with gofmt."

exit 0
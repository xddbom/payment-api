echo "Running golangci-lint..."

golangci-lint run ./...  

if [ $? -ne 0 ]; then
  echo "❌ Linter has found errors."
  exit 1
fi

echo "✅ The code is clear."
exit 0
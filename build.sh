#!/bin/bash
echo "Building..."
GOOS=linux GOARCH=amd64 go build -o main main.go
chmod +x main
echo "Building completed"

echo "Packing..."
zip deployment.zip main
echo "Packing completed"

echo "Cleaning up"
rm main

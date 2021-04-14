#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/jobs/"

rm ./docker/jobs
mv ./jobs ./docker/

docker build -t jobs:"latest" docker/

docker-compose up -d
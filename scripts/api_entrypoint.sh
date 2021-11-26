#!/bin/bash
set -e

echo "#################### downloading CompileDaemon for api"
# disable go modules to avoid this package from getting into go.mod
# as we only need it locally to watch and rebuild server on change
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

echo "#################### starting deamon"
CompileDaemon --exclude-dir=cmd/crawler --build="go build -o api cmd/api/main.go" --command=./api

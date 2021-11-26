#!/bin/bash
set -e

echo "#################### downloading CompileDaemon for crawler"
# disable go modules to avoid this package from getting into go.mod
# as we only need it locally to watch and rebuild server on change
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

echo "#################### starting deamon"
CompileDaemon --exclude-dir=cmd/api --build="go build -o crawler cmd/crawler/main.go" --command=./crawler

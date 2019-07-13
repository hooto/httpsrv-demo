#!/bin/bash

go build -o bin/demo-server cmd/server/main.go

./bin/demo-server --logtostderr true


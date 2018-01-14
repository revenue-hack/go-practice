#!/bin/sh

go build gopl.io/ch1/fetch
go build main.go
./fetch https://golang.org | ./main
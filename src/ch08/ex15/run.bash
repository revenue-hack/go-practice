#!/bin/bash

go build chat.go
go build netcat3.go
./chat &
# ./netcat3 1
# ./netcat3 2
# ./netcat3 3

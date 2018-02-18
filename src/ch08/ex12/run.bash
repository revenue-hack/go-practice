#!/bin/bash

go build chat.go
go build netcat3.go
./chat &
./netcat3

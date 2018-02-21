#!/bin/bash

go build ftp.go
go build netcat.go
./ftp &
./netcat

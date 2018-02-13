#!/bin/bash
go build netcat3.go
go build reverb2.go
./reverb2 &
./netcat3

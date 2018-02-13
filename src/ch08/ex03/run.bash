#!/bin/bash
go build netcat3.go
go build reverb1.go
./reverb1 &
./netcat3

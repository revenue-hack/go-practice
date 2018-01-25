#!/bin/bash

go build tempflag.go
./tempflag -temp 20C
./tempflag -temp 20F
./tempflag -temp 20K

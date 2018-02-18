#!/bin/bash

go build crawl.go
./crawl -depth=10 https://golang.org

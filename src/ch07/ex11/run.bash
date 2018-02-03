#!/bin/bash

# start server
go run http.go

# localhost:8888/list
# localhost:8888/create?item=hoge&price=200
# localhost:8888/read?item=hoge
# localhost:8888/update?item=hoge&price=10000
# localhost:8888/delete?item=hoge


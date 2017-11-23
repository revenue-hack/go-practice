#!/bin/bash
go build fetchall.go
./fetchall https://www.flickr.com/photos/tags/flicker/ http://www.skincare-univ.com/ https://echo.labstack.com/ https://golang.org/
./fetchall https://www.flickr.com/photos/tags/flicker/ http://www.skincare-univ.com/ https://echo.labstack.com/ https://golang.org/


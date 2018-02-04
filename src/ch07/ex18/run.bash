#!/bin/bash
go build gopl.io/ch1/fetch
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 > input.xml
go build xmltree.go
./xmltree < input.xml > output.html


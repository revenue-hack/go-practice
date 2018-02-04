#!/bin/bash
go build xmlselect.go
go build gopl.io/ch1/fetch
#./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div div h2
echo "class pattern"
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect .toc h2
echo "id pattern"
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect \#dt-parentchild
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect \#dt-extlitentval


package unarchive

import (
	"bufio"
	"errors"
	"os"
)

type Reader interface {
	ReadFiles() []File
}

type format struct {
	name, magic      string
	peekSize, offset int
	decode           func(string) (Reader, error)
}

type File interface {
	Name() string
	FileInfo() os.FileInfo
}

var formats []format

func OpenReader(fileName string) (Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for _, format := range formats {
		b, err := r.Peek(format.peekSize)
		if err != nil {
			return nil, err
		}
		if match(format.magic, b[format.offset:]) {
			return format.decode(fileName)
		}
	}
	return nil, errors.New("format nothing")
}

func match(magicNum string, fileMagicNum []byte) bool {
	return magicNum == string(fileMagicNum)

}

func RegisterFormat(name, magic string, peekSize, offset int, decode func(string) (Reader, error)) {
	formats = append(formats, format{name, magic, peekSize, offset, decode})
}

package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/revenue-hack/go-practice/src/ch10/ex02/unarchive"
)

const tarHeader = "ustar"

func init() {
	unarchive.RegisterFormat("tar", tarHeader, 257+len(tarHeader), 257, Decode)
}

type Reader struct {
	f *os.File
	r *tar.Reader
}

type TarFile struct {
	h *tar.Header
}

func (f *TarFile) FileInfo() os.FileInfo {
	return f.h.FileInfo()
}

func (f *TarFile) Name() string {
	return f.h.Name
}

func (r *Reader) ReadFiles() []unarchive.File {
	var files []unarchive.File
	for {
		h, err := r.r.Next()
		if err == io.EOF {
			r.f.Close()
			break
		}
		f := &TarFile{h}
		files = append(files, f)
	}
	return files
}

func Decode(fileName string) (unarchive.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	r := tar.NewReader(f)
	return &Reader{f, r}, nil
}

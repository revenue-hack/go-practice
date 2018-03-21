package zip

import (
	"os"

	"archive/zip"
	"ch10/ex02/unarchive"
)

const zipHeader = "PK\003\004"

func init() {
	unarchive.RegisterFormat("zip", zipHeader, len(zipHeader), 0, Decode)
}

type Reader struct {
	r *zip.ReadCloser
}

type ZipFile struct {
	f *zip.File
}

func (f *ZipFile) FileInfo() os.FileInfo {
	return f.f.FileInfo()
}

func (f *ZipFile) Name() string {
	return f.f.Name
}

func (r *Reader) ReadFiles() []unarchive.File {
	defer r.r.Close()
	var files []unarchive.File
	for _, f := range r.r.File {
		zf := &ZipFile{f}
		files = append(files, zf)
	}
	return files
}

func Decode(fileName string) (unarchive.Reader, error) {
	f, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}
	return &Reader{f}, nil
}

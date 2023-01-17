package file

import "io"

type File struct {

}

func (f *File) NeedsUpload() bool {
	return true
}

func (f *File) UploadData() (name string, file io.Reader, err error) {

}

func (f *File) SendData() string {

}

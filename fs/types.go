package fs

import (
	"os"
	"time"
)

type File struct {
	name     string
	fullPath string
	data     []byte
}

type FileStat struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	Mode    os.FileMode `json:"mode"`
	ModTime time.Time   `json:"modTime"`
	IsDir   bool        `json:"isDir"`
}

// type FileTree struct {

// }

type FileList []*FileStat

func (f *File) StringData() string {
	return string(f.data)
}

func (f *File) Data() []byte {
	return f.data
}

func NewFile(name string, fullPath string, data []byte) *File {
	return &File{
		name,
		fullPath,
		data,
	}
}

func NewFileStat(fileInfo os.FileInfo) *FileStat {
	return &FileStat{
		fileInfo.Name(),
		fileInfo.Size(),
		fileInfo.Mode(),
		fileInfo.ModTime(),
		fileInfo.IsDir(),
	}
}

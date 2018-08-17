package fs

type File struct {
	name     string
	fullPath string
	data     []byte
}

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

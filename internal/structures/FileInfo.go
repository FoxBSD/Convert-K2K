package structures

type FileData struct {
	FileName  string
	FilePath  string
	Extension string
}

// This function returns the File name
func (f *FileData) GetName() string {
	return f.FileName
}

// This function returns the File path
func (f *FileData) GetPath() string {
	return f.FilePath
}

// This function retuns the File extension
func (f *FileData) GetExt() string {
	return f.Extension
}

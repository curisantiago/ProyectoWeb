package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

type FileStore struct {
	FileName string
}

func NewStore(store Type, filename string) Store {
	switch store {
	case FileType:
		return &FileStore{filename}
	}
	return nil
}
func (fs *FileStore) Write(data interface{}) error {
	filedata, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, filedata, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	filedata, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, &data)
}

package filesystem

import "os"

type Filesystem interface {
	GetCurrentDirectory() (string, error)
	ReadFiles(dir string) ([]os.FileInfo, error)
	GetExtension(file string) string
	RemoveExtension(file string) string
	RemoveFile(file string) error
}

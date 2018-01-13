package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"

	"github.com/pkg/errors"
)

type filesystem struct {
	CurDir string
}

func NewFilesystem() *filesystem {
	return &filesystem{}
}

func (f *filesystem) GetCurrentDirectory() (string, error) {
	if f.CurDir != "" {
		return f.CurDir, nil
	}

	dir, err := filepath.Abs("./")
	if err != nil {
		return "", errors.Wrap(err, "Unable to get the current directory")
	}

	f.CurDir = dir

	return dir, nil
}

func (f *filesystem) ReadFiles(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(f.CurDir + "/" + dir)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to read files from the following directory: %s", dir)
	}

	return files, nil
}

func (f *filesystem) GetExtension(file string) string {
	return filepath.Ext(file)
}

func (f *filesystem) RemoveExtension(file string) string {
	return strings.Replace(file, f.GetExtension(file), "", 1)
}

func (f *filesystem) RemoveFile(file string) error {
	return os.Remove(file)
}

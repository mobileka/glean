package glean

import (
	"io/ioutil"
	"os"
	"strings"

	"fmt"

	"github.com/mobileka/glean/pkg/debug"
	"github.com/mobileka/glean/pkg/filesystem"
	"github.com/pkg/errors"
)

type Arguments struct {
	JpegDir string
	RawDir  string
	JpegExt []string
	RawExt  []string
}

type glean struct {
	dbg    debug.Debugger
	config Config
	fs     filesystem.Filesystem
	argv   Arguments
}

func NewGlean(dbg debug.Debugger, config Config, fs filesystem.Filesystem, argv Arguments) *glean {
	return &glean{dbg: dbg, config: config, fs: fs, argv: argv}
}

func (g *glean) Run() error {
	g.dbg.Logf("Config Path: %s\n", g.config.GetPath())

	confContent, err := ioutil.ReadFile(g.config.GetPath())
	if err != nil {
		return errors.Wrapf(err, "Unable to read the config file: %s", g.config.GetPath())
	}
	g.dbg.Logf("Config Content:\n%s\n", confContent)

	if err = g.config.Unmarshal(confContent); err != nil {
		return errors.Wrapf(err, "Unable to unmarshall the config file: %s", confContent)
	}
	g.dbg.Logf("Unserialized Config Contents: %s\n", g.config)

	g.config.SetJpegDir(g.argv.JpegDir).
		SetRawDir(g.argv.RawDir).
		SetJpegExt(g.argv.JpegExt).
		SetRawExt(g.argv.RawExt)
	g.dbg.Logf("Customised Config:\n%s\n", g.config)

	JPEGs, err := g.fs.ReadFiles(g.config.GetJpegDir())
	if err != nil {
		return err
	}

	JPEGMap := g.filesToMap(JPEGs)

	RAWs, err := g.fs.ReadFiles(g.config.GetRawDir())
	if err != nil {
		return err
	}

	if err := g.glean(RAWs, JPEGMap); err != nil {
		return err
	}

	return nil
}

func (g *glean) filesToMap(files []os.FileInfo) map[string]bool {
	var result = map[string]bool{}
	for _, f := range files {
		fullPath := g.config.GetRawDir() + "/" + f.Name()
		ext := g.fs.GetExtension(fullPath)

		if !f.IsDir() && isRelevantExtension(ext, g.config.GetJpegExt()) {
			pathWithNoExtension := g.fs.RemoveExtension(fullPath)
			result[pathWithNoExtension] = true
		}
	}

	return result
}

func (g *glean) glean(RAWs []os.FileInfo, JPEGs map[string]bool) error {
	var failedToRemove []string
	var failedToRemoveErr error

	for _, f := range RAWs {
		fullPath := g.config.GetRawDir() + "/" + f.Name()
		ext := g.fs.GetExtension(fullPath)

		if !f.IsDir() && isRelevantExtension(ext, g.config.GetRawExt()) {
			pathWithNoExtension := strings.Replace(fullPath, ext, "", 1)

			if _, ok := JPEGs[pathWithNoExtension]; !ok {
				fmt.Println("Gleaning " + fullPath)
				if err := g.fs.RemoveFile(fullPath); err != nil {
					failedToRemoveErr = err
					failedToRemove = append(failedToRemove, fullPath)
				}
			}
		}
	}

	if failedToRemoveErr != nil {
		return errors.Wrapf(failedToRemoveErr, "Unable to remove one or several files: %s", failedToRemove)
	}

	return nil
}

func isRelevantExtension(str string, slice []string) bool {
	for _, el := range slice {
		if str == el {
			return true
		}
	}
	return false
}

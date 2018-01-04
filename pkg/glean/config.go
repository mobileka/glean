package glean

import "gopkg.in/yaml.v2"

type config struct {
	path    string   `yaml:"-"`
	JpegDir string   `yaml:"jpeg_dir"`
	JpegExt []string `yaml:"jpeg_ext"`
	RawDir  string   `yaml:"raw_dir"`
	RawExt  []string `yaml:"raw_ext"`
}

type Config interface {
	GetPath() string
	SetJpegDir(val string) *config
	SetRawDir(val string) *config
	SetJpegExt(val []string) *config
	SetRawExt(val []string) *config
	GetJpegDir() string
	GetRawDir() string
	GetJpegExt() []string
	GetRawExt() []string
	Unmarshal(content []byte) error
}

func NewConfig(path string, curDir string) *config {
	if path == "" {
		path = "/glean.yaml"
	}

	path = curDir + path

	return &config{path: path}
}

func (c *config) GetPath() string {
	return c.path
}

func (c *config) SetJpegDir(val string) *config {
	if val != "" {
		c.JpegDir = val
	}

	return c
}

func (c *config) SetRawDir(val string) *config {
	if val != "" {
		c.RawDir = val
	}

	return c
}

func (c *config) SetJpegExt(val []string) *config {
	if len(val) > 0 {
		c.JpegExt = val
	}

	return c
}

func (c *config) SetRawExt(val []string) *config {
	if len(val) > 0 {
		c.RawExt = val
	}

	return c
}

func (c *config) GetJpegDir() string {
	return c.JpegDir
}

func (c *config) GetRawDir() string {
	return c.RawDir
}

func (c *config) GetJpegExt() []string {
	return c.JpegExt
}

func (c *config) GetRawExt() []string {
	return c.RawExt
}

func (c *config) Unmarshal(content []byte) error {
	return yaml.Unmarshal(content, &c)
}

package glean

import "gopkg.in/yaml.v2"

type config struct {
	path      string   `yaml:"-"`
	RefDir    string   `yaml:"ref_dir"`
	RefExt    []string `yaml:"ref_ext"`
	TargetDir string   `yaml:"target_dir"`
	TargetExt []string `yaml:"target_ext"`
}

type Config interface {
	GetPath() string
	SetRefDir(val string) *config
	SetTargetDir(val string) *config
	SetRefExt(val []string) *config
	SetTargetExt(val []string) *config
	GetRefDir() string
	GetTargetDir() string
	GetRefExt() []string
	GetTargetExt() []string
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

func (c *config) SetRefDir(val string) *config {
	if val != "" {
		c.RefDir = val
	}

	return c
}

func (c *config) SetTargetDir(val string) *config {
	if val != "" {
		c.TargetDir = val
	}

	return c
}

func (c *config) SetRefExt(val []string) *config {
	if len(val) > 0 {
		c.RefExt = val
	}

	return c
}

func (c *config) SetTargetExt(val []string) *config {
	if len(val) > 0 {
		c.TargetExt = val
	}

	return c
}

func (c *config) GetRefDir() string {
	return c.RefDir
}

func (c *config) GetTargetDir() string {
	return c.TargetDir
}

func (c *config) GetRefExt() []string {
	return c.RefExt
}

func (c *config) GetTargetExt() []string {
	return c.TargetExt
}

func (c *config) Unmarshal(content []byte) error {
	return yaml.Unmarshal(content, &c)
}

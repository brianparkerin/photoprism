package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/photoprism/photoprism/pkg/fs"
	"gopkg.in/yaml.v2"
)

type MapsSettings struct {
	Animate int    `json:"animate" yaml:"animate"`
	Style   string `json:"style" yaml:"style"`
}

// Settings contains Web UI settings
type Settings struct {
	Theme    string       `json:"theme" yaml:"theme"`
	Language string       `json:"language" yaml:"language"`
	Maps     MapsSettings `json:"maps" yaml:"maps"`
}

// NewSettings returns a empty Settings
func NewSettings() *Settings {
	return &Settings{
		Theme:    "default",
		Language: "en",
		Maps: MapsSettings{
			Animate: 0,
			Style:   "streets",
		},
	}
}

// Load uses a yaml config file to initiate the configuration entity.
func (s *Settings) Load(fileName string) error {
	if !fs.FileExists(fileName) {
		return fmt.Errorf("settings file not found: \"%s\"", fileName)
	}

	yamlConfig, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlConfig, s)
}

// Save uses a yaml config file to initiate the configuration entity.
func (s *Settings) Save(fileName string) error {
	data, err := yaml.Marshal(s)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, data, os.ModePerm)
}

// Settings returns the current user settings.
func (c *Config) Settings() *Settings {
	s := NewSettings()
	p := c.SettingsFile()

	if err := s.Load(p); err != nil {
		log.Error(err)
	}

	return s
}

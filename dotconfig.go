package dotconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

var (
	// ErrConfigNotFound is used when the config cannot be found during the Loading.
	// This can happen if the configuration never got saved.
	ErrConfigNotFound = errors.New("config not found")
)

// Save serializes and saves the v interface into a YAML file at
// ~/.config/{appName}/config.yml
func Save(appName string, v interface{}) error {
	home, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "unable to get home directory")
	}

	cfgDir := filepath.Join(home, ".config", appName)
	if err := os.MkdirAll(cfgDir, 0700); err != nil {
		return errors.Wrap(err, "unable to create config directory")
	}

	cfgFile := filepath.Join(cfgDir, "config.yml")

	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cfgFile, data, 0600)
}

// Load loads the configuration from ~/.config/{appName}/config.yaml and deserialises it to v.
func Load(appName string, v interface{}) error {
	home, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "unable to get home directory")
	}

	cfgDir := filepath.Join(home, ".config", appName)
	cfgFile := filepath.Join(cfgDir, "config.yml")
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		return ErrConfigNotFound
	} else if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, v)
}

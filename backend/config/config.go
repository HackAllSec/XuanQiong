package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
	"xuanqiong/backend/types"
)

var (
	configFile = "config.yaml"
	Config     types.Config
	Version    = "v1.1.0"
)

func init() {
	resolvedConfigFile, err := resolveConfigFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	data, err := os.ReadFile(resolvedConfigFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	if isTestBinary() {
		Config.Database.Type = "sqlite"
		Config.Database.Connection.File = "file::memory:?cache=shared"
		Config.Log.Level = "silent"
	}
}

func resolveConfigFile(name string) (string, error) {
	if filepath.IsAbs(name) {
		if _, err := os.Stat(name); err != nil {
			return "", err
		}
		return name, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(wd, name)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			break
		}
		wd = parent
	}
	return "", os.ErrNotExist
}

func isTestBinary() bool {
	return len(os.Args) > 0 && strings.HasSuffix(os.Args[0], ".test")
}

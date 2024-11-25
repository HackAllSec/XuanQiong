package config

import (
    "os"
    "log"

    "gopkg.in/yaml.v2"
    "xuanqiong/backend/types"
)

var (
    configFile = "config.yaml"
    Config types.Config
    Version = "v1.0.7"
)

func init() {
    data, err := os.ReadFile(configFile)
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    err = yaml.Unmarshal(data, &Config)
    if err != nil {
        log.Fatalf("Error unmarshalling config: %v", err)
    }
}
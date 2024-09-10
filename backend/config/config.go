package config

import (
    "os"
    "log"

    "gopkg.in/yaml.v2"
)

var (
    Config ConfigStruct
    Version = "v1.0.0"
)

// ConfigStruct 结构体定义
type ConfigStruct struct {
    Database struct {
        Type       string `yaml:"type"`
        Connection struct {
            Host     string `yaml:"host"`
            Port     int64    `yaml:"port"`
            User     string `yaml:"user"`
            Password string `yaml:"password"`
            Name     string `yaml:"name"`
            Charset  string `yaml:"charset"`
            File     string `yaml:"file"`
        } `yaml:"connection"`
    } `yaml:"database"`
    JWT struct {
        Secret     string `yaml:"secret"`
        ExpiresIn  int64 `yaml:"expires_in"`
    } `yaml:"jwt"`
    Server struct {
        Mode          string `yaml:"mode"`
        Host          string `yaml:"host"`
        Port          int64 `yaml:"port"`
        ReadTimeout   int64 `yaml:"read_timeout"`
        WriteTimeout  int64 `yaml:"write_timeout"`
    } `yaml:"server"`
    Log struct {
        Level string `yaml:"level"`
        File  string `yaml:"file"`
    } `yaml:"log"`
    Login struct {
        MaxAttempts    int64    `yaml:"max_attempts"`
        LockoutDuration int64 `yaml:"lockout_duration"`
    } `yaml:"login"`
}

func init() {
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    err = yaml.Unmarshal(data, &Config)
    if err != nil {
        log.Fatalf("Error unmarshalling config: %v", err)
    }
}

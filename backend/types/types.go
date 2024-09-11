package types

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

// Config 结构体定义
type Config struct {
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

// User 用户表
type User struct {
    ID         uint64      `gorm:"primaryKey"`
    Username   string      `gorm:"unique"`
    Password   string
    Role       int64
    CreateTime time.Time
    Status     int64
    Token      string
}

// Vulnerability 漏洞表
type Vulnerability struct {
    ID             string    `gorm:"primaryKey" json:"id"`
    CVE            string    `json:"cve"`
    CNNVD          string    `json:"cnnvd"`
    CNVD           string    `json:"cnvd"`
    VulnName       string    `json:"vuln_name"`
    VulnType       string    `json:"vuln_type"`
    VulnLevel      string    `json:"vuln_level"`
    Description    string    `json:"description"`
    AffectedProduct string   `json:"affected_product"`
    Poc            string    `json:"poc"`
    Exp            string    `json:"exp"`
    RepairSuggestion string  `json:"repair_suggestion"`
    Submit         string    `json:"submit"`
    CreateTime     time.Time `json:"create_time"`
    UpdateTime     time.Time `json:"update_time"`
}

// Lockip 锁定IP表
type Lockip struct {
    ID         uint64    `gorm:"primaryKey"`
    ClientIP   string
    LockoutUntil  *time.Time
}

type Jwt struct {
    Role       uint64    `json:"role"`
    Username   string    `json:"username"`
    jwt.StandardClaims
}

type LoginData struct {
    Username string    `json:"username"`
    Password string    `json:"password"`
}

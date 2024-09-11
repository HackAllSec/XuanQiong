package types

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

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

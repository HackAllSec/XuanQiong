package types

import (
    "time"
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
        Mode           string `yaml:"mode"`
        Host           string `yaml:"host"`
        Port           int64 `yaml:"port"`
        ReadTimeout    int64 `yaml:"read_timeout"`
        WriteTimeout   int64 `yaml:"write_timeout"`
        FrontendPath   string `yaml:"frontend_path"`
        StaticUrl      string `yaml:"static_url"`
        AdminPath      string `yaml:"admin_path"`
        AdminStaticUrl string `yaml:"admin_static_url"`
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
    Avatar     string
    Email      string
    Phone      string
    Ranking    int64
    Role       int64
    CreateTime time.Time
    UpdateTime time.Time
    Status     int64
    Token      string
}

// Vulnerability 漏洞表
type Vulnerability struct {
    ID                       string    `gorm:"primaryKey" json:"id"`
    UserID                   uint64    `json:"user_id"`
    CVE                      string    `json:"cve"`
    NVD                      string    `json:"nvd"`
    EDB                      string    `json:"edb"`
    CNNVD                    string    `json:"cnnvd"`
    CNVD                     string    `json:"cnvd"`
    VulnName                 string    `json:"vuln_name"`
    VulnType                 string    `json:"vuln_type"`
    VulnLevel                string    `json:"vuln_level"`
    Description              string    `json:"description"`
    AffectedProduct          string    `json:"affected_product"`
    AffectedProductVersion   string    `json:"affected_product_version"`
    FofaQuery                string    `json:"fofa_query"`
    ZoomEyeQuery             string    `json:"zoomeye_query"`
    QuakeQuery               string    `json:"quake_query"`
    HunterQuery              string    `json:"hunter_query"`
    GoogleQuery              string    `json:"google_query"`
    ShodanQuery              string    `json:"shodan_query"`
    CensysQuery              string    `json:"censys_query"`
    GreynoiseQuery           string    `json:"greynoise_query"`
    Poc                      string    `json:"poc"`
    PocType                  string    `json:"poc_type"`
    Exp                      string    `json:"exp"`
    ExpType                  string    `json:"exp_type"`
    RepairSuggestion         string    `json:"repair_suggestion"`
    AttachmentID             string    `json:"attachment_id"`
    AttachmentName           string    `json:"attachment_name"`
    Submitter                string    `json:"submitter"`
    IsPublic                 bool      `json:"is_public"`
    Status                   int64     `json:"status"`
    CreateTime               time.Time `json:"create_time"`
    UpdateTime               time.Time `json:"update_time"`
}

// Lockip 锁定IP表
type Lockip struct {
    ID             uint64    `gorm:"primaryKey"`
    ClientIP       string
    Status         int64
    CreateTime     time.Time
    LockoutUntil   *time.Time
}

// 附件表
type Attachment struct {
    ID          string    `gorm:"primaryKey"`
    UserID      uint64
    Name        string
    Type        string
    Data        []byte
    Status      int64
    CreateTime  time.Time
    UpdateTime  time.Time
}

type LoginData struct {
    Username string    `json:"username"`
    Password string    `json:"password"`
}

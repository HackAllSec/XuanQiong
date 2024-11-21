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
}

// 系统配置表
type XqSystemConfig struct {
    ID                  uint64    `gorm:"primaryKey" json:"id"`
    UserRegister        bool      `json:"user_register"`
    UserDisplay         string    `json:"user_display"`
    MaxAttempts         int64     `json:"max_attempts"`
    LockoutDuration     int64     `json:"lockout_duration"`
    CreateTime          time.Time `json:"create_time"`
    UpdateTime          time.Time `json:"update_time"`
}

// Jwt配置表
type XqJwtConfig struct {
    ID                  uint64    `gorm:"primaryKey" json:"id"`
    JwtSecret           string    `json:"jwt_secret"`
    JwtExpires          int64     `json:"jwt_expires"`
    CreateTime          time.Time `json:"create_time"`
    UpdateTime          time.Time `json:"update_time"`
}

// 邮箱配置表
type XqEmailConfig struct {
    ID                  uint64    `gorm:"primaryKey" json:"id"`
    EmailHost           string    `json:"email_host"`
    EmailPort           int64     `json:"email_port"`
    EmailUser           string    `json:"email_user"`
    EmailPassword       string    `json:"email_password"`
    EmailSender         string    `json:"email_sender"`
    CreateTime          time.Time `json:"create_time"`
    UpdateTime          time.Time `json:"update_time"`
}

// 通知配置表
type XqNoticeConfig struct {
    ID              uint64      `gorm:"primaryKey" json:"id"`
    Type            int64      `json:"type"`
    Secret          string      `json:"secret"`
    Webhook         string      `json:"webhook"`
    CreateTime      time.Time   `json:"create_time"`
    UpdateTime      time.Time   `json:"update_time"`
}

// User 用户表
type XqUser struct {
    ID         uint64      `gorm:"primaryKey" json:"id"`
    Username   string      `gorm:"unique" json:"username"`
    Password   string      `json:"password"`
    Avatar     string      `json:"avatar"`
    Email      string      `json:"email"`
    Phone      string      `json:"phone"`
    Ranking    int64       `json:"ranking"`
    Role       int64       `json:"role"`
    Status     int64       `json:"status"`
    Token      string      `json:"token"`
    CreateTime time.Time   `json:"create_time"`
    UpdateTime time.Time   `json:"update_time"`
}

// 漏洞类型表
type XqVulnType struct {
    ID          uint64    `gorm:"primaryKey" json:"id"`
    Name       string     `json:"name"`
    CreateTime time.Time  `json:"create_time"`
    UpdateTime time.Time  `json:"update_time"`
}

// Vulnerability 漏洞表
type XqVulnerability struct {
    ID                       string    `gorm:"primaryKey" json:"id"`
    UserID                   uint64    `json:"user_id"`
    CVE                      string    `json:"cve"`
    NVD                      string    `json:"nvd"`
    EDB                      string    `json:"edb"`
    CNNVD                    string    `json:"cnnvd"`
    CNVD                     string    `json:"cnvd"`
    VulnName                 string    `json:"vuln_name"`
    VulnTypeID               uint64    `json:"vuln_type_id"`
    VulnType                 string    `json:"vuln_type"`
    VulnLevel                string    `json:"vuln_level"`
    CVSS                     float64   `json:"cvss"`
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
    ReviewComments           string    `json:"review_comments"`
    CreateTime               time.Time `json:"create_time"`
    UpdateTime               time.Time `json:"update_time"`
}

// Lockip 锁定IP表
type XqLockip struct {
    ID             uint64    `gorm:"primaryKey"`
    ClientIP       string
    LockoutUntil   *time.Time
    Status         int64
    CreateTime     time.Time
    UpdateTime     time.Time
}

// 附件表
type XqAttachment struct {
    ID          string    `gorm:"primaryKey"`
    UserID      uint64
    Name        string
    Type        string
    Data        []byte
    Status      int64
    CreateTime  time.Time
    UpdateTime  time.Time
}

// 用户ranking明细表
type XqRankingDetail struct {
    ID          uint64    `gorm:"primaryKey"`
    UserID      uint64
    VulnID      string
    Ranking     int64
    CreateTime  time.Time
    UpdateTime  time.Time
}

// 评分规则表
type XqScoreRule struct {
    ID              uint64    `gorm:"primaryKey"`
    Type            int64
    Rule            string
    Score           float64
    Coefficient     float64
    CreateTime      time.Time
    UpdateTime      time.Time
}

// 验证码表
type XqVerifyCode struct {
    ID             uint64    `gorm:"primaryKey"`
    Email          string
    Code           string
    CreateTime     time.Time
    UpdateTime     time.Time
    ExpiredTime     time.Time
}

type SystemConfigData struct {
    EmailConfig    XqEmailConfig    `json:"emailconf"`
    JwtConfig      XqJwtConfig      `json:"jwtconf"`
    NoticeConfig   XqNoticeConfig   `json:"noticeconf"`
    SysConfig      XqSystemConfig   `json:"sysconf"`
}
package models

import (
    "fmt"
    "log"
    "time"

    "github.com/natefinch/lumberjack"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    //"gorm.io/driver/sqlite"
    "github.com/glebarez/sqlite"
    "gorm.io/driver/sqlserver"
    "xuanqiong/config"
    "xuanqiong/types"
    "xuanqiong/utils"
)

var (
    db *gorm.DB
)

func init() {
    lumberjackLogger := &lumberjack.Logger{
        Filename:   config.Config.Log.File,
        MaxSize:    10, // megabytes
        MaxBackups: 7,  // number of old log files to retain
        MaxAge:     30, // days
        Compress:   true, // whether to compress the log files
    }
    gormLogger := logger.New(
        log.New(lumberjackLogger, "\r\n", log.LstdFlags),
        logger.Config{
            IgnoreRecordNotFoundError: true,  // 忽略记录未找到错误
        },
    )

    switch config.Config.Log.Level {
        case "silent":
            gormLogger = gormLogger.LogMode(logger.Silent)
        case "error":
            gormLogger = gormLogger.LogMode(logger.Error)
        case "warn":
            gormLogger = gormLogger.LogMode(logger.Warn)
        case "info":
            gormLogger = gormLogger.LogMode(logger.Info)
    }
    dsn := generateDSN(config.Config)
    var err error
    var user types.XqUser
    switch config.Config.Database.Type {
    case "mysql":
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
    case "postgres":
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
    case "sqlite":
        db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: gormLogger})
    case "sqlserver":
        db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: gormLogger})
    default:
        log.Fatalf("Unsupported database type: %s", config.Config.Database.Type)
    }

    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    res := db.Raw("SELECT * FROM xq_users WHERE username = ?", "admin").Scan(&user)
    if res.Error != nil {
        log.Println("Initializing Database...")
        db.AutoMigrate(&types.XqSystemConfig{}, &types.XqJwtConfig{}, &types.XqEmailConfig{}, &types.XqNoticeConfig{}, &types.XqUser{},
            &types.XqVulnType{}, &types.XqVulnerability{}, &types.XqLockip{}, &types.XqAttachment{},
            &types.XqRankingDetail{}, &types.XqScoreRule{})
        res = db.First(&types.XqSystemConfig{})
        if res.RowsAffected == 0 {
            db.Create(&types.XqSystemConfig{UserRegister: false, UserDisplay: "佚名", MaxAttempts: 5, LockoutDuration: 3600, CreateTime: time.Now()})
        }
        res = db.First(&types.XqJwtConfig{})
        if res.RowsAffected == 0 {
            db.Create(&types.XqJwtConfig{JwtSecret: utils.GenerateRandomJwtSecret(), JwtExpires: 3600, CreateTime: time.Now()})
        }
        res = db.First(&types.XqVulnType{})
        if res.RowsAffected == 0 {
            db.Create(&types.XqVulnType{Name: "信息泄露", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "远程代码执行", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "反序列化", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "SQL注入", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "跨站脚本攻击", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "本地文件包含", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "远程文件包含", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "跨站请求伪造", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "服务器端请求伪造", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "任意文件上传", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "任意文件读取/下载", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "拒绝服务", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "暴力破解", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "弱口令", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "XML外部实体注入", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "未授权访问", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "水平越权", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "垂直越权", CreateTime: time.Now()})
            db.Create(&types.XqVulnType{Name: "其他", CreateTime: time.Now()})
        }
        res = db.First(&types.XqScoreRule{})
        if res.RowsAffected == 0 {
            db.Create(&types.XqScoreRule{Type: 1, Rule: "Xray、Nuclei、Goby等完整Poc，误报低", Score: 20, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 1, Rule: "Xray、Nuclei、Goby等完整Poc，误报较高", Score: 10, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 1, Rule: "仅包含Payload或无法工具化的Poc", Score: 5, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 2, Rule: "Xray、Nuclei、Goby等完整Exp，误报低", Score: 30, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 2, Rule: "Xray、Nuclei、Goby等完整Exp，误报较高", Score: 15, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 2, Rule: "仅包含Payload或无法工具化的Exp", Score: 5, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数大于 5000", Score: 30, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数介于 1000 到 5000", Score: 20, Coefficient: 1.0, CreateTime: time.Now()})
            db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数小于 1000", Score: 10, Coefficient: 1.0, CreateTime: time.Now()})
        }
        password, err := utils.GenerateRandomPassword(12)
        if err != nil {
            log.Fatalf("Error generating password: %v", err)
        }
        err = CreateUser("admin", password, "", "", 1)
        if err == nil {
            log.Println("Database initialized successfully!")
            log.Println("Successfully created administrator account:")
            log.Println("Username: admin")
            log.Println("Password:", password)
        }
    } else {
        if res.RowsAffected == 0 {
            password, err := utils.GenerateRandomPassword(12)
            if err != nil {
                log.Fatalf("Error generating password: %v", err)
            }
            err = CreateUser("admin", password, "", "", 1)
            if err == nil {
                log.Println("Successfully created administrator account:")
                log.Println("Username: admin")
                log.Println("Password:", password)
            }
        }
    }
}

// generateDSN 根据配置生成 DSN
func generateDSN(config types.Config) string {
    switch config.Database.Type {
    case "mysql":
        return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
            config.Database.Connection.User,
            config.Database.Connection.Password,
            config.Database.Connection.Host,
            config.Database.Connection.Port,
            config.Database.Connection.Name,
            config.Database.Connection.Charset,
        )
    case "postgres":
        return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable",
            config.Database.Connection.User,
            config.Database.Connection.Password,
            config.Database.Connection.Name,
            config.Database.Connection.Port,
        )
    case "sqlite":
        return config.Database.Connection.File // SQLite 使用文件路径
    case "sqlserver":
        return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
            config.Database.Connection.User,
            config.Database.Connection.Password,
            config.Database.Connection.Host,
            config.Database.Connection.Port,
            config.Database.Connection.Name,
        )
    default:
        return ""
    }
}

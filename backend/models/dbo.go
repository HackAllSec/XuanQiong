package models

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/natefinch/lumberjack"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/sqlserver"
    "xuanqiong/backend/config"
    "xuanqiong/backend/types"
    "xuanqiong/backend/utils"
)

var (
    db *gorm.DB
)

func initDatabase() {
    log.Println("Initializing Database...")
    db.AutoMigrate(&types.XqSystemConfig{}, &types.XqJwtConfig{}, &types.XqEmailConfig{}, &types.XqNoticeConfig{}, &types.XqUser{},
        &types.XqVulnType{}, &types.XqVulnerability{}, &types.XqLockip{}, &types.XqAttachment{},
        &types.XqRankingDetail{}, &types.XqScoreRule{}, &types.XqVerifyCode{})
    res := db.First(&types.XqSystemConfig{})
    if res.RowsAffected == 0 {
        db.Create(&types.XqSystemConfig{UserRegister: false, UserDisplay: "佚名", MaxAttempts: 5, LockoutDuration: 3600, CreateTime: time.Now(), UpdateTime: time.Now()})
    }
    res = db.First(&types.XqJwtConfig{})
    if res.RowsAffected == 0 {
        db.Create(&types.XqJwtConfig{JwtSecret: utils.GenerateRandomJwtSecret(), JwtExpires: 3600, CreateTime: time.Now(), UpdateTime: time.Now()})
    }
    res = db.First(&types.XqVulnType{})
    if res.RowsAffected == 0 {
        db.Create(&types.XqVulnType{Name: "信息泄露", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "远程代码执行", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "反序列化", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "SQL注入", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "命令注入", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "跨站脚本攻击", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "本地文件包含", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "远程文件包含", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "跨站请求伪造", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "服务器端请求伪造", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "任意文件上传", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "任意文件读取/下载", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "拒绝服务", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "权限提升", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "暴力破解", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "弱口令", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "XML外部实体注入", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "未授权访问", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "水平越权", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "垂直越权", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "登录绕过", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "任意密码重置", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "负值反冲", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "短信轰炸", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "邮件轰炸", CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqVulnType{Name: "其他", CreateTime: time.Now(), UpdateTime: time.Now()})
    }
    res = db.First(&types.XqScoreRule{})
    if res.RowsAffected == 0 {
        db.Create(&types.XqScoreRule{Type: 1, Rule: "Xray、Nuclei、Goby等完整Poc，误报低", Score: 20, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 1, Rule: "Xray、Nuclei、Goby等完整Poc，误报较高", Score: 10, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 1, Rule: "仅包含Payload或无法工具化的Poc", Score: 5, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 2, Rule: "Xray、Nuclei、Goby等完整Exp，误报低", Score: 30, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 2, Rule: "Xray、Nuclei、Goby等完整Exp，误报较高", Score: 15, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 2, Rule: "仅包含Payload或无法工具化的Exp", Score: 5, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数大于 5000", Score: 30, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数介于 1000 到 5000", Score: 20, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
        db.Create(&types.XqScoreRule{Type: 3, Rule: "互联网资产数小于 1000", Score: 10, Coefficient: 1.0, CreateTime: time.Now(), UpdateTime: time.Now()})
    }
    log.Println("Database initialized successfully!")
    log.Println("Created administrator account:")
    initAdminPassword()
}

func initAdminPassword() {
    password, err := utils.GenerateRandomChars(12, 5)
    if err != nil {
        log.Fatalf("Error generating password: %v", err)
    }
    err = CreateUser("admin", password, "", "", 1)
    if err == nil {
        log.Println("Username: admin")
        log.Println("Password:", password)
    }
}

func checkTablesExist(db *gorm.DB, tables []interface{}) bool {
    for _, table := range tables {
        if !db.Migrator().HasTable(table) {
            return false
        }
    }
    return true
}

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
    switch config.Config.Database.Type {
    case "mysql":
        // 检测和创建数据库
        dbSQL, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.Database.Connection.User, config.Config.Database.Connection.Password, config.Config.Database.Connection.Host, config.Config.Database.Connection.Port))
        if err != nil {
            log.Fatalf("Error opening database connection: %v", err)
        }
        defer dbSQL.Close()

        // 检查数据库是否存在
        var count int
        err = dbSQL.QueryRow("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", config.Config.Database.Connection.Name).Scan(&count)
        if err != nil {
            log.Fatalf("Error checking database existence: %v", err)
        }

        if count == 0 {
            // 数据库不存在，创建数据库
            _, err = dbSQL.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", config.Config.Database.Connection.Name))
            if err != nil {
                log.Fatalf("Error creating database: %v", err)
            }
        }
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
    case "postgres":
        dbSQL, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%d/", config.Config.Database.Connection.User, config.Config.Database.Connection.Password, config.Config.Database.Connection.Host, config.Config.Database.Connection.Port))
        if err != nil {
            log.Fatalf("Error opening database connection: %v", err)
        }
        defer dbSQL.Close()

        var count int
        err = dbSQL.QueryRow("SELECT COUNT(*) FROM pg_database WHERE datname = $1", config.Config.Database.Connection.Name).Scan(&count)
        if err != nil {
            log.Fatalf("Error checking database existence: %v", err)
        }

        if count == 0 {
            _, err = dbSQL.Exec(fmt.Sprintf("CREATE DATABASE %s", config.Config.Database.Connection.Name))
            if err != nil {
                log.Fatalf("Error creating database: %v", err)
            }
        }
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
    case "sqlite":
        var err error
        db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: gormLogger})
        if err != nil {
            log.Fatalf("Error connecting to database: %v", err)
        }
    case "sqlserver":
        dbSQL, err := sql.Open("sqlserver", fmt.Sprintf("server=%s,%d;user id=%s;password=%s;", config.Config.Database.Connection.Host, config.Config.Database.Connection.Port, config.Config.Database.Connection.User, config.Config.Database.Connection.Password))
        if err != nil {
            log.Fatalf("Error opening database connection: %v", err)
        }
        defer dbSQL.Close()

        var count int
        err = dbSQL.QueryRow("SELECT COUNT(*) FROM sys.databases WHERE name = @name", sql.Named("name", config.Config.Database.Connection.Name)).Scan(&count)
        if err != nil {
            log.Fatalf("Error checking database existence: %v", err)
        }

        if count == 0 {
            _, err = dbSQL.Exec(fmt.Sprintf("CREATE DATABASE [%s]", config.Config.Database.Connection.Name))
            if err != nil {
                log.Fatalf("Error creating database: %v", err)
            }
        }
        db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: gormLogger})
    default:
        log.Fatalf("Unsupported database type: %s", config.Config.Database.Type)
    }
    tables := []interface{}{
        &types.XqSystemConfig{},
        &types.XqJwtConfig{},
        &types.XqEmailConfig{},
        &types.XqNoticeConfig{},
        &types.XqUser{},
        &types.XqAttachment{},
        &types.XqVulnerability{},
        &types.XqScoreRule{},
        &types.XqVulnType{},
        &types.XqRankingDetail{},
        &types.XqVerifyCode{},
        &types.XqLockip{},
    }
    allTablesExist := checkTablesExist(db, tables)

    if !allTablesExist {
        initDatabase()
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

package models

import (
    "fmt"
    "log"

	"github.com/natefinch/lumberjack"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
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
            LogLevel: logger.Info,
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
    }
    dsn := generateDSN(config.Config)
    var err error
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

    err = db.Where("username = ?", "admin").First(&user).Error
    if err != nil {
        log.Println("Initializing Database...")
        db.AutoMigrate(&types.User{}, &types.Vulnerability{}, &types.Lockip{})
        password, err := utils.GenerateRandomPassword(12)
        if err != nil {
            log.Fatalf("Error generating password: %v", err)
        }
        err = CreateUser("admin", password, 1)
        if err == nil {
            log.Println("Database initialized successfully!")
            log.Println("Successfully created administrator account:")
            log.Println("Username: admin")
            log.Println("Password:", password)
        }
    }
}

// generateDSN 根据配置生成 DSN
func generateDSN(config config.ConfigStruct) string {
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

package models

import (
    "time"

    "xuanqiong/backend/types"
    "xuanqiong/backend/utils"
)

// 获取系统配置
func GetSystemConfig() (types.XqSystemConfig, types.XqEmailConfig, types.XqJwtConfig, types.XqNoticeConfig) {
    var LoginPolicy types.XqSystemConfig
    var emailConfig types.XqEmailConfig
    var jwtConfig types.XqJwtConfig
    var noticeConfig types.XqNoticeConfig
    
    db.First(&LoginPolicy)
    db.First(&emailConfig)
    db.First(&jwtConfig)
    db.First(&noticeConfig)
    
    return LoginPolicy, emailConfig, jwtConfig, noticeConfig
}

// 更新系统配置
func UpdateSystemConfig(configData types.SystemConfigData) error {
    var err error
    var sysConf types.XqSystemConfig
    var emailConfig types.XqEmailConfig
    var jwtConfig types.XqJwtConfig
    var noticeConfig types.XqNoticeConfig
    configData.SysConfig.UpdateTime = time.Now()
    
    configData.JwtConfig.UpdateTime = time.Now()
    
    db.First(&sysConf)
    err = db.Model(&sysConf).Updates(configData.SysConfig).Error
    
    err = db.First(&emailConfig).Error
    if err == nil {
        configData.EmailConfig.UpdateTime = time.Now()
        err = db.Model(&emailConfig).Updates(configData.EmailConfig).Error
    } else {
        configData.EmailConfig.CreateTime = time.Now()
        db.Create(&configData.EmailConfig)
    }
    
    db.First(&jwtConfig)
    err = db.Model(&jwtConfig).Updates(configData.JwtConfig).Error
    
    err = db.First(&noticeConfig).Error
    if err == nil {
        configData.NoticeConfig.UpdateTime = time.Now()
        err = db.Model(&noticeConfig).Updates(configData.NoticeConfig).Error
    } else {
        configData.NoticeConfig.CreateTime = time.Now()
        db.Create(&configData.NoticeConfig)
    }
    return err
}

// 检查IP是否被锁定
func IsLocked(ip string) bool {
    var lockip types.XqLockip
    if res := db.Raw("SELECT * FROM xq_lockips WHERE client_ip = ? AND status = 1", ip).Scan(&lockip).RowsAffected; res == 1 {
        if lockip.LockoutUntil != nil && time.Now().Before(*lockip.LockoutUntil) {
            return true
        }
    }
    return false
}

// 锁定IP地址
func LockIP(ip string, duration int64) {
    var lockip types.XqLockip
    db.Raw("SELECT * FROM xq_lockips WHERE client_ip = ?", ip).Scan(&lockip)
    lockoutUntil := time.Now().Add(time.Duration(duration) * time.Second)
    lockip.CreateTime = time.Now()
    lockip.ClientIP = ip
    lockip.Status = 1
    lockip.LockoutUntil = &lockoutUntil
    db.Save(&lockip)
}

func getCurrentQuarter() (time.Time, time.Time) {
    now := time.Now()
    currentMonth := now.Month()
    quarter := (currentMonth - 1) / 3 + 1

    // 计算季度的开始月份
    startMonth := time.Month((quarter - 1) * 3 + 1)
    // 计算季度的结束月份
    endMonth := startMonth + 2

    var startTime time.Time
    var endTime time.Time
    if quarter == 1 {
        // 如果是第一季度，结束月份是3月
        endTime = time.Date(now.Year(), 3, 31, 23, 59, 59, 0, now.Location())
    } else {
        // 否则，结束月份是当前季度的最后一个月
        endTime = time.Date(now.Year(), endMonth, 31, 23, 59, 59, 59, now.Location())
        if endMonth == time.December {
            // 如果是第四季度，年份要加1
            endTime = time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location()).Add(-1 * time.Second)
        } else {
            // 否则，年份不变
            endTime = time.Date(now.Year(), endMonth+1, 1, 0, 0, 0, 0, now.Location()).Add(-1 * time.Second)
        }
    }
    startTime = time.Date(now.Year(), startMonth, 1, 0, 0, 0, 0, now.Location())

    return startTime, endTime
}

func getCurrentMonth() (time.Time, time.Time) {
    now := time.Now()
    startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
    endOfMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-1 * time.Second)
    return startOfMonth, endOfMonth
}

func GetSystemStatus() (map[string]interface{}) {
    var userslist []types.XqUser
    var users types.XqUser
    var adminCount int64
    var userCount int64
    var onlineUsers int64
    db.Model(&users).Where("role = 1").Count(&adminCount)
    db.Model(&users).Where("role = 0").Count(&userCount)
    db.Where("token <> '' AND role = 0").Find(&userslist)
    for _, user := range userslist {
        if res := GetUserByToken(user.Token); res != nil {
            onlineUsers++
        }
    }
    cpuUsage, memUsage, diskUsage := utils.GetSystemUsage()
    return map[string]interface{}{
        "cpu":    cpuUsage,
        "mem":    memUsage,
        "disk":   diskUsage,
        "admintotal": adminCount,
        "usertotal": userCount,
        "onlineuser": onlineUsers,
    }
}
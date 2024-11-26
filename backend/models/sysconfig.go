package models

import (
    "fmt"
    "mime"
    "time"

    "xuanqiong/backend/types"
    "xuanqiong/backend/utils"
    "github.com/go-gomail/gomail"
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
    if err != nil {
        return err
    }
    res := db.First(&emailConfig).RowsAffected
    if res != 0 {
        configData.EmailConfig.UpdateTime = time.Now()
        err = db.Model(&emailConfig).Updates(configData.EmailConfig).Error
        if err != nil {
            return err
        }
    } else {
        configData.EmailConfig.CreateTime = time.Now()
        err = db.Create(&configData.EmailConfig).Error
        if err != nil {
            return err
        }
    }
    
    db.First(&jwtConfig)
    err = db.Model(&jwtConfig).Updates(configData.JwtConfig).Error
    if err != nil {
        return err
    }
    res = db.First(&noticeConfig).RowsAffected
    if res != 0 {
        configData.NoticeConfig.UpdateTime = time.Now()
        err = db.Model(&noticeConfig).Updates(configData.NoticeConfig).Error
        if err != nil {
            return err
        }
    } else {
        configData.NoticeConfig.CreateTime = time.Now()
        err = db.Create(&configData.NoticeConfig).Error
        if err != nil {
            return err
        }
    }
    return nil
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
    lockip.UpdateTime = time.Now()
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

// 发送验证码
func SendCaptcha(email string) error {
    var verifycode types.XqVerifyCode
    
    captcha, err := utils.GenerateRandomChars(6, 1)
    if err != nil {
        return err
    }
    
    res := db.Where("email = ?", email).First(&verifycode).RowsAffected
    if res != 0 {
        timeDiff := time.Now().Sub(verifycode.UpdateTime)
        if timeDiff < 2*time.Minute {
            return nil
        }
        verifycode.Code = captcha
        verifycode.UpdateTime = time.Now()
        verifycode.ExpiredTime = time.Now().Add(time.Minute * 5)
        db.Save(&verifycode)
        return sendEmail(email, captcha)
    }
    verifycode.Email = email
    verifycode.Code = captcha
    verifycode.CreateTime = time.Now()
    verifycode.UpdateTime = time.Now()
    verifycode.ExpiredTime = time.Now().Add(time.Minute * 5)
    db.Create(&verifycode)
    return sendEmail(email, captcha)
}

func sendEmail(email, captcha string) error {
    var emailConfig types.XqEmailConfig
    db.First(&emailConfig)
    m := gomail.NewMessage()
    from := mime.QEncoding.Encode("UTF-8", emailConfig.EmailSender) + " <" + emailConfig.EmailUser + ">"
    m.SetHeader("From", from)
    m.SetHeader("To", email)
    m.SetHeader("Subject", "欢迎使用玄穹漏洞平台")
    m.SetBody("text/plain", "尊敬的用户，\r\n您好！欢迎使用玄穹漏洞平台，您的验证码是：" + captcha +
        "\r\n请注意为了您的账户安全，请勿将验证码透露给任何人，验证码有效期为5分钟。")
    d := gomail.NewDialer(emailConfig.EmailHost, int(emailConfig.EmailPort), emailConfig.EmailUser, emailConfig.EmailPassword)
    return d.DialAndSend(m)
}

// 批量删除
func MultiDelete(model string, ids []interface{}) error {
    switch model {
        case "user":
            return db.Delete(&types.XqUser{}, ids).Error
        case "vuln":
            return db.Delete(&types.XqVulnerability{}, ids).Error
        case "vulntype":
            return db.Delete(&types.XqVulnType{}, ids).Error
        case "scorerule":
            return db.Delete(&types.XqScoreRule{}, ids).Error
        default:
            return fmt.Errorf("不支持的模型")
    }
}
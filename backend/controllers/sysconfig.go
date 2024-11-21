package controllers

import (
    "xuanqiong/backend/types"
    "xuanqiong/backend/models"
    "github.com/gin-gonic/gin"
)

// 获取系统状态-管理员
func GetSystemStatus(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        data := models.GetSystemStatus()
        c.JSON(200, gin.H{"code": 1, "data": data})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取系统配置
func GetSystemConfig(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        sysconf, emailconf, jwtconf, noticeconf := models.GetSystemConfig()
        c.JSON(200, gin.H{"code": 1, "data": gin.H{"sysconf": sysconf, "emailconf": emailconf, "jwtconf": jwtconf, "noticeconf": noticeconf}})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 更新系统配置
func UpdateSystemConfig(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var configData types.SystemConfigData
        if err := c.ShouldBindJSON(&configData); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input", "err": err})
            return
        }
        err := models.UpdateSystemConfig(configData)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": "Update failed", "err": err})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update success"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取验证码
func GetCaptcha(c *gin.Context) {
    email := c.Query("email")
    err := models.SendCaptcha(email)
    if err != nil {
        c.JSON(200, gin.H{"code": 0, "msg": "Send email failed"})
        return
    }
    c.JSON(200, gin.H{"code": 1, "msg": "Send email success"})
}
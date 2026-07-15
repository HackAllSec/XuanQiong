package controllers

import (
	"github.com/gin-gonic/gin"
	"xuanqiong/backend/models"
	"xuanqiong/backend/types"
)

// 获取系统状态-管理员
func GetSystemStatus(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser != nil {
		data := models.GetSystemStatus()
		c.JSON(200, gin.H{"code": 1, "data": data})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取系统配置
func GetSystemConfig(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser != nil {
		sysconf, emailconf, jwtconf, noticeconf := models.GetSystemConfig()
		if emailconf.EmailPassword != "" {
			emailconf.EmailPassword = models.MaskedSecretValue
		}
		if jwtconf.JwtSecret != "" {
			jwtconf.JwtSecret = models.MaskedSecretValue
		}
		if noticeconf.Secret != "" {
			noticeconf.Secret = models.MaskedSecretValue
		}
		c.JSON(200, gin.H{"code": 1, "data": gin.H{"sysconf": sysconf, "emailconf": emailconf, "jwtconf": jwtconf, "noticeconf": noticeconf}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 更新系统配置
func UpdateSystemConfig(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser != nil {
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
	if !models.IsEmailValid(email) {
		c.JSON(200, gin.H{"code": 2, "msg": "Invalid email"})
		return
	}
	if err := models.AllowCaptchaRequest(c.ClientIP()); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": "Captcha request too frequent"})
		return
	}
	err := models.SendCaptcha(email)
	if err != nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Send email failed"})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Send email success"})
}

func GetBrandConfig(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "data": models.GetBrandPublicConfig()})
}

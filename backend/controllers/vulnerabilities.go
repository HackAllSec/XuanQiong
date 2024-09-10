package controllers

import (
    "xuanqiong/models"
    "github.com/gin-gonic/gin"
)

// 获取漏洞摘要，无需登录
func GetVulnAbstract(c *gin.Context) {
    res := models.GetVulnAbstract()
    c.JSON(200, gin.H{"message": res})
}

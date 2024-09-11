package controllers

import (
    "xuanqiong/types"
    "xuanqiong/models"
    "github.com/gin-gonic/gin"
)

// 获取漏洞摘要，无需登录
func GetVulnAbstract(c *gin.Context) {
    res, total, hasPoc, hasExp := models.GetVulnAbstract()
    c.JSON(200, gin.H{"total": total, "hasPoc": hasPoc, "hasExp": hasExp, "data": res})
}

// 获取漏洞详细信息，登录和未登录情况
func GetVulnDetail(c *gin.Context) {
    id := c.Query("id")
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil{
        res := models.GetVulnDetailAuthed(id)
        c.JSON(200, gin.H{"message": res})
        return
    }
    res := models.GetVulnDetail(id)
    c.JSON(200, gin.H{"message": res})
}

// 添加漏洞信息
func AddVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var vulnerabilities types.Vulnerability
        if err := c.ShouldBindJSON(&vulnerabilities); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"+err.Error()})
            return
        }
        vulnerabilities.Submit = currentUser.Username
        res, err := models.InsertVuln(vulnerabilities)
        if err != nil {
            c.JSON(200, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "提交成功","result": res})
        return
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 搜索漏洞
func SearchVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        keyword := c.Query("keyword")
        res := models.SearchVuln(keyword)
        c.JSON(200, gin.H{"message": res})
        return
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 高级搜索
func SearchVulnAdv(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        res := models.SearchVulnAdv(data)
        c.JSON(200, gin.H{"message": res})
        return
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

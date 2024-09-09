package controllers

import (
	"strconv"

	"xuanqiong/config"
	"xuanqiong/types"
	"xuanqiong/models"
	"github.com/gin-gonic/gin"
)

var (
	cfg = config.Config
	maxAttempts = cfg.Login.MaxAttempts
)

// 登录
func Login(c *gin.Context) {
	if models.IsLocked(c.ClientIP()) {
		c.JSON(429, gin.H{"error": "Too many login attempts. Please try again later."})
		return
	}
	var logindata types.LoginData
	if err := c.ShouldBindJSON(&logindata); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	loginUser := models.CheckLogin(logindata.Username, logindata.Password)
	if loginUser != nil {
		c.JSON(200, gin.H{"msg":"Login Successful", "Username":loginUser.Username, "token": loginUser.Token})
	} else {
		maxAttempts--
		if maxAttempts == 0 {
			models.LockIP(c.ClientIP(), cfg.Login.LockoutDuration)
			c.JSON(429, gin.H{"error": "Too many login attempts. Please try again later."})
			maxAttempts = cfg.Login.MaxAttempts
			return
		}
		c.JSON(401, gin.H{"error": "Invalid username or password, you can try " + strconv.Itoa(maxAttempts) + " times."})
		return
	}
}

// 退出
func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		err := models.CleanToken(token)
		if err == nil {
			c.JSON(200, gin.H{"已退出": "显示未登录页面"})
			return
		}
	}
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 创建用户
func CreateUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		if currentUser.Role != 1 {
			c.JSON(200, gin.H{"error": "Permission denied"})
			return
		}
		var logindata types.LoginData
		if err := c.ShouldBindJSON(&logindata); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
	}
	//c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

func Index(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		c.JSON(200, gin.H{"message": currentUser})
	}
	c.JSON(200, gin.H{"message": "未登录"})
}

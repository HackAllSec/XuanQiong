package controllers

import (
    "xuanqiong/models"
    "github.com/gin-gonic/gin"
)

var (
    maxAttempts int64
)

// 登录
func Login(c *gin.Context) {
    LoginPolicy := models.GetSystemConfig()
    if models.IsLocked(c.ClientIP()) {
        c.JSON(200, gin.H{"code": 0, "msg": "Too many login attempts. Please try again later."})
        return
    }
    var data map[string]interface{}
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
        return
    }
    username, _ := data["username"].(string)
    password, _ := data["password"].(string)
    loginUser := models.CheckLogin(username, password)
    if loginUser != nil {
        maxAttempts = 0
        c.JSON(200, gin.H{"code": 1, "msg":"Login Successful", "username":loginUser.Username, "avatar":loginUser.Avatar, "token": loginUser.Token})
    } else {
        maxAttempts++
        if maxAttempts >= LoginPolicy.MaxAttempts {
            models.LockIP(c.ClientIP(), LoginPolicy.LockoutDuration)
            c.JSON(200, gin.H{"code": 0, "msg": "Too many login attempts. Please try again later."})
            return
        }
        c.JSON(200, gin.H{"code": 3, "times": LoginPolicy.MaxAttempts - maxAttempts, "msg": "Invalid username or password"})
        return
    }
}

// 退出
func Logout(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        err := models.CleanToken(currentUser.Username)
        if err == nil {
            c.JSON(200, gin.H{"code": 1, "msg": "The user has logged out."})
            return
        } else {
            c.JSON(200, gin.H{"code": 0, "msg": err})
        }
    }
}

// 创建用户
func CreateUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        username, _ := data["username"].(string)
        password, _ := data["password"].(string)
        err := models.CreateUser(username, password, 0)
        if err == nil {
            c.JSON(200, gin.H{"code": 1, "msg": "User created successfully."})
            return
        } else {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
    }
    c.JSON(302, gin.H{"code": 0, "msg": "Permission denied"})
}

// 删除用户
func DeleteUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        err := models.DeleteUser(username)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete " + username + " successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 启用或禁用用户
func SetUserStatus(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        // 设置状态
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        err := models.SetUserStatus(username)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update " + username + " status successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取所有用户信息
func GetUsers(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        users := models.GetUsers()
        c.JSON(200, gin.H{"code": 1, "data": users})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        password, _ := data["password"].(string)
        role, ok := data["role"].(int64)
        if !ok {
            role = -1
        }
        status, ok := data["status"].(int64)
        if !ok {
            status = -1
        }
        err := models.UpdateUser(username, password, role, status)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update " + username + " successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取当前用户信息
func GetUserInfo(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        totalCount, authCount, _ := models.GetUservulns(currentUser.ID)
        c.JSON(200, gin.H{"code": 1, "data": gin.H{"username": currentUser.Username,
            "avatar": currentUser.Avatar, "email": currentUser.Email,
            "phone": currentUser.Phone, "ranking": currentUser.Ranking,
            "total": totalCount, "authed": authCount},
        })
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 修改头像
func UpdateAvatar(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        res := models.UpdateAvatar(file, currentUser.ID)
        c.JSON(200, gin.H{"code": 1, "msg": "Update avatar successfully","data": res})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 修改用户个人信息
func UpdateUserInfo(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        username, _ := data["username"].(string)
        email, _ := data["email"].(string)
        phone, _ := data["phone"].(string)
        err := models.UpdateUserInfo(currentUser.ID, username, email, phone)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update user info successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 修改用户密码
func UpdateUserPassword(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        oldpassword, _ := data["oldpassword"].(string)
        newpassword, _ := data["newpassword"].(string)
        err := models.UpdateUserPassword(currentUser.ID, oldpassword, newpassword)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update user password successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 注册用户
func Register(c *gin.Context) {
    LoginPolicy := models.GetSystemConfig()
    if LoginPolicy.UserRegister == false {
        c.JSON(200, gin.H{"code": 0, "msg": "Register function is disabled"})
        return
    }
    var data map[string]interface{}
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
        return
    }
    username, _ := data["username"].(string)
    password, _ := data["password"].(string)
    email, _ := data["email"].(string)
    phone, _ := data["phone"].(string)
    err := models.Register(username, password, email, phone)
    if err == 2 {
        c.JSON(200, gin.H{"code": 2, "msg": "User already exists"})
        return
    }
    if err == 3 {
        c.JSON(200, gin.H{"code": 3, "msg": "Username and password cannot be the empty"})
        return
    }
    if err == 4 {
        c.JSON(200, gin.H{"code": 4, "msg": "Email format error"})
        return
    }
    if err == 5 {
        c.JSON(200, gin.H{"code": 5, "msg": "Email already exists"})
    }
    // 注册功能禁用返回code 0
    c.JSON(200, gin.H{"code": 1, "msg": "Register successfully"})
}

// 获取用户提交的漏洞
func GetUserVulnList(c *gin.Context) {
    page := c.Query("page")
    limit := c.Query("limit")
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        total, data := models.GetUserVulnList(currentUser.ID, page, limit)
        c.JSON(200, gin.H{"code": 1, "total": total, "data": data})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取用户积分排名
func GetUserTop10(c *gin.Context) {
    annual, quarterly, monthly := models.GetUserTop10()
    c.JSON(200, gin.H{"annual": annual, "quarterly": quarterly, "monthly": monthly})
}

// 审核漏洞-管理员
func AuditVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        id, _ := data["id"].(string)
        audit, _ := data["audit"].(float64)
        review, _ := data["review"].(string)
        cvss, _ := data["cvss"].(float64)
        prid, _ := data["prid"].(float64)
        erid, _ := data["erid"].(float64)
        irid, _ := data["irid"].(float64)
        orid, _ := data["orid"].(float64)
        err := models.AuditVuln(id, int64(audit), review, cvss, uint64(prid), uint64(erid), uint64(irid), uint64(orid))
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Audit successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取系统状态-管理员
func GetSystemStatus(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
            return
        }
        data := models.GetSystemStatus()
        c.JSON(200, gin.H{"code": 1, "data": data})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}
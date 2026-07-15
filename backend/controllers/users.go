package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"xuanqiong/backend/models"
)

var (
	loginAttemptsMu   sync.Mutex
	loginAttemptsByIP = map[string]int64{}
)

// 登录
func Login(c *gin.Context) {
	LoginPolicy, _, _, _ := models.GetSystemConfig()
	clientIP := c.ClientIP()
	if models.IsLocked(c.ClientIP()) {
		c.JSON(200, gin.H{"code": 0, "msg": "Too many login attempts. Please try again later."})
		clearLoginAttempts(clientIP)
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
		clearLoginAttempts(clientIP)
		c.JSON(200, gin.H{
			"code":                  1,
			"msg":                   "Login Successful",
			"username":              loginUser.Username,
			"avatar":                loginUser.Avatar,
			"token":                 loginUser.Token,
			"force_password_change": loginUser.ForcePasswordChange,
			"permissions":           models.GetUserPermissionCodes(loginUser.ID),
			"roles":                 models.GetUserRoleNames(loginUser.ID),
			"role_ids":              models.GetUserRoleIDs(loginUser.ID),
		})
	} else {
		attemptsLeft := increaseLoginAttempts(clientIP, LoginPolicy.MaxAttempts)
		if attemptsLeft <= 0 {
			models.LockIP(clientIP, LoginPolicy.LockoutDuration)
			clearLoginAttempts(clientIP)
			c.JSON(200, gin.H{"code": 0, "msg": "Too many login attempts. Please try again later."})
			return
		}
		c.JSON(200, gin.H{"code": 3, "times": attemptsLeft, "msg": "Invalid username or password"})
		return
	}
}

func increaseLoginAttempts(ip string, maxAttempts int64) int64 {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()

	loginAttemptsByIP[ip]++
	attempts := loginAttemptsByIP[ip]
	return maxAttempts - attempts
}

func clearLoginAttempts(ip string) {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()

	delete(loginAttemptsByIP, ip)
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
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
			return
		}
		username, _ := data["username"].(string)
		password, _ := data["password"].(string)
		email, _ := data["email"].(string)
		phone, _ := data["phone"].(string)
		if username == "" || password == "" {
			c.JSON(200, gin.H{"code": 3, "msg": "Username and password cannot be empty"})
			return
		}
		roleIDs := parseRoleIDs(data["role_ids"])
		err := models.CreateUserWithRoles(username, password, email, phone, roleIDs)
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
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
			return
		}
		userid, ok := data["userid"].(float64)
		if !ok {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
			return
		}
		if uint64(userid) == currentUser.ID {
			c.JSON(200, gin.H{"code": 0, "msg": "You can't delete yourself"})
			return
		}
		err := models.DeleteUser(uint64(userid))
		if err != nil {
			c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": 1, "msg": "Delete successfully"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 分页获取所有用户信息
func GetUsers(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		page := c.Query("page")
		limit := c.Query("limit")
		total, users := models.GetUsers(page, limit)
		c.JSON(200, gin.H{"code": 1, "total": total, "data": users})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input1"})
			return
		}
		userid, _ := data["id"].(float64)
		username, ok := data["username"].(string)
		if !ok {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input3"})
			return
		}
		password, _ := data["password"].(string)
		email, _ := data["email"].(string)
		phone, _ := data["phone"].(string)
		status, _ := data["status"].(float64)
		roleIDs := parseRoleIDs(data["role_ids"])
		err := models.UpdateUserWithRoles(uint64(userid), username, password, email, phone, int64(status), roleIDs)
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
			"total": totalCount, "authed": authCount,
			"permissions": models.GetUserPermissionCodes(currentUser.ID),
			"roles":       models.GetUserRoleNames(currentUser.ID),
			"role_ids":    models.GetUserRoleIDs(currentUser.ID)},
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
		c.JSON(200, gin.H{"code": 1, "msg": "Update avatar successfully", "data": res})
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
	LoginPolicy, _, _, _ := models.GetSystemConfig()
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
	captcha, _ := data["captcha"].(string)
	err := models.Register(username, password, email, phone, captcha)
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
		return
	}
	if err == 6 {
		c.JSON(200, gin.H{"code": 6, "msg": "Captcha error"})
		return
	}
	if err != 1 {
		c.JSON(200, gin.H{"code": 0, "msg": "Register failed"})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Register successfully"})
}

// 忘记密码
func ForgetPassword(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	email, _ := data["email"].(string)
	captcha, _ := data["captcha"].(string)
	password, _ := data["password"].(string)
	err := models.ForgetPassword(email, captcha, password)
	if err != nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Reset failed"})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Reset success"})
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

// 批量删除用户
func MultiDeleteUsers(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	currentUser := models.GetUserByToken(token)
	if currentUser != nil {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
			return
		}
		ids, _ := data["ids"].([]interface{})
		for _, id := range ids {
			userID, ok := normalizeDeleteUserID(id)
			if ok && userID == currentUser.ID {
				c.JSON(200, gin.H{"code": 0, "msg": "You can't delete yourself"})
				return
			}
		}
		err := models.MultiDelete("user", ids)
		if err != nil {
			c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

func normalizeDeleteUserID(value interface{}) (uint64, bool) {
	switch v := value.(type) {
	case float64:
		if v < 0 || v != float64(uint64(v)) {
			return 0, false
		}
		return uint64(v), true
	case string:
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, false
		}
		return id, true
	default:
		return 0, false
	}
}

func parseRoleIDs(value interface{}) []uint64 {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}
	result := make([]uint64, 0, len(items))
	for _, item := range items {
		if roleID, ok := item.(float64); ok && roleID > 0 {
			result = append(result, uint64(roleID))
		}
	}
	return result
}

package models

import (
    "fmt"
    "strings"
    "time"
    "mime/multipart"

    "xuanqiong/types"
    "xuanqiong/utils"
)

// 检查IP是否被锁定
func IsLocked(ip string) bool {
    var lockip types.Lockip
    if res := db.Raw("SELECT * FROM lockips WHERE client_ip = ? AND status = 1", ip).Scan(&lockip).RowsAffected; res == 1 {
        if lockip.LockoutUntil != nil && time.Now().Before(*lockip.LockoutUntil) {
            return true
        }
    }
    return false
}

// 检查登录凭据
func CheckLogin(username string, password string) *types.User {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ? AND status = 1", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    if user.Status != 1 {
        return nil
    }
    if utils.IsHashEqual(user.Password, password) {
        token, _ := utils.GenJWTToken(user.Username)
        db.Model(&user).Update("token", token)
        return &user
    }
    return nil
}

// 根据用户名获取用户信息
func GetUserByUsername(username string) *types.User {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ? AND status = 1", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    return &user
}

// 根据token获取用户信息
func GetUserByToken(token string) *types.User {
    var user types.User
    if token != "" {
        token = strings.TrimPrefix(token, "Bearer ")
        res := db.Raw("SELECT * FROM users WHERE token = ? AND status = 1", token).Scan(&user).RowsAffected
        if res == 0 {
            return nil
        }
        claims, _ := utils.DecryptJWTToken(token)
        if claims != nil {
            return &user
        }
    }
    return nil
}

// 锁定IP地址
func LockIP(ip string, duration int64) {
    var lockip types.Lockip
    db.Raw("SELECT * FROM lockips WHERE client_ip = ?", ip).Scan(&lockip)
    lockoutUntil := time.Now().Add(time.Duration(duration) * time.Minute)
    lockip.ClientIP = ip
    lockip.Status = 1
    lockip.LockoutUntil = &lockoutUntil
    db.Save(&lockip)
}

// 清除token
func CleanToken(username string) error {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    err := db.Model(&user).Update("token", "").Error
    if err != nil {
        return err
    }
    return nil
}

// 创建用户
func CreateUser(username string, password string, role int64) error {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        passwdHash := utils.GenPasswordHash(password)
        userData := types.User{
            Username:   username,
            Password:   passwdHash,
            Role:       role,
            CreateTime: time.Now(),
            Status:     1,
        }
        db.Create(&userData)
        return nil
    } else {
        return fmt.Errorf("User %s already exists.", username)
    }
}

// 删除用户
func DeleteUser(username string) error {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        return fmt.Errorf("User %s not found.", username)
    }
    db.Delete(&user)
    return nil
}

// 修改用户信息
func UpdateUser(username string, password string, role int64, status int64) error {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        return fmt.Errorf("User %s not found.", username)
    }
    updates := make(map[string]interface{})
    if password != "" {
        password = utils.GenPasswordHash(password)
        updates["password"] = password
    }
    if role != -1 {
        updates["role"] = role
    }
    if status != -1 {
        updates["status"] = status
    }
    db.Model(&user).Where("username = ?", username).Updates(updates)
    return nil
}

// 获取所有用户
func GetUsers() ([]types.User) {
    var users []types.User
    db.Select("id, username, role, create_time, status").Find(&users)
    return users
}

// 启用或禁用用户
func SetUserStatus(username string) error {
    var user types.User
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        return fmt.Errorf("User %s not found.", username)
    }
    if user.Status == 0 {
        db.Model(&user).Update("status", 1)
    } else {
        db.Model(&user).Update("status", 0)
    }
    return nil
}

// 修改头像
func UpdateAvatar(file *multipart.FileHeader, userid uint64) string {
    var user types.User
    db.Where("id = ? AND status = 1", userid).First(&user)
    fileid, err := StoreFile(file, userid)
    if err != nil {
        return ""
    }
    if user.Avatar != "" {
        DeleteFile(user.Avatar, userid)
    }
    db.Model(&user).Update("avatar", fileid)
    return fileid
}

// 获取用户漏洞情况
func GetUservulns(userid uint64) (int64, int64, []types.Vulnerability) {
    var totalCount int64
    var authCount int64
    var vulnDatas []types.Vulnerability
    db.Model(&vulnDatas).Where("user_id = ?", userid).Count(&totalCount)
    db.Model(&vulnDatas).Where("user_id = ? AND status = 1", userid).Count(&authCount)
    db.Model(&vulnDatas).Where("user_id = ?", userid).Order("id desc").Find(&vulnDatas)
    return totalCount, authCount, vulnDatas
}

// 修改用户个人信息
func UpdateUserInfo(userid uint64, username string, email string, phone string) error {
    var user types.User
    updates := make(map[string]interface{})
    res := db.Where("id = ? AND status = 1", userid).First(&user)
    if res.RowsAffected != 1 {
        return fmt.Errorf("User is not exits.")
    }
    if username != "" {
        res = db.Where("username = ? AND id != ?", username, userid).First(&user)
        if res.RowsAffected != 0 {
            return fmt.Errorf("Username is already in use.")
        }
        updates["username"] = username
    }
    if email != "" {
        if IsEmailValid(email) == false {
            return fmt.Errorf("Invalid email address.")
        }
        res = db.Where("email = ? AND id != ?", email, userid).First(&user)
        if res.RowsAffected != 0 {
            return fmt.Errorf("Email is already in use.")
        }
        updates["email"] = email
    }
    if phone != "" {
        updates["phone"] = phone
    }
    updates["update_time"] = time.Now()
    db.Model(&user).Where("id = ?", userid).Updates(updates)
    return nil
}

// 修改用户密码
func UpdateUserPassword(userid uint64, oldpassword string, newpassword string) error {
    var user types.User
    updates := make(map[string]interface{})
    res := db.Where("id = ? AND status = 1", userid).First(&user)
    if res.RowsAffected != 1 {
        return fmt.Errorf("User is not exits.")
    }
    if oldpassword == "" || newpassword == "" {
        return fmt.Errorf("oldpassword and newpassword cannot be empty")
    }
    if !utils.IsHashEqual(user.Password, oldpassword) {
        return fmt.Errorf("Old password is incorrect.")
    }
    newpassword = utils.GenPasswordHash(newpassword)
    updates["password"] = newpassword
    updates["update_time"] = time.Now()
    db.Model(&user).Where("id = ?", userid).Updates(updates)
    return nil
}

// 用户注册
func Register(username string, password string, email string, phone string) int64 {
    var user types.User
    res := db.Where("username = ?", username).First(&user)
    if res.RowsAffected != 0 {
        return 2
    }
    if username == "" || password == "" {
        return 3
    }
    if email != "" {
        if IsEmailValid(email) == false {
            return 4
        }
        res = db.Where("email = ?", email).First(&user)
        if res.RowsAffected != 0 {
            return 5
        }
    }
    userData := types.User{
        Username: username,
        Password: utils.GenPasswordHash(password),
        Email:    email,
        Phone:    phone,
        Role:     0,
        Status:   1,
        CreateTime: time.Now(),
    }
    db.Create(&userData)
    return 1
}

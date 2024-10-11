package models

import (
    "fmt"
    "strings"
    "time"

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
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
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
    res := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).RowsAffected
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
        res := db.Raw("SELECT * FROM users WHERE token = ?", token).Scan(&user).RowsAffected
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

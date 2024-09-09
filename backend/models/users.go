package models

import (
	"fmt"
	"strings"
	"time"

	"xuanqiong/types"
	"xuanqiong/utils"
)

var (
	user types.User
)

// 检查IP是否被锁定
func IsLocked(ip string) bool {
	var lockip types.Lockip
	if err := db.Raw("SELECT * FROM lockip WHERE client_ip = ?", ip).Scan(&lockip).Error; err == nil {
		if lockip.LockoutUntil != nil && time.Now().Before(*lockip.LockoutUntil) {
			return true
		}
	}
	return false
}

// 检查登录凭据
func CheckLogin(username string, password string) *types.User {
    err := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
	if err != nil {
		return nil
	}
	if user.Status != 1 {
		return nil
	}
	if utils.IsHashEqual(user.Password, password) {
		token, _ := utils.GenJWTToken(user.Username, user.Role)
		db.Model(&user).Update("token", token)
		return &user
	}
	return nil
}

// 根据用户名获取用户信息
func GetUserByUsername(username string) *types.User {
    err := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

// 根据token获取用户信息
func GetUserByToken(token string) *types.User {
    if token != "" {
        token = strings.TrimPrefix(token, "Bearer ")
		err := db.Raw("SELECT * FROM users WHERE token = ?", token).Scan(&user).Error
		if err != nil {
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
func LockIP(ip string, duration int) {
	var lockip types.Lockip
	db.Raw("SELECT * FROM lockip WHERE client_ip = ?", ip).Scan(&lockip)
	lockoutUntil := time.Now().Add(time.Duration(duration) * time.Minute)
	lockip.ClientIP = ip
	lockip.LockoutUntil = &lockoutUntil
	db.Save(&lockip)
}

// 清除token
func CleanToken(username string) error {
	err := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
	if err != nil {
		return err
	}
	err = db.Model(&user).Update("token", "").Error
	if err != nil {
		return err
	}
	return nil
}

// 创建用户
func CreateUser(username string, password string, role int) error {
	err := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
	if err != nil {
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
    result := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
    if result.Error != nil {
        return fmt.Errorf("User %s not found.", username)
    }
    db.Delete(&user)
    return nil
}

// 启用或禁用用户
func SetUserStatus(username string, status int) error {
    result := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
    if result.Error != nil {
        return fmt.Errorf("User %s not found.", username)
    }
    db.Model(&user).Update("status", status)
    return nil
}

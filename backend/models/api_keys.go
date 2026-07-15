package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"xuanqiong/backend/types"
	"xuanqiong/backend/utils"
)

const apiKeyPrefix = "xqk_"

func hashAPIKey(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}

func GenerateAPIKey(userID uint64, name string, expiresAt string) (types.XqAPIKey, string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return types.XqAPIKey{}, "", fmt.Errorf("name is required")
	}
	secret, err := utils.GenerateRandomChars(40, 4)
	if err != nil {
		return types.XqAPIKey{}, "", err
	}
	key := apiKeyPrefix + secret
	prefix := key
	if len(prefix) > 12 {
		prefix = prefix[:12]
	}

	var expires *time.Time
	if strings.TrimSpace(expiresAt) != "" {
		parsed, err := time.Parse(time.RFC3339, expiresAt)
		if err != nil {
			return types.XqAPIKey{}, "", fmt.Errorf("invalid expires_at")
		}
		expires = &parsed
	}

	record := types.XqAPIKey{
		UserID:     userID,
		Name:       name,
		KeyPrefix:  prefix,
		KeyHash:    hashAPIKey(key),
		Scopes:     "",
		Status:     1,
		ExpiresAt:  expires,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := db.Create(&record).Error; err != nil {
		return types.XqAPIKey{}, "", err
	}
	return record, key, nil
}

func ListAPIKeys(userID uint64) []types.XqAPIKey {
	var keys []types.XqAPIKey
	db.Select("id, user_id, name, key_prefix, scopes, status, last_used_at, expires_at, create_time, update_time").
		Where("user_id = ?", userID).
		Order("create_time desc").
		Find(&keys)
	return keys
}

func DeleteAPIKey(userID uint64, keyID uint64) error {
	res := db.Where("id = ? AND user_id = ?", keyID, userID).Delete(&types.XqAPIKey{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("api key not found")
	}
	return nil
}

func GetUserByAPIKey(key string) *types.XqUser {
	key = strings.TrimSpace(key)
	if key == "" || !strings.HasPrefix(key, apiKeyPrefix) {
		return nil
	}

	var record types.XqAPIKey
	if db.Where("key_hash = ? AND status = ?", hashAPIKey(key), 1).First(&record).RowsAffected != 1 {
		return nil
	}
	if record.ExpiresAt != nil && time.Now().After(*record.ExpiresAt) {
		return nil
	}

	var user types.XqUser
	if db.Where("id = ? AND status = 1", record.UserID).First(&user).RowsAffected != 1 {
		return nil
	}
	now := time.Now()
	_ = db.Model(&record).Updates(map[string]interface{}{
		"last_used_at": &now,
		"update_time":  now,
	}).Error
	return &user
}

package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"xuanqiong/backend/types"
	"xuanqiong/backend/utils"
)

const apiKeyPrefix = "xqk_"
const defaultAPIKeyTTL = 90 * 24 * time.Hour
const maxAPIKeyTTL = 365 * 24 * time.Hour

var allowedAPIKeyScopes = map[string]bool{
	"profile.read":      true,
	"vuln.self.read":    true,
	"attachment.upload": true,
	"attachment.delete": true,
	"vuln.submit":       true,
	"vuln.edit":         true,
	"vuln.read":         true,
	"vuln.audit.read":   true,
	"vuln.audit.write":  true,
	"vuln.import":       true,
	"vuln.export":       true,
	"message.read":      true,
	"message.update":    true,
}

func hashAPIKey(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}

func AvailableAPIKeyScopes(userID uint64) []string {
	userPermissions := map[string]bool{}
	for _, code := range GetUserPermissionCodes(userID) {
		userPermissions[code] = true
	}
	candidates := []string{
		"profile.read",
		"vuln.self.read",
		"attachment.upload",
		"attachment.delete",
		"vuln.submit",
		"vuln.edit",
		"vuln.read",
		"vuln.audit.read",
		"vuln.audit.write",
		"vuln.import",
		"vuln.export",
		"message.read",
		"message.update",
	}
	scopes := make([]string, 0, len(candidates))
	for _, scope := range candidates {
		if userPermissions[scope] {
			scopes = append(scopes, scope)
		}
	}
	return scopes
}

func normalizeAPIKeyScopes(userID uint64, scopes []string) ([]string, error) {
	if len(scopes) == 0 {
		return nil, fmt.Errorf("at least one scope is required")
	}
	userScopes := map[string]bool{}
	for _, scope := range AvailableAPIKeyScopes(userID) {
		userScopes[scope] = true
	}
	seen := map[string]bool{}
	normalized := make([]string, 0, len(scopes))
	for _, scope := range scopes {
		scope = strings.TrimSpace(scope)
		if scope == "" || seen[scope] {
			continue
		}
		if !allowedAPIKeyScopes[scope] || !userScopes[scope] {
			return nil, fmt.Errorf("invalid api key scope: %s", scope)
		}
		seen[scope] = true
		normalized = append(normalized, scope)
	}
	if len(normalized) == 0 {
		return nil, fmt.Errorf("at least one scope is required")
	}
	return normalized, nil
}

func encodeAPIKeyScopes(scopes []string) (string, error) {
	data, err := json.Marshal(scopes)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func parseAPIKeyScopes(raw string) []string {
	var scopes []string
	if err := json.Unmarshal([]byte(raw), &scopes); err == nil {
		return scopes
	}
	return nil
}

func validStoredAPIKeyScopes(userID uint64, scopes []string) []string {
	userScopes := map[string]bool{}
	for _, scope := range AvailableAPIKeyScopes(userID) {
		userScopes[scope] = true
	}
	valid := make([]string, 0, len(scopes))
	for _, scope := range scopes {
		if allowedAPIKeyScopes[scope] && userScopes[scope] {
			valid = append(valid, scope)
		}
	}
	return valid
}

func GenerateAPIKey(userID uint64, name string, expiresAt string, scopes []string) (types.XqAPIKey, string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return types.XqAPIKey{}, "", fmt.Errorf("name is required")
	}
	normalizedScopes, err := normalizeAPIKeyScopes(userID, scopes)
	if err != nil {
		return types.XqAPIKey{}, "", err
	}
	encodedScopes, err := encodeAPIKeyScopes(normalizedScopes)
	if err != nil {
		return types.XqAPIKey{}, "", err
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
		if !parsed.After(time.Now()) {
			return types.XqAPIKey{}, "", fmt.Errorf("expires_at must be in the future")
		}
		if parsed.After(time.Now().Add(maxAPIKeyTTL)) {
			return types.XqAPIKey{}, "", fmt.Errorf("expires_at exceeds maximum allowed lifetime")
		}
		expires = &parsed
	} else {
		defaultExpires := time.Now().Add(defaultAPIKeyTTL)
		expires = &defaultExpires
	}

	record := types.XqAPIKey{
		UserID:     userID,
		Name:       name,
		KeyPrefix:  prefix,
		KeyHash:    hashAPIKey(key),
		Scopes:     encodedScopes,
		ScopeList:  normalizedScopes,
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
	for i := range keys {
		keys[i].ScopeList = parseAPIKeyScopes(keys[i].Scopes)
	}
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

func GetUserByAPIKeyWithScopes(key string) (*types.XqUser, []string) {
	key = strings.TrimSpace(key)
	if key == "" || !strings.HasPrefix(key, apiKeyPrefix) {
		return nil, nil
	}

	var record types.XqAPIKey
	if db.Where("key_hash = ? AND status = ?", hashAPIKey(key), 1).First(&record).RowsAffected != 1 {
		return nil, nil
	}
	if record.ExpiresAt != nil && time.Now().After(*record.ExpiresAt) {
		return nil, nil
	}

	var user types.XqUser
	if db.Where("id = ? AND status = 1", record.UserID).First(&user).RowsAffected != 1 {
		return nil, nil
	}
	scopes := validStoredAPIKeyScopes(user.ID, parseAPIKeyScopes(record.Scopes))
	if len(scopes) == 0 {
		return nil, nil
	}
	now := time.Now()
	_ = db.Model(&record).Updates(map[string]interface{}{
		"last_used_at": &now,
		"update_time":  now,
	}).Error
	return &user, scopes
}

func GetUserByAPIKey(key string) *types.XqUser {
	user, _ := GetUserByAPIKeyWithScopes(key)
	return user
}

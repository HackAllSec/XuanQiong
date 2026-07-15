package models

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"xuanqiong/backend/types"
)

func TestAPIKeyLifecycle_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	user := types.XqUser{
		Username:   "api-user",
		Password:   "hash",
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}

	record, plaintext, err := GenerateAPIKey(user.ID, "integration", "")
	if err != nil {
		t.Fatalf("generate api key: %v", err)
	}
	if plaintext == "" || record.KeyHash == "" || record.KeyHash == plaintext {
		t.Fatalf("api key should return plaintext once and store only hash")
	}
	if record.KeyPrefix == "" || len(record.KeyPrefix) > 12 {
		t.Fatalf("unexpected key prefix: %q", record.KeyPrefix)
	}

	keys := ListAPIKeys(user.ID)
	if len(keys) != 1 {
		t.Fatalf("list api keys length = %d, want 1", len(keys))
	}
	if keys[0].KeyHash != "" {
		t.Fatalf("list api keys should not expose hash")
	}
	var backupBuffer bytes.Buffer
	if err := CreateSystemBackup(&backupBuffer); err != nil {
		t.Fatalf("create system backup: %v", err)
	}
	var backup SystemBackup
	if err := json.Unmarshal(backupBuffer.Bytes(), &backup); err != nil {
		t.Fatalf("unmarshal system backup: %v", err)
	}
	if len(backup.APIKeys) != 1 || backup.APIKeys[0].KeyHash != record.KeyHash {
		t.Fatalf("backup should preserve api key hash, got %#v", backup.APIKeys)
	}

	currentUser := GetUserByAPIKey(plaintext)
	if currentUser == nil || currentUser.ID != user.ID {
		t.Fatalf("api key did not authenticate expected user")
	}
	var refreshed types.XqAPIKey
	if err := db.First(&refreshed, record.ID).Error; err != nil {
		t.Fatalf("load refreshed api key: %v", err)
	}
	if refreshed.LastUsedAt == nil {
		t.Fatalf("api key last_used_at should be updated after use")
	}
	if err := DeleteAPIKey(user.ID, record.ID); err != nil {
		t.Fatalf("delete api key: %v", err)
	}
	if GetUserByAPIKey(plaintext) != nil {
		t.Fatalf("deleted api key should not authenticate")
	}
}

func TestMessageReadLifecycle_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	if err := CreateMessage(7, "review", "done", "vuln_audit", "HVD-1"); err != nil {
		t.Fatalf("create message: %v", err)
	}
	if unread := GetUnreadMessageCount(7); unread != 1 {
		t.Fatalf("unread count = %d, want 1", unread)
	}
	total, messages := GetUserMessages(7, "1", "10")
	if total != 1 || len(messages) != 1 || messages[0].IsRead {
		t.Fatalf("unexpected messages: total=%d data=%#v", total, messages)
	}
	if err := MarkMessageRead(7, messages[0].ID); err != nil {
		t.Fatalf("mark message read: %v", err)
	}
	if unread := GetUnreadMessageCount(7); unread != 0 {
		t.Fatalf("unread count after read = %d, want 0", unread)
	}
	if err := CreateMessage(7, "review2", "done", "vuln_audit", "HVD-2"); err != nil {
		t.Fatalf("create second message: %v", err)
	}
	if err := MarkAllMessagesRead(7); err != nil {
		t.Fatalf("mark all messages read: %v", err)
	}
	if unread := GetUnreadMessageCount(7); unread != 0 {
		t.Fatalf("unread count after mark all = %d, want 0", unread)
	}
}

func TestRestoreSystemBackupSkipsEmptyCollections_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	if err := db.Create(&types.XqUser{
		Username:   "before-restore",
		Password:   "hash",
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}).Error; err != nil {
		t.Fatalf("create user before restore: %v", err)
	}
	payload, err := json.Marshal(SystemBackup{Version: 1, CreatedAt: time.Now()})
	if err != nil {
		t.Fatalf("marshal backup: %v", err)
	}
	if err := RestoreSystemBackup(bytes.NewReader(payload)); err != nil {
		t.Fatalf("restore empty backup: %v", err)
	}
	var count int64
	db.Model(&types.XqUser{}).Count(&count)
	if count != 0 {
		t.Fatalf("users after empty restore = %d, want 0", count)
	}
}

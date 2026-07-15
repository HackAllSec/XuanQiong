package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"xuanqiong/backend/types"
)

func CreateMessage(userID uint64, title string, content string, messageType string, relatedID string) error {
	if userID == 0 || title == "" {
		return nil
	}
	return db.Create(&types.XqMessage{
		UserID:     userID,
		Title:      title,
		Content:    content,
		Type:       messageType,
		RelatedID:  relatedID,
		IsRead:     false,
		CreateTime: time.Now(),
	}).Error
}

func GetUserMessages(userID uint64, page string, pageSize string) (int64, []types.XqMessage) {
	var messages []types.XqMessage
	var total int64
	pageNum, pageSizeNum := normalizePagination(page, pageSize)
	db.Model(&types.XqMessage{}).Where("user_id = ?", userID).Count(&total)
	db.Where("user_id = ?", userID).
		Order("create_time desc").
		Limit(pageSizeNum).
		Offset((pageNum - 1) * pageSizeNum).
		Find(&messages)
	return total, messages
}

func GetUnreadMessageCount(userID uint64) int64 {
	var total int64
	db.Model(&types.XqMessage{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&total)
	return total
}

func MarkMessageRead(userID uint64, messageID uint64) error {
	var message types.XqMessage
	if db.Where("id = ? AND user_id = ?", messageID, userID).First(&message).RowsAffected != 1 {
		return fmt.Errorf("message not found")
	}
	now := time.Now()
	return db.Model(&message).Updates(map[string]interface{}{
		"is_read":   true,
		"read_time": &now,
	}).Error
}

func MarkAllMessagesRead(userID uint64) error {
	now := time.Now()
	return db.Model(&types.XqMessage{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Updates(map[string]interface{}{
			"is_read":   true,
			"read_time": &now,
		}).Error
}

func createAuditResultMessage(vuln types.XqVulnerability) error {
	title := "漏洞审核结果"
	status := "未通过"
	if vuln.Status == 1 {
		status = "已通过"
	}
	content := fmt.Sprintf("你提交的漏洞 %s（%s）审核%s。", vuln.ID, vuln.VulnName, status)
	if vuln.ReviewComments != "" {
		content += " 审核意见：" + vuln.ReviewComments
	}
	return CreateMessage(vuln.UserID, title, content, "vuln_audit", vuln.ID)
}

func clearUserMessages(tx *gorm.DB, userID uint64) error {
	return tx.Where("user_id = ?", userID).Delete(&types.XqMessage{}).Error
}

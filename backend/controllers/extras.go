package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xuanqiong/backend/models"
	"xuanqiong/backend/types"
)

func currentUserFromRequest(c *gin.Context) *types.XqUser {
	if value, exists := c.Get("current_user"); exists {
		if currentUser, ok := value.(*types.XqUser); ok {
			return currentUser
		}
	}
	return models.GetUserByToken(c.Request.Header.Get("Authorization"))
}

func GetMessages(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	total, messages := models.GetUserMessages(currentUser.ID, c.Query("page"), c.Query("limit"))
	c.JSON(200, gin.H{"code": 1, "total": total, "unread": models.GetUnreadMessageCount(currentUser.ID), "data": messages})
}

func MarkMessageRead(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	var payload types.MessagePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if err := models.MarkMessageRead(currentUser.ID, payload.ID); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Message updated"})
}

func MarkAllMessagesRead(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	if err := models.MarkAllMessagesRead(currentUser.ID); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Messages updated"})
}

func GetAPIKeys(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	c.JSON(200, gin.H{"code": 1, "data": models.ListAPIKeys(currentUser.ID)})
}

func CreateAPIKey(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	var payload types.CreateAPIKeyPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	record, key, err := models.GenerateAPIKey(currentUser.ID, payload.Name, payload.ExpiresAt)
	if err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "data": record, "api_key": key})
}

func DeleteAPIKey(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	var payload types.MessagePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if err := models.DeleteAPIKey(currentUser.ID, payload.ID); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "API key deleted"})
}

func ExportVulns(c *gin.Context) {
	var buffer bytes.Buffer
	if err := models.ExportVulnsCSV(&buffer); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.Set("audit_skip_response_body", true)
	filename := fmt.Sprintf("xuanqiong_vulns_%s.csv", time.Now().Format("20060102150405"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "text/csv; charset=utf-8", buffer.Bytes())
}

func ImportVulns(c *gin.Context) {
	currentUser := currentUserFromRequest(c)
	if currentUser == nil {
		c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if file.Size <= 0 || file.Size > 5<<20 {
		c.JSON(400, gin.H{"code": 3, "msg": "File size exceeds limit"})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	defer src.Close()
	imported, errors := models.ImportVulnsCSV(src, currentUser.ID)
	if imported == 0 && len(errors) > 0 {
		c.JSON(200, gin.H{"code": 3, "imported": imported, "errors": errors, "msg": "No vulnerabilities imported"})
		return
	}
	c.JSON(200, gin.H{"code": 1, "imported": imported, "errors": errors})
}

func ExportBackup(c *gin.Context) {
	var buffer bytes.Buffer
	if err := models.CreateSystemBackup(&buffer); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.Set("audit_skip_response_body", true)
	filename := fmt.Sprintf("xuanqiong_backup_%s.json", time.Now().Format("20060102150405"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/json; charset=utf-8", buffer.Bytes())
}

func RestoreBackup(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if file.Size <= 0 || file.Size > 50<<20 {
		c.JSON(400, gin.H{"code": 3, "msg": "File size exceeds limit"})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	defer src.Close()
	limited := io.LimitReader(src, 50<<20)
	if err := models.RestoreSystemBackup(limited); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Backup restored"})
}

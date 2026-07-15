package controllers

import (
	"github.com/gin-gonic/gin"
	"xuanqiong/backend/models"
	"xuanqiong/backend/types"
)

func GetPermissions(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "data": models.GetPermissionList()})
}

func GetRoles(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "data": models.GetRoleList()})
}

func CreateRole(c *gin.Context) {
	var payload types.RolePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if err := models.CreateRole(payload); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Role created successfully"})
}

func UpdateRole(c *gin.Context) {
	var payload types.RolePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if err := models.UpdateRole(payload); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Role updated successfully"})
}

func DeleteRole(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	roleID, ok := data["id"].(float64)
	if !ok {
		c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
		return
	}
	if err := models.DeleteRole(uint64(roleID)); err != nil {
		c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "Role deleted successfully"})
}

func GetAuditLogs(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	keyword := c.Query("keyword")
	action := c.Query("action")
	total, logs := models.GetAuditLogs(page, limit, keyword, action)
	c.JSON(200, gin.H{"code": 1, "total": total, "data": logs})
}

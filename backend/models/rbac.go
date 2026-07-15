package models

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"xuanqiong/backend/types"
)

const (
	PermissionAdminPanelAccess = "admin.panel.access"
)

var defaultPermissions = []types.XqPermission{
	{Code: PermissionAdminPanelAccess, Name: "后台访问", Category: "platform", Description: "允许登录并访问后台管理界面"},
	{Code: "dashboard.read", Name: "查看仪表盘", Category: "dashboard", Description: "允许查看后台仪表盘"},
	{Code: "system.config.read", Name: "查看系统配置", Category: "system", Description: "允许查看系统配置"},
	{Code: "system.config.update", Name: "修改系统配置", Category: "system", Description: "允许修改系统配置和品牌配置"},
	{Code: "user.read", Name: "查看用户", Category: "user", Description: "允许查看用户列表"},
	{Code: "user.create", Name: "创建用户", Category: "user", Description: "允许创建用户"},
	{Code: "user.update", Name: "编辑用户", Category: "user", Description: "允许编辑用户"},
	{Code: "user.delete", Name: "删除用户", Category: "user", Description: "允许删除用户"},
	{Code: "role.read", Name: "查看角色", Category: "role", Description: "允许查看角色列表"},
	{Code: "role.create", Name: "创建角色", Category: "role", Description: "允许创建角色"},
	{Code: "role.update", Name: "编辑角色", Category: "role", Description: "允许编辑角色"},
	{Code: "role.delete", Name: "删除角色", Category: "role", Description: "允许删除角色"},
	{Code: "audit.log.read", Name: "查看操作日志", Category: "audit", Description: "允许查看操作日志"},
	{Code: "vuln.type.read", Name: "查看漏洞类型", Category: "vulnerability", Description: "允许查看漏洞类型"},
	{Code: "vuln.type.manage", Name: "管理漏洞类型", Category: "vulnerability", Description: "允许增删改漏洞类型"},
	{Code: "vuln.read", Name: "查看漏洞", Category: "vulnerability", Description: "允许查看后台漏洞列表"},
	{Code: "vuln.delete", Name: "删除漏洞", Category: "vulnerability", Description: "允许删除漏洞"},
	{Code: "vuln.audit.read", Name: "查看审核队列", Category: "vulnerability", Description: "允许查看待审核和已审核漏洞"},
	{Code: "vuln.audit.write", Name: "执行漏洞审核", Category: "vulnerability", Description: "允许审核漏洞"},
	{Code: "score.rule.read", Name: "查看评分规则", Category: "score", Description: "允许查看评分规则"},
	{Code: "score.rule.manage", Name: "管理评分规则", Category: "score", Description: "允许增删改评分规则"},
	{Code: "profile.read", Name: "查看个人资料", Category: "profile", Description: "允许查看个人资料"},
	{Code: "profile.update", Name: "编辑个人资料", Category: "profile", Description: "允许编辑个人资料"},
	{Code: "password.update", Name: "修改密码", Category: "profile", Description: "允许修改个人密码"},
}

var defaultRoles = []struct {
	Name            string
	Code            string
	Description     string
	IsSystem        bool
	PermissionCodes []string
}{
	{
		Name:        "超级管理员",
		Code:        "super_admin",
		Description: "拥有全部后台权限",
		IsSystem:    true,
		PermissionCodes: []string{
			PermissionAdminPanelAccess,
			"dashboard.read",
			"system.config.read",
			"system.config.update",
			"user.read",
			"user.create",
			"user.update",
			"user.delete",
			"role.read",
			"role.create",
			"role.update",
			"role.delete",
			"audit.log.read",
			"vuln.type.read",
			"vuln.type.manage",
			"vuln.read",
			"vuln.delete",
			"vuln.audit.read",
			"vuln.audit.write",
			"score.rule.read",
			"score.rule.manage",
			"profile.read",
			"profile.update",
			"password.update",
		},
	},
	{
		Name:        "漏洞审核员",
		Code:        "vuln_reviewer",
		Description: "负责漏洞审核和漏洞内容查看",
		IsSystem:    true,
		PermissionCodes: []string{
			PermissionAdminPanelAccess,
			"dashboard.read",
			"vuln.type.read",
			"vuln.read",
			"vuln.audit.read",
			"vuln.audit.write",
			"score.rule.read",
			"profile.read",
			"profile.update",
			"password.update",
		},
	},
	{
		Name:        "日志审计员",
		Code:        "audit_viewer",
		Description: "负责查看后台操作日志",
		IsSystem:    true,
		PermissionCodes: []string{
			PermissionAdminPanelAccess,
			"dashboard.read",
			"audit.log.read",
			"profile.read",
			"profile.update",
			"password.update",
		},
	},
}

func syncRBACDefaults() error {
	permissionByCode := map[string]types.XqPermission{}
	for _, permission := range defaultPermissions {
		permission.CreateTime = time.Now()
		permission.UpdateTime = time.Now()

		var existing types.XqPermission
		if err := db.Where("code = ?", permission.Code).First(&existing).Error; err == nil {
			updates := map[string]interface{}{
				"name":        permission.Name,
				"category":    permission.Category,
				"description": permission.Description,
				"update_time": time.Now(),
			}
			db.Model(&existing).Updates(updates)
			existing.Name = permission.Name
			existing.Category = permission.Category
			existing.Description = permission.Description
			permissionByCode[existing.Code] = existing
			continue
		}
		if err := db.Create(&permission).Error; err != nil {
			return err
		}
		permissionByCode[permission.Code] = permission
	}

	for _, roleDef := range defaultRoles {
		role := types.XqRole{
			Name:        roleDef.Name,
			Code:        roleDef.Code,
			Description: roleDef.Description,
			IsSystem:    roleDef.IsSystem,
			Status:      1,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}

		var existingRole types.XqRole
		if err := db.Where("code = ?", role.Code).First(&existingRole).Error; err == nil {
			db.Model(&existingRole).Updates(map[string]interface{}{
				"name":        role.Name,
				"description": role.Description,
				"is_system":   role.IsSystem,
				"status":      1,
				"update_time": time.Now(),
			})
			role = existingRole
		} else {
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		}

		if err := saveRolePermissions(role.ID, roleDef.PermissionCodes); err != nil {
			return err
		}
	}

	var adminUser types.XqUser
	if db.Where("username = ?", "admin").First(&adminUser).RowsAffected == 1 {
		if err := AssignRoleByCode(adminUser.ID, "super_admin"); err != nil {
			return err
		}
	}
	return nil
}

func saveRolePermissions(roleID uint64, permissionCodes []string) error {
	codes := normalizeStringList(permissionCodes)
	var permissions []types.XqPermission
	if len(codes) > 0 {
		if err := db.Where("code IN ?", codes).Find(&permissions).Error; err != nil {
			return err
		}
	}

	if err := db.Where("role_id = ?", roleID).Delete(&types.XqRolePermission{}).Error; err != nil {
		return err
	}
	for _, permission := range permissions {
		if err := db.Create(&types.XqRolePermission{
			RoleID:       roleID,
			PermissionID: permission.ID,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func normalizeStringList(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func AssignRoleByCode(userID uint64, roleCode string) error {
	var role types.XqRole
	if db.Where("code = ?", roleCode).First(&role).RowsAffected != 1 {
		return fmt.Errorf("role %s not found", roleCode)
	}
	return UpdateUserRoles(userID, []uint64{role.ID})
}

func UpdateUserRoles(userID uint64, roleIDs []uint64) error {
	roleIDs = normalizeUint64List(roleIDs)
	if err := db.Where("user_id = ?", userID).Delete(&types.XqUserRole{}).Error; err != nil {
		return err
	}

	for _, roleID := range roleIDs {
		if err := db.Create(&types.XqUserRole{
			UserID:     userID,
			RoleID:     roleID,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}).Error; err != nil {
			return err
		}
	}

	var user types.XqUser
	if db.Where("id = ?", userID).First(&user).RowsAffected != 1 {
		return fmt.Errorf("user not found")
	}

	legacyRole := int64(0)
	if UserHasAnyPermission(userID, PermissionAdminPanelAccess) {
		legacyRole = 1
	}
	return db.Model(&user).Updates(map[string]interface{}{
		"role":        legacyRole,
		"update_time": time.Now(),
	}).Error
}

func normalizeUint64List(values []uint64) []uint64 {
	seen := map[uint64]bool{}
	result := make([]uint64, 0, len(values))
	for _, value := range values {
		if value == 0 || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func GetUserRoles(userID uint64) []types.XqRole {
	var roles []types.XqRole
	db.Table("xq_roles").
		Select("xq_roles.*").
		Joins("join xq_user_roles on xq_user_roles.role_id = xq_roles.id").
		Where("xq_user_roles.user_id = ? AND xq_roles.status = 1", userID).
		Order("xq_roles.id asc").
		Scan(&roles)
	return roles
}

func GetUserRoleIDs(userID uint64) []uint64 {
	roles := GetUserRoles(userID)
	result := make([]uint64, 0, len(roles))
	for _, role := range roles {
		result = append(result, role.ID)
	}
	return result
}

func GetUserRoleNames(userID uint64) []string {
	roles := GetUserRoles(userID)
	result := make([]string, 0, len(roles))
	for _, role := range roles {
		result = append(result, role.Name)
	}
	return result
}

func GetUserPermissionCodes(userID uint64) []string {
	var permissions []types.XqPermission
	db.Table("xq_permissions").
		Select("distinct xq_permissions.code, xq_permissions.id, xq_permissions.name, xq_permissions.category, xq_permissions.description, xq_permissions.create_time, xq_permissions.update_time").
		Joins("join xq_role_permissions on xq_role_permissions.permission_id = xq_permissions.id").
		Joins("join xq_user_roles on xq_user_roles.role_id = xq_role_permissions.role_id").
		Where("xq_user_roles.user_id = ?", userID).
		Order("xq_permissions.code asc").
		Scan(&permissions)

	result := make([]string, 0, len(permissions))
	for _, permission := range permissions {
		result = append(result, permission.Code)
	}
	return result
}

func UserHasAnyPermission(userID uint64, codes ...string) bool {
	if len(codes) == 0 {
		return true
	}
	userPermissions := map[string]bool{}
	for _, code := range GetUserPermissionCodes(userID) {
		userPermissions[code] = true
	}
	for _, code := range codes {
		if userPermissions[code] {
			return true
		}
	}
	return false
}

func GetPermissionList() []types.XqPermission {
	var permissions []types.XqPermission
	db.Order("category asc, id asc").Find(&permissions)
	return permissions
}

func GetRoleList() []map[string]interface{} {
	var roles []types.XqRole
	db.Order("id asc").Find(&roles)
	result := make([]map[string]interface{}, 0, len(roles))
	for _, role := range roles {
		result = append(result, map[string]interface{}{
			"id":               role.ID,
			"name":             role.Name,
			"code":             role.Code,
			"description":      role.Description,
			"is_system":        role.IsSystem,
			"status":           role.Status,
			"permission_codes": GetRolePermissionCodes(role.ID),
		})
	}
	return result
}

func GetRolePermissionCodes(roleID uint64) []string {
	var permissions []types.XqPermission
	db.Table("xq_permissions").
		Select("xq_permissions.*").
		Joins("join xq_role_permissions on xq_role_permissions.permission_id = xq_permissions.id").
		Where("xq_role_permissions.role_id = ?", roleID).
		Order("xq_permissions.code asc").
		Scan(&permissions)

	result := make([]string, 0, len(permissions))
	for _, permission := range permissions {
		result = append(result, permission.Code)
	}
	return result
}

func CreateRole(payload types.RolePayload) error {
	payload.PermissionCodes = normalizeStringList(payload.PermissionCodes)
	if payload.Name == "" || payload.Code == "" {
		return fmt.Errorf("role name and code are required")
	}

	var existing types.XqRole
	if db.Where("code = ?", payload.Code).First(&existing).RowsAffected != 0 {
		return fmt.Errorf("role code already exists")
	}

	role := types.XqRole{
		Name:        payload.Name,
		Code:        payload.Code,
		Description: payload.Description,
		IsSystem:    false,
		Status:      1,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if payload.Status != 0 {
		role.Status = payload.Status
	}
	if err := db.Create(&role).Error; err != nil {
		return err
	}
	return saveRolePermissions(role.ID, payload.PermissionCodes)
}

func UpdateRole(payload types.RolePayload) error {
	var role types.XqRole
	if db.Where("id = ?", payload.ID).First(&role).RowsAffected != 1 {
		return fmt.Errorf("role not found")
	}
	if payload.Name == "" || payload.Code == "" {
		return fmt.Errorf("role name and code are required")
	}

	var duplicate types.XqRole
	if db.Where("code = ? AND id <> ?", payload.Code, payload.ID).First(&duplicate).RowsAffected != 0 {
		return fmt.Errorf("role code already exists")
	}

	updates := map[string]interface{}{
		"name":        payload.Name,
		"code":        payload.Code,
		"description": payload.Description,
		"status":      payload.Status,
		"update_time": time.Now(),
	}
	if err := db.Model(&role).Updates(updates).Error; err != nil {
		return err
	}
	if err := saveRolePermissions(role.ID, payload.PermissionCodes); err != nil {
		return err
	}

	var userRoles []types.XqUserRole
	db.Where("role_id = ?", role.ID).Find(&userRoles)
	for _, userRole := range userRoles {
		if err := UpdateUserRoles(userRole.UserID, GetUserRoleIDs(userRole.UserID)); err != nil {
			return err
		}
	}
	return nil
}

func DeleteRole(roleID uint64) error {
	var role types.XqRole
	if db.Where("id = ?", roleID).First(&role).RowsAffected != 1 {
		return fmt.Errorf("role not found")
	}
	if role.IsSystem {
		return fmt.Errorf("system role cannot be deleted")
	}
	if err := db.Where("role_id = ?", roleID).Delete(&types.XqRolePermission{}).Error; err != nil {
		return err
	}
	if err := db.Where("role_id = ?", roleID).Delete(&types.XqUserRole{}).Error; err != nil {
		return err
	}
	return db.Delete(&role).Error
}

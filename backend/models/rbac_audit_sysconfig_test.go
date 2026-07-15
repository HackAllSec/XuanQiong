package models

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"xuanqiong/backend/types"
)

func setupModelTestDB(t *testing.T) {
	t.Helper()

	dbName := strings.NewReplacer("/", "_", " ", "_").Replace(t.Name())
	testDB, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", dbName)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("open test database: %v", err)
	}
	if err := testDB.AutoMigrate(
		&types.XqSystemConfig{},
		&types.XqJwtConfig{},
		&types.XqEmailConfig{},
		&types.XqNoticeConfig{},
		&types.XqUser{},
		&types.XqRole{},
		&types.XqPermission{},
		&types.XqRolePermission{},
		&types.XqUserRole{},
		&types.XqVerifyCode{},
		&types.XqAuditLog{},
	); err != nil {
		t.Fatalf("migrate test database: %v", err)
	}
	db = testDB
}

func TestAuditSanitization_BitsUT(t *testing.T) {
	response := []byte(`{"code":1,"token":"raw-token","data":{"password":"raw-password","nested":[{"jwt":"raw-jwt"}]}}`)
	sanitized := SanitizeResponseBody(response)

	for _, leaked := range []string{"raw-token", "raw-password", "raw-jwt"} {
		if strings.Contains(sanitized, leaked) {
			t.Fatalf("sensitive value %q leaked in sanitized response: %s", leaked, sanitized)
		}
	}
	if strings.Count(sanitized, `"***"`) != 3 {
		t.Fatalf("expected all sensitive fields to be masked, got: %s", sanitized)
	}

	jsonReq, err := http.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(`{"username":"admin"}`))
	if err != nil {
		t.Fatalf("build json request: %v", err)
	}
	jsonReq.Header.Set("Content-Type", "application/json")
	if !ShouldCaptureRequestBody(jsonReq) {
		t.Fatalf("json request body should be captured for audit")
	}

	multipartReq, err := http.NewRequest(http.MethodPost, "/api/v1/upload", strings.NewReader("large-body"))
	if err != nil {
		t.Fatalf("build multipart request: %v", err)
	}
	multipartReq.Header.Set("Content-Type", "multipart/form-data; boundary=----test")
	if ShouldCaptureRequestBody(multipartReq) {
		t.Fatalf("multipart request body should not be captured by audit middleware")
	}
}

func TestRBACDefaultRoleAssignment_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	if err := syncRBACDefaults(); err != nil {
		t.Fatalf("sync default rbac data: %v", err)
	}

	verifyCode := types.XqVerifyCode{
		Email:       "new-user@example.com",
		Code:        "123456",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		ExpiredTime: time.Now().Add(time.Minute),
	}
	if err := db.Create(&verifyCode).Error; err != nil {
		t.Fatalf("create verify code: %v", err)
	}
	if status := Register("new-user", "password", "new-user@example.com", "18800000000", "123456"); status != 1 {
		t.Fatalf("register status = %d, want 1", status)
	}

	registered := GetUserByUsername("new-user")
	if registered == nil {
		t.Fatalf("registered user not found")
	}
	roleNames := GetUserRoleNames(registered.ID)
	if len(roleNames) != 1 || roleNames[0] != "普通用户" {
		t.Fatalf("registered user roles = %#v, want 普通用户", roleNames)
	}
	for _, permission := range []string{"profile.read", "vuln.submit", "attachment.upload"} {
		if !UserHasAnyPermission(registered.ID, permission) {
			t.Fatalf("registered user missing permission %s", permission)
		}
	}

	if err := CreateUserWithRoles("created-user", "password", "created@example.com", "", nil); err != nil {
		t.Fatalf("create user with default role: %v", err)
	}
	created := GetUserByUsername("created-user")
	if created == nil || !UserHasAnyPermission(created.ID, "vuln.submit") {
		t.Fatalf("created user should receive default user role permissions")
	}
}

func TestSyncRBACDefaultsMigratesLegacyAdmins_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	legacyAdmin := types.XqUser{
		Username:   "legacy-admin",
		Password:   "hash",
		Role:       1,
		Status:     1,
		Token:      "old-token",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := db.Create(&legacyAdmin).Error; err != nil {
		t.Fatalf("create legacy admin: %v", err)
	}

	if err := syncRBACDefaults(); err != nil {
		t.Fatalf("sync default rbac data: %v", err)
	}

	if !UserHasAnyPermission(legacyAdmin.ID, PermissionAdminPanelAccess) {
		t.Fatalf("legacy admin should receive super admin permission")
	}
	var refreshed types.XqUser
	if err := db.First(&refreshed, legacyAdmin.ID).Error; err != nil {
		t.Fatalf("load refreshed admin: %v", err)
	}
	if refreshed.Token != "" {
		t.Fatalf("legacy admin token should be cleared after role migration")
	}
}

func TestUpdateSystemConfigPersistsZeroValues_BitsUT(t *testing.T) {
	setupModelTestDB(t)
	if err := db.Create(&types.XqSystemConfig{
		UserRegister:        true,
		UserDisplay:         "visible",
		MaxAttempts:         9,
		LockoutDuration:     99,
		SiteName:            "site",
		FrontendTitle:       "front",
		AdminTitle:          "admin",
		LogoAttachmentID:    "logo",
		FaviconAttachmentID: "favicon",
		FooterText:          "footer",
		HelpURL:             "help",
		SuggestURL:          "suggest",
		CreateTime:          time.Now(),
		UpdateTime:          time.Now(),
	}).Error; err != nil {
		t.Fatalf("create system config: %v", err)
	}
	if err := db.Create(&types.XqJwtConfig{JwtSecret: "secret", JwtExpires: 3600}).Error; err != nil {
		t.Fatalf("create jwt config: %v", err)
	}
	if err := db.Create(&types.XqEmailConfig{EmailHost: "smtp", EmailPort: 25, EmailUser: "user", EmailPassword: "pass", EmailSender: "sender"}).Error; err != nil {
		t.Fatalf("create email config: %v", err)
	}
	if err := db.Create(&types.XqNoticeConfig{Type: 1, Secret: "secret", Webhook: "hook"}).Error; err != nil {
		t.Fatalf("create notice config: %v", err)
	}

	err := UpdateSystemConfig(types.SystemConfigData{
		SysConfig: types.XqSystemConfig{
			UserRegister:        false,
			UserDisplay:         "",
			MaxAttempts:         0,
			LockoutDuration:     0,
			SiteName:            "",
			FrontendTitle:       "",
			AdminTitle:          "",
			LogoAttachmentID:    "",
			FaviconAttachmentID: "",
			FooterText:          "",
			HelpURL:             "",
			SuggestURL:          "",
		},
		JwtConfig:    types.XqJwtConfig{JwtSecret: "", JwtExpires: 0},
		EmailConfig:  types.XqEmailConfig{EmailHost: "", EmailPort: 0, EmailUser: "", EmailPassword: "", EmailSender: ""},
		NoticeConfig: types.XqNoticeConfig{Type: 0, Secret: "", Webhook: ""},
	})
	if err != nil {
		t.Fatalf("update system config: %v", err)
	}

	sysConf, emailConf, jwtConf, noticeConf := GetSystemConfig()
	if sysConf.UserRegister || sysConf.UserDisplay != "" || sysConf.MaxAttempts != 0 || sysConf.LogoAttachmentID != "" || sysConf.FooterText != "" {
		t.Fatalf("system zero values were not persisted: %#v", sysConf)
	}
	if jwtConf.JwtSecret != "" || jwtConf.JwtExpires != 0 {
		t.Fatalf("jwt zero values were not persisted: %#v", jwtConf)
	}
	if emailConf.EmailHost != "" || emailConf.EmailPort != 0 || emailConf.EmailPassword != "" {
		t.Fatalf("email zero values were not persisted: %#v", emailConf)
	}
	if noticeConf.Type != 0 || noticeConf.Secret != "" || noticeConf.Webhook != "" {
		t.Fatalf("notice zero values were not persisted: %#v", noticeConf)
	}
}

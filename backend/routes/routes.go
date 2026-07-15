package routes

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"xuanqiong/backend/config"
	"xuanqiong/backend/controllers"
	"xuanqiong/backend/models"
)

var (
	route  *gin.Engine
	apiuri string = "/api/v1"
)

func resolveStaticFile(root string, requestPath string) (string, bool) {
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return "", false
	}
	cleanedPath := path.Clean("/" + requestPath)
	relativePath := strings.TrimPrefix(cleanedPath, "/")
	if relativePath == "" || relativePath == "." {
		return "", false
	}
	filePath, err := filepath.Abs(filepath.Join(rootAbs, relativePath))
	if err != nil {
		return "", false
	}
	relPath, err := filepath.Rel(rootAbs, filePath)
	if err != nil {
		return "", false
	}
	if relPath == ".." || strings.HasPrefix(relPath, ".."+string(os.PathSeparator)) {
		return "", false
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil || fileInfo.IsDir() {
		return "", false
	}
	return filePath, true
}

func init() {
	switch config.Config.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	}
	route = gin.Default()
	route.Use(normalizeAccessTokenHeaderMiddleware(), currentUserMiddleware(), auditLogMiddleware(), passwordChangeRequiredMiddleware())
	fmt.Println("Welcome to XuanQiong", config.Version)
	fmt.Println("Server running on " + config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

func passwordChangeRequiredMiddleware() gin.HandlerFunc {
	allowedPaths := map[string]bool{
		apiuri + "/login":          true,
		apiuri + "/logout":         true,
		apiuri + "/updatepassword": true,
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Next()
			return
		}

		currentUser := models.GetUserByToken(token)
		if currentUser == nil || !currentUser.ForcePasswordChange || allowedPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(200, gin.H{"code": 9, "msg": "Password change required"})
	}
}

func defineRouter() {
	// 无需登录
	route.POST(apiuri+"/register", auditedPublicRoute("auth.register", controllers.Register)...)
	route.POST(apiuri+"/login", auditedPublicRoute("auth.login", controllers.Login)...)
	route.GET(apiuri+"/getvulnabs", controllers.GetVulnAbstract)
	route.GET(apiuri+"/getvulntypes", controllers.GetVulnTypeList)
	route.GET(apiuri+"/getvulntype", controllers.GetVulnType)
	route.GET(apiuri+"/getvulnlist", controllers.GetVulnList)
	route.GET(apiuri+"/getvulndtl", controllers.GetVulnDetail)
	route.GET(apiuri+"/search", controllers.SearchVuln)
	route.POST(apiuri+"/advsearch", controllers.SearchVulnAdv)
	route.GET("/download/file", controllers.DownloadFile)
	route.GET(apiuri+"/usertop", controllers.GetUserTop10)
	route.GET(apiuri+"/getcaptcha", auditedPublicRoute("auth.captcha", controllers.GetCaptcha)...)
	route.POST(apiuri+"/forgetpassword", auditedPublicRoute("auth.forget_password", controllers.ForgetPassword)...)
	route.GET(apiuri+"/getbrandconfig", controllers.GetBrandConfig)
	// 需要登录
	route.GET(apiuri+"/logout", protectedRoute(nil, "auth.logout", controllers.Logout)...)
	route.GET(apiuri+"/userinfo", protectedRoute(nil, "profile.read", controllers.GetUserInfo)...)
	route.GET(apiuri+"/uservulnlist", protectedRoute(nil, "vuln.self.read", controllers.GetUserVulnList)...)
	route.POST(apiuri+"/upload", protectedRoute(nil, "attachment.upload", controllers.UploadFile)...)
	route.GET("/delete/file", protectedRoute(nil, "attachment.delete", controllers.DeleteFile)...)
	route.POST(apiuri+"/addvuln", protectedRoute(nil, "vuln.submit", controllers.AddVuln)...)
	route.POST(apiuri+"/editvuln", protectedRoute(nil, "vuln.edit", controllers.EditVuln)...)
	route.POST(apiuri+"/updateavatar", protectedRoute(nil, "profile.avatar.update", controllers.UpdateAvatar)...)
	route.POST(apiuri+"/updateuserinfo", protectedRoute(nil, "profile.update", controllers.UpdateUserInfo)...)
	route.POST(apiuri+"/updatepassword", protectedRoute(nil, "password.update", controllers.UpdateUserPassword)...)

	// 管理员权限
	route.POST(apiuri+"/adduser", protectedRoute([]string{"user.create"}, "user.create", controllers.CreateUser)...)
	route.POST(apiuri+"/deluser", protectedRoute([]string{"user.delete"}, "user.delete", controllers.DeleteUser)...)
	route.POST(apiuri+"/multidelusers", protectedRoute([]string{"user.delete"}, "user.multi_delete", controllers.MultiDeleteUsers)...)
	route.GET(apiuri+"/getusers", protectedRoute([]string{"user.read"}, "user.read", controllers.GetUsers)...)
	route.POST(apiuri+"/updateuser", protectedRoute([]string{"user.update"}, "user.update", controllers.UpdateUser)...)
	route.POST(apiuri+"/auditvuln", protectedRoute([]string{"vuln.audit.write"}, "vuln.audit", controllers.AuditVuln)...)
	route.GET(apiuri+"/getsystemstatus", protectedRoute([]string{"dashboard.read"}, "dashboard.read", controllers.GetSystemStatus)...)
	route.POST(apiuri+"/addvulntype", protectedRoute([]string{"vuln.type.manage"}, "vuln_type.create", controllers.AddVulnType)...)
	route.POST(apiuri+"/delvulntype", protectedRoute([]string{"vuln.type.manage"}, "vuln_type.delete", controllers.DeleteVulnType)...)
	route.POST(apiuri+"/multidelvulntypes", protectedRoute([]string{"vuln.type.manage"}, "vuln_type.multi_delete", controllers.MultiDeleteVulnTypes)...)
	route.POST(apiuri+"/updatevulntype", protectedRoute([]string{"vuln.type.manage"}, "vuln_type.update", controllers.UpdateVulnType)...)
	route.GET(apiuri+"/getunauditlist", protectedRoute([]string{"vuln.audit.read"}, "vuln.audit.read", controllers.GetUnauditList)...)
	route.GET(apiuri+"/getauditedlist", protectedRoute([]string{"vuln.audit.read"}, "vuln.audited.read", controllers.GetAuditedList)...)
	route.GET(apiuri+"/getsysconfig", protectedRoute([]string{"system.config.read"}, "system.config.read", controllers.GetSystemConfig)...)
	route.POST(apiuri+"/updatesysconfig", protectedRoute([]string{"system.config.update"}, "system.config.update", controllers.UpdateSystemConfig)...)
	route.POST(apiuri+"/addscorerule", protectedRoute([]string{"score.rule.manage"}, "score_rule.create", controllers.AddScoreRule)...)
	route.GET(apiuri+"/getallscorerules", protectedRoute([]string{"score.rule.read"}, "score_rule.read_all", controllers.GetAllScoreRules)...)
	route.GET(apiuri+"/getscorerules", protectedRoute([]string{"score.rule.read"}, "score_rule.read", controllers.GetScoreRules)...)
	route.POST(apiuri+"/editscorerule", protectedRoute([]string{"score.rule.manage"}, "score_rule.update", controllers.EditScoreRule)...)
	route.POST(apiuri+"/delscorerule", protectedRoute([]string{"score.rule.manage"}, "score_rule.delete", controllers.DeleteScoreRule)...)
	route.POST(apiuri+"/multidelscorerules", protectedRoute([]string{"score.rule.manage"}, "score_rule.multi_delete", controllers.MultiDeleteScoreRules)...)
	route.POST(apiuri+"/delvuln", protectedRoute([]string{"vuln.delete"}, "vuln.delete", controllers.DeleteVuln)...)
	route.POST(apiuri+"/multidelvulns", protectedRoute([]string{"vuln.delete"}, "vuln.multi_delete", controllers.MultiDeleteVulns)...)
	route.GET(apiuri+"/getroles", protectedRoute([]string{"role.read"}, "role.read", controllers.GetRoles)...)
	route.POST(apiuri+"/addrole", protectedRoute([]string{"role.create"}, "role.create", controllers.CreateRole)...)
	route.POST(apiuri+"/updaterole", protectedRoute([]string{"role.update"}, "role.update", controllers.UpdateRole)...)
	route.POST(apiuri+"/delrole", protectedRoute([]string{"role.delete"}, "role.delete", controllers.DeleteRole)...)
	route.GET(apiuri+"/getpermissions", protectedRoute([]string{"role.read", "role.create", "role.update"}, "permission.read", controllers.GetPermissions)...)
	route.GET(apiuri+"/getauditlogs", protectedRoute([]string{"audit.log.read"}, "audit.read", controllers.GetAuditLogs)...)
}

// 前后端分离的路由
func api() {
	// 配置CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = strings.Split(config.Config.Server.AllowOrigins, ",")
	corsConfig.AllowMethods = strings.Split(config.Config.Server.AllowMethods, ",")
	corsConfig.AllowHeaders = strings.Split(config.Config.Server.AllowHeaders, ",")
	if !containsHeader(corsConfig.AllowHeaders, "X-Auth-Token") {
		corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Auth-Token")
	}

	route.Use(cors.New(corsConfig))
	defineRouter()

	// 添加调试输出
	fmt.Println("CORS Configuration:")
	fmt.Printf("AllowOrigins: %v\n", corsConfig.AllowOrigins)
	fmt.Printf("AllowMethods: %v\n", corsConfig.AllowMethods)
	fmt.Printf("AllowHeaders: %v\n", corsConfig.AllowHeaders)

	route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

func containsHeader(headers []string, header string) bool {
	for _, item := range headers {
		if strings.EqualFold(strings.TrimSpace(item), header) {
			return true
		}
	}
	return false
}

// 前后端不分离的路由
func all() {
	frontendPath := config.Config.Server.FrontendPath
	frontendStaticUrl := config.Config.Server.StaticUrl
	adminPath := config.Config.Server.AdminPath
	AdminStaticUrl := config.Config.Server.AdminStaticUrl
	defineRouter()
	// 前端静态文件
	route.Static(frontendStaticUrl, frontendPath+frontendStaticUrl)
	route.Static(AdminStaticUrl, adminPath+AdminStaticUrl)

	// 设置前端模板文件的路由
	route.LoadHTMLGlob(frontendPath + "/*.html")
	route.LoadHTMLGlob(adminPath + "/*.html")

	// 定义路由
	route.GET("/", func(c *gin.Context) {
		c.File(frontendPath + "/index.html")
	})
	route.GET("/admin", func(c *gin.Context) {
		c.File(adminPath + "/index.html")
	})

	// 通配符路由
	route.NoRoute(func(c *gin.Context) {
		if filePath, ok := resolveStaticFile(frontendPath, c.Request.URL.Path); ok {
			c.File(filePath)
			return
		}
		c.HTML(404, "", nil)
	})
	route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

func StartServer() {
	switch config.Config.Server.StartMode {
	case "api":
		api()
	case "all":
		all()
	}
}

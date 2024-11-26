package routes

import (
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "xuanqiong/backend/controllers"
    "xuanqiong/backend/config"
)

var (
    route *gin.Engine
    apiuri string = "/api/v1"
)

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
    fmt.Println("Welcome to XuanQiong",config.Version)
    fmt.Println("Server running on " + config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

func defineRouter() {
    // 无需登录
    route.POST(apiuri + "/register", controllers.Register)
    route.POST(apiuri + "/login", controllers.Login)
    route.GET(apiuri + "/getvulnabs", controllers.GetVulnAbstract)
    route.GET(apiuri + "/getvulntypes", controllers.GetVulnTypeList)
    route.GET(apiuri + "/getvulntype", controllers.GetVulnType)
    route.GET(apiuri + "/getvulnlist", controllers.GetVulnList)
    route.GET(apiuri + "/getvulndtl", controllers.GetVulnDetail)
    route.GET(apiuri + "/search", controllers.SearchVuln)
    route.POST(apiuri + "/advsearch", controllers.SearchVulnAdv)
    route.GET("/download/file", controllers.DownloadFile)
    route.GET(apiuri + "/usertop", controllers.GetUserTop10)
    route.GET(apiuri + "/getcaptcha", controllers.GetCaptcha)
    route.POST(apiuri + "/forgetpassword", controllers.ForgetPassword)
    // 需要登录
    route.GET(apiuri + "/logout", controllers.Logout)
    route.GET(apiuri + "/userinfo", controllers.GetUserInfo)
    route.GET(apiuri + "/uservulnlist", controllers.GetUserVulnList)
    route.POST(apiuri + "/upload", controllers.UploadFile)
    route.GET("/delete/file", controllers.DeleteFile)
    route.POST(apiuri + "/addvuln", controllers.AddVuln)
    route.POST(apiuri + "/editvuln", controllers.EditVuln)
    route.POST(apiuri + "/updateavatar", controllers.UpdateAvatar)
    route.POST(apiuri + "/updateuserinfo", controllers.UpdateUserInfo)
    route.POST(apiuri + "/updatepassword", controllers.UpdateUserPassword)

    // 管理员权限
    route.POST(apiuri + "/adduser", controllers.CreateUser)
    route.POST(apiuri + "/deluser", controllers.DeleteUser)
    route.POST(apiuri + "/multidelusers", controllers.MultiDeleteUsers)
    route.GET(apiuri + "/getusers", controllers.GetUsers)
    route.POST(apiuri + "/updateuser", controllers.UpdateUser)
    route.POST(apiuri + "/auditvuln", controllers.AuditVuln)
    route.GET(apiuri + "/getsystemstatus", controllers.GetSystemStatus)
    route.POST(apiuri + "/addvulntype", controllers.AddVulnType)
    route.POST(apiuri + "/delvulntype", controllers.DeleteVulnType)
    route.POST(apiuri + "/multidelvulntypes", controllers.MultiDeleteVulnTypes)
    route.POST(apiuri + "/updatevulntype", controllers.UpdateVulnType)
    route.GET(apiuri + "/getunauditlist", controllers.GetUnauditList)
    route.GET(apiuri + "/getauditedlist", controllers.GetAuditedList)
    route.GET(apiuri + "/getsysconfig", controllers.GetSystemConfig)
    route.POST(apiuri + "/updatesysconfig", controllers.UpdateSystemConfig)
    route.POST(apiuri + "/addscorerule", controllers.AddScoreRule)
    route.GET(apiuri + "/getallscorerules", controllers.GetAllScoreRules)
    route.GET(apiuri + "/getscorerules", controllers.GetScoreRules)
    route.POST(apiuri + "/editscorerule", controllers.EditScoreRule)
    route.POST(apiuri + "/delscorerule", controllers.DeleteScoreRule)
    route.POST(apiuri + "/multidelscorerules", controllers.MultiDeleteScoreRules)
    route.POST(apiuri + "/delvuln", controllers.DeleteVuln)
    route.POST(apiuri + "/multidelvulns", controllers.MultiDeleteVulns)
}

// 前后端分离的路由
func api() {
    // 配置CORS
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = strings.Split(config.Config.Server.AllowOrigins, ",")
    corsConfig.AllowMethods = strings.Split(config.Config.Server.AllowMethods, ",")
    corsConfig.AllowHeaders = strings.Split(config.Config.Server.AllowHeaders, ",")
    
    route.Use(cors.New(corsConfig))
    defineRouter()

    // 添加调试输出
    fmt.Println("CORS Configuration:")
    fmt.Printf("AllowOrigins: %v\n", corsConfig.AllowOrigins)
    fmt.Printf("AllowMethods: %v\n", corsConfig.AllowMethods)
    fmt.Printf("AllowHeaders: %v\n", corsConfig.AllowHeaders)

    route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

// 前后端不分离的路由
func all() {
    frontendPath := config.Config.Server.FrontendPath
    frontendStaticUrl := config.Config.Server.StaticUrl
    adminPath := config.Config.Server.AdminPath
    AdminStaticUrl := config.Config.Server.AdminStaticUrl
    defineRouter()
    // 前端静态文件
    route.Static(frontendStaticUrl, frontendPath + frontendStaticUrl)
    route.Static(AdminStaticUrl, adminPath + AdminStaticUrl)

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
        path := c.Request.URL.Path
        if fileInfo, err := os.Stat(frontendPath + path); err == nil {
            if !fileInfo.IsDir() {
                c.File(frontendPath + path)
            }
        }
        c.HTML(404, "", nil)
    })
    route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

func StartServer(){
    switch config.Config.Server.StartMode {
        case "api":
            api()
        case "all":
            all()
    }
}
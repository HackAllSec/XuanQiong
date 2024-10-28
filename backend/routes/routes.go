package routes

import (
    "fmt"
    "os"
    "strconv"

    "github.com/gin-gonic/gin"
    "xuanqiong/controllers"
    "xuanqiong/config"
)

var (
    route *gin.Engine
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

// 前后端分离的路由
func InitRoutes() {
    route.POST("/api/v1/login", controllers.Login)
    route.GET("/api/v1/logout", controllers.Logout)
    route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}

// 前后端不分离的路由
func StartServer() {
    frontendPath := config.Config.Server.FrontendPath
    frontendStaticUrl := config.Config.Server.StaticUrl
    adminPath := config.Config.Server.AdminPath
    AdminStaticUrl := config.Config.Server.AdminStaticUrl
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
    /*
    route.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
        */
    // 无需登录
    route.POST("/api/v1/register", controllers.Register)
    route.POST("/api/v1/login", controllers.Login)
    route.GET("/api/v1/getvulnabs", controllers.GetVulnAbstract)
    route.GET("/api/v1/getvulntypes", controllers.GetVulnTypeList)
    route.GET("/api/v1/getvulnlist", controllers.GetVulnList)
    route.GET("/api/v1/getvulndtl", controllers.GetVulnDetail)
    route.GET("/api/v1/search", controllers.SearchVuln)
    route.POST("/api/v1/advsearch", controllers.SearchVulnAdv)
    route.GET("/download/file", controllers.DownloadFile)
    route.GET("/api/v1/usertop", controllers.GetUserTop10)
    // 需要登录
    route.GET("/api/v1/logout", controllers.Logout)
    route.GET("/api/v1/userinfo", controllers.GetUserInfo)
    route.GET("/api/v1/uservulnlist", controllers.GetUserVulnList)
    route.POST("/api/v1/upload", controllers.UploadFile)
    route.GET("/delete/file", controllers.DeleteFile)
    route.POST("/api/v1/addvuln", controllers.AddVuln)
    route.POST("/api/v1/editvuln", controllers.EditVuln)
    route.POST("/api/v1/updateavatar", controllers.UpdateAvatar)
    route.POST("/api/v1/updateuserinfo", controllers.UpdateUserInfo)
    route.POST("/api/v1/updatepassword", controllers.UpdateUserPassword)
    
    // 管理员权限
    route.POST("/api/v1/adduser", controllers.CreateUser)
    route.POST("/api/v1/deluser", controllers.DeleteUser)
    route.POST("/api/v1/userstatus", controllers.SetUserStatus)
    route.GET("/api/v1/getusers", controllers.GetUsers)
    route.POST("/api/v1/updateuser", controllers.UpdateUser)
    route.POST("/api/v1/auditvuln", controllers.AuditVuln)
    route.GET("/api/v1/getsystemstatus", controllers.GetSystemStatus)
    
    // 通配符路由
    route.NoRoute(func(c *gin.Context) {
        path := c.Request.URL.Path
        if fileInfo, err := os.Stat(frontendPath + path); err == nil {
            if !fileInfo.IsDir() {
                c.File(frontendPath + path)
            }
        }
        c.HTML(404, "404.html", nil)
    })
    route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
}
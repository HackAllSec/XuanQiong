
package routes

import (
	"fmt"
	"strings"
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
	fmt.Println("Welcome to VulnData")
	fmt.Println("Server running on " + config.Config.Server.Host + ":" + strconv.Itoa(config.Config.Server.Port))
}

// 前后端分离的路由
func InitRoutes() {
	route.POST("/login", controllers.Login)
	route.GET("/logout", controllers.Logout)
	route.Run(config.Config.Server.Host + ":" + strconv.Itoa(config.Config.Server.Port))
}

// 前后端不分离的路由
func StartServer() {
	// 前端静态文件
	route.Static("/assets", "../frontend/assets")

	// 设置前端模板文件的路由
	route.LoadHTMLGlob("../frontend/*.html")

	// 定义路由
	route.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	route.POST("/login", controllers.Login)
	route.GET("/logout", controllers.Logout)
	route.GET("/test", controllers.Index)

	// 通配符路由
	route.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		// 检查路径是否以 .html 结尾
		if strings.HasSuffix(path, ".html") {
			c.HTML(200, path[1:], nil)
		} else {
			c.HTML(404, "404.html", nil)
		}
	})
	route.Run(config.Config.Server.Host + ":" + strconv.Itoa(config.Config.Server.Port))
}

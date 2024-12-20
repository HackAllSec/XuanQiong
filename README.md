# 简介

## 简体中文 | [English](README_EN.md)

**XuanQiong（玄穹）**——一款高性能的开源漏洞库平台，中小型团队自建漏洞库的合适之选。支持漏洞提交、漏洞审核、漏洞搜索、漏洞排行榜、消息推送等功能。

## 技术栈和功能

### 前端

- Vite
- Vue3
- Element Plus
- Typescript

### 后端

- Gin
- GORM
- JWT

支持MySQL、PostgreSQL、SQLite3、SQL Server等数据库，具体查看GORM支持的数据库。  
支持JWT、Webhook、邮件通知等。  
内置漏洞类型，漏洞评分规则，积分计算规则查看[积分规则](ScoreRules.md)。

演示环境：[https://demo.hackall.cn](https://demo.hackall.cn)  
管理员账号：admin/Admin@123  
普通用户账号：test/123456  

### 用户界面功能

|功能|说明|前端实现情况|后端实现情况|
|-|-|-|-|
|查看漏洞|查看漏洞摘要和详情，分页查看|✅|✅|
|注册|用户注册|✅|✅|
|登录退出|用户登录和退出|✅|✅|
|忘记密码|忘记密码后重置密码|✅|✅|
|我的信息|头像、用户名、邮箱和手机号修改，积分，漏洞提交情况展示|✅|✅|
|修改个人信息|修改用户名、邮箱、手机号、密码|✅|✅|
|消息|查看漏洞审核消息，积分变动消息等|❌|❌|
|我提交的漏洞|查看自己提交的漏洞|✅|✅|
|获取漏洞|获取漏洞摘要、详情|✅|✅|
|提交漏洞|提交漏洞信息|✅|✅|
|编辑漏洞|未审核通过的漏洞信息修改|✅|✅|
|附件上传|上传漏洞信息附件|✅|✅|
|简单搜索|模糊搜索|✅|✅|
|高级搜索|精确搜索|✅|✅|
|排行榜|月度、季度、年度排行榜|✅|✅|
|语言切换|支持中文和英文|✅|✅|

### 管理员界面功能

|功能|说明|前端实现情况|后端实现情况|
|-|-|-|-|
|登录退出|管理员登录和退出|✅|✅|
|忘记密码|忘记密码|✅|✅|
|修改个人信息|修改用户名、邮箱、手机号、密码|✅|✅|
|创建用户|创建管理员或普通用户|✅|✅|
|查看用户|分页查看用户列表|✅|✅|
|修改用户信息|修改管理员或普通用户信息，包含密码、邮箱、手机号、状态等|✅|✅|
|删除用户|删除管理员或普通用户信息|✅|✅|
|仪表盘|展示漏洞总数、Poc总数、Exp总数、最近新增、用户总数、管理员总数、CPU、内存、磁盘使用情况|✅|✅|
|系统设置|功能启停，登录锁定策略配置，邮箱配置，JWT配置、Webhook通知配置|✅|✅|
|漏洞管理|添加、修改和删除漏洞类型，查看、更新和审核漏洞，导入、导出漏洞|✅|❌导入导出待完成|
|评分规则管理|添加、修改和删除评分规则|✅|✅|
|语言切换|支持中文和英文|✅|✅|
|消息推送|支持钉钉和企业微信Webhook通知|❌|❌|

## 目录结构

```
XuanQiong/
│
├── backend/               # 后端目录
│   ├── config/            # 配置解析和验证
│   ├── controllers/       # 控制器，处理路由和逻辑
│   ├── models/            # 数据模型
│   ├── routes/            # 路由定义
│   ├── types/             # 类型定义
│   ├── utils/             # 工具和辅助函数
│
├── frontend/              # 用户界面目录，Vue3+Vite+Element Plus+Typescript
|   ├── dist               # build生成的目录
│   │   └── assets/        # 静态资源，如图片、样式等
│   │   └── index.html     # 用户端入口文件
|
├── admin/                 # 后台管理界面目录，Vue3+Vite+Element Plus+Typescript
│   ├── dist               # build生成的目录
│   │   └── assets/        # 静态资源，如图片、样式等
│   │   └── index.html     # 后台入口文件
|
├── config.yaml            # 配置文件
├── go.mod                 # Go Modules 依赖文件
├── go.sum
├── main.go                # 主入口
└── README.md              # 项目说明文件
```

## 部署方式

支持`前后端一体`和`前后端分离`，已在 MySQL 和 Sqlite 测试，建议使用MySQL。如果遇到问题请查看[FAQ](FAQ.md)。  
**使用 sqlite 时，需要设置环境变量 `CGO_ENABLED=1` 后重新 build，Releases 中的二进制文件使用 `CGO_ENABLED=0` 编译。**  

### 前后端一体化

默认使用一体化模式启动，步骤如下：
- 修改 `config.yaml` 中的数据库配置，修改数据库名称，初始化时自动创建数据库。
- 运行 `go run main.go` 或 Releases 中的二进制文件即可。

启动后随机生成管理员密码

### 前后端分离

用户前端文件和管理员前端文件独立，可分别部署在不同web目录下。后台服务通过 `config.yaml` 中的 `start_mode` 参数切换启动模式。

- 前端配置 `src/api.ts` 中 baseURL 地址，然后编译。
- 编译后的用户前端文件位于：`frontend/dist` 目录下，复制目录下的文件到 web 目录即可。
- 编译后的管理员前端文件位于：`admin/dist` 目录下，复制目录下的文件到 web 目录即可。
- 后端配置CORS、运行模式，修改 `config.yaml` 中的 `start_mode`, `allow_origins`, `allow_methods` 和 `allow_headers` 参数，然后运行即可。

## ChangeLog

[CHANGELOG](CHANGELOG.md)

## API

- [HTML格式](API/XuanQiong.html)
- [Apifox格式](API/XuanQiong.apifox.json)
- [Openapi格式](API/XuanQiong.openapi.json)

## Star History

![](https://api.star-history.com/svg?repos=HackAllSec/XuanQiong&type=Date)

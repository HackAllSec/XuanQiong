# 简介

**XuanQiong（玄穹）**——开源漏洞库平台

## 目录结构

```
XuanQiong/
│
├── backend/               # 后端目录
│   ├── cmd/               # 可执行文件（如果使用 Go Modules）
│   │   └── main.go
│   ├── controllers/       # 控制器，处理路由和逻辑
│   ├── models/           # 数据模型
│   ├── routes/           # 路由定义
│   ├── utils/            # 工具和辅助函数
│   ├── go.mod            # Go Modules 依赖文件
│   ├── go.sum
│   └── main.go            # 后端主入口
│
├── frontend/              # 前端目录
│   ├── assets/           # 静态资源，如图片、样式等
│   ├── index.html
└── README.md            # 项目说明文件
```

## 技术栈

- Gin
- GORM
- JWT

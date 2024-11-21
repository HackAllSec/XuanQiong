# Summary

## English | [简体中文](README.md)

**XuanQiong（玄穹）**——A high-performance open-source vulnerability library platform, suitable for small and medium-sized teams to build their own vulnerability libraries. Support functions such as vulnerability submission, vulnerability review, vulnerability search, vulnerability ranking list, and message push.

## Technology stack and functionality

### Frontend

- Vite
- Vue3
- Element Plus
- Typescript

### Backend

- Gin
- GORM
- JWT

Supports databases such as MySQL, PostgreSQL, SQLite3, SQL Server, etc. Please refer to the databases supported by GORM for details.
Support JWT, Webhook, email notifications, etc.
View the built-in vulnerability types, vulnerability rating rules, and point calculation rules [Score Rules](ScoreRules.md).

Demo：[https://demo.hackall.cn](https://demo.hackall.cn)  
Administrator Account：admin/Admin@123  
Ordinary user accounts：test/123456  

### User UI functions

|Function|Description|Is the front-end implemented|Is the back-end implemented|
|-|-|-|-|
|View Vulnerabilities|View vulnerability summary and details, paginated view|✅|✅|
|Register|User Register|✅|✅|
|Login/Logout|USer Login/Logout|✅|✅|
|Forgot Password|Forgot Password and Reset Password|✅|✅|
|My Profile|Avatar, username, email, and phone number modification, points, vulnerability submission status display|✅|✅|
|Modify personal information|Change username, email, phone number, password|✅|✅|
|Message|View vulnerability review messages, points change messages, etc|❌|❌|
|My vulnerabilities|View vulnerabilities submitted by oneself|✅|✅|
|Get vulnerabilities|Get vulnerability summary and details|✅|✅|
|Submit vulnerabilities|Submit vulnerabilities details|✅|✅|
|Edit vulnerabilities|Unauthorized vulnerability information modification|✅|✅|
|Attachment upload|Upload vulnerability information attachment|✅|✅|
|Simple Search|fuzzy search|✅|✅|
|Advanced search|Accurate search|✅|✅|
|Ranking list|Monthly, quarterly, and annual rankings|✅|✅|
|Language switching|Supports both Chinese and English|✅|✅|

### Admin UI functions

|Function|Description|Is the front-end implemented|Is the back-end implemented|
|-|-|-|-|
|Login/Logout|Admin Login/Logout|✅|✅|
|Forgot Password|Forgot Password|✅|✅|
|Modify personal information|Change username, email, phone number, password|✅|✅|
|Create Users|Create administrator and normal users|✅|✅|
|View Users|Page view user list|✅|✅|
|Modify user information|Modify administrator or regular user information, including password, email, phone number, status, etc|✅|✅|
|Delete Users|Delete administrator or regular user information|✅|❌Multiple delete to be completed|
|Dashboard|Display the total number of vulnerabilities, PoC, Exp, recent additions, total number of users, total number of administrators, CPU, memory, and disk usage|✅|✅|
|System Setting|Function start stop, login lock policy configuration, email configuration, JWT configuration, Webhook notification configuration|✅|✅|
|Vulnerability management|Add, modify, and delete vulnerability types, view, update, and review vulnerabilities, import and export vulnerabilities|✅|❌Import and export to be completed, Multiple delete to be completed|
|Scoring rule management|Add, modify, and delete rating rules|✅|❌Multiple delete to be completed|
|Language switching|Supports both Chinese and English|✅|✅|
|Push messages|Support DingTalk and WxWork webhook notifications|❌|❌|

## Directory structure

```
XuanQiong/
│
├── backend/               # backend source code
│   ├── config/            # Configuration parsing and validation
│   ├── controllers/       # Controller, responsible for handling routes and logic
│   ├── models/            # Models, responsible for database operations
│   ├── routes/            # Routes, responsible for routing
│   ├── types/             # Types, responsible for type definitions
│   ├── utils/             # Utils, responsible for common functions
│
├── frontend/              # User Frontend, Vue3+Vite+Element Plus+Typescript
|   ├── dist               # build
│   │   └── assets/
│   │   └── index.html
|
├── admin/                 # Admin Frontend, Vue3+Vite+Element Plus+Typescript
│   ├── dist               # build
│   │   └── assets/
│   │   └── index.html
|
├── config.yaml            # Configuration file
├── go.mod
├── go.sum
├── main.go                # Main entry point
└── README.md
```

## Deployment method

Support `Front and rear integrated` and `Front and rear separation`.

### Front and rear integrated

Just run `main. go` directly.

### Front and rear separation

User front-end files and administrator front-end files are independent and can be deployed separately in different web directories

- Configure API address on the front-end and then compile it
- The compiled user frontend files are located in the frontend/dist directory. Simply copy the files from the directory to the web directory
- Backend configuration CORS and running mode

## ChangeLog

[CHANGELOG](CHANGELOG.md)

## API

- [HTML](API/XuanQiong.html)
- [Apifox](API/XuanQiong.apifox.json)
- [Openapi](API/XuanQiong.openapi.json)

## Star History

![](https://api.star-history.com/svg?repos=HackAllSec/XuanQiong&type=Date)
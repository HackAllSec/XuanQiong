# Summary

## English | [简体中文](README.md)

**XuanQiong（玄穹）**——A high-performance open-source vulnerability library platform, suitable for small and medium-sized teams to build their own vulnerability libraries. It supports vulnerability submission, vulnerability review, vulnerability search, vulnerability rankings, RBAC permission management, operation audit logs, and brand customization.

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
Support JWT, `X-Auth-Token` application authentication, email captcha, Webhook configuration, etc.
View the built-in vulnerability types, vulnerability rating rules, and point calculation rules [Score Rules](ScoreRules.md).

Demo：[https://demo.hackall.cn](https://demo.hackall.cn)  
Demo credentials are provided by maintainers when needed.  

### User UI functions

|Function|Description|Is the front-end implemented|Is the back-end implemented|
|-|-|-|-|
|View Vulnerabilities|View vulnerability summary and details, paginated view|✅|✅|
|Register|User Register|✅|✅|
|Login/Logout|USer Login/Logout|✅|✅|
|Forgot Password|Forgot Password and Reset Password|✅|✅|
|My Profile|Avatar, username, email, and phone number modification, points, vulnerability submission status display|✅|✅|
|Modify personal information|Change username, email, phone number, password|✅|✅|
|Message|View vulnerability review messages and point-change notifications|✅|✅|
|My vulnerabilities|View vulnerabilities submitted by oneself|✅|✅|
|API keys / personal access tokens|Create, view, and delete personal API keys for automation integrations|✅|✅|
|Get vulnerabilities|Get vulnerability summary and details|✅|✅|
|Submit vulnerabilities|Submit vulnerabilities details|✅|✅|
|Edit vulnerabilities|Unauthorized vulnerability information modification|✅|✅|
|Attachment upload|Upload vulnerability information attachment|✅|✅|
|Attachment download|Download attachments based on vulnerability visibility, submitter, and permissions|✅|✅|
|Simple Search|fuzzy search|✅|✅|
|Advanced search|Accurate search|✅|✅|
|Ranking list|Monthly, quarterly, and annual rankings|✅|✅|
|Language switching|Supports both Chinese and English|✅|✅|
|Brand display|Display configured site name, logo, footer, help URL, and suggestion URL|✅|✅|

### Admin UI functions

|Function|Description|Is the front-end implemented|Is the back-end implemented|
|-|-|-|-|
|Login/Logout|Admin Login/Logout|✅|✅|
|Forgot Password|Forgot Password|✅|✅|
|Modify personal information|Change username, email, phone number, password|✅|✅|
|Create Users|Create users and assign roles|✅|✅|
|View Users|Page view user list|✅|✅|
|Modify user information|Modify user information, including password, email, phone number, status, roles, etc|✅|✅|
|Delete Users|Delete administrator or regular user information|✅|✅|
|Dashboard|Display the total number of vulnerabilities, PoC, Exp, recent additions, total number of users, total number of administrators, CPU, memory, and disk usage|✅|✅|
|System Setting|Feature toggles, login lock policy, email, JWT, and Webhook configuration|✅|✅|
|Brand customization|Configure site name, frontend/admin titles, logo, favicon, footer, help URL, and suggestion URL|✅|✅|
|Role and permission management|Create, update, and delete roles; assign action-level permissions to roles|✅|✅|
|Dynamic permission menus|Admin menus are rendered according to the current user's permissions|✅|✅|
|Operation audit logs|Record login, configuration, user, role, and vulnerability review operations; support search and detail view|✅|✅|
|Vulnerability management|Add, modify, and delete vulnerability types; view, update, review, and delete vulnerabilities|✅|✅|
|Scoring rule management|Add, modify, and delete rating rules|✅|✅|
|Language switching|Supports both Chinese and English|✅|✅|
|Push messages|In-site messages or Webhook pushes triggered by vulnerability review, score changes, and other events|✅|✅|
|Vulnerability import/export|Import unaudited vulnerabilities from CSV and export vulnerability data|✅|✅|
|API key management|Create, view, and delete personal API keys, with `X-API-Key` API access|✅|✅|
|Data backup and restore|Export JSON backups and restore core data by permission|✅|✅|

### Security and operation capabilities

|Function|Description|Status|
|-|-|-|
|Application auth header isolation|The application layer only uses `X-Auth-Token`, avoiding conflicts with reverse proxy Basic Auth in the `Authorization` header|✅|
|Initial administrator password|A random administrator password is generated on first startup, and password change is required after login|✅|
|Forced re-login|Existing tokens are invalidated after password changes or role/permission changes|✅|
|Attachment access control|Attachment downloads are checked against vulnerability visibility, submitter, and management permissions|✅|
|Captcha rate limiting|Captcha requests are protected by email, IP, and global rate limits|✅|
|Sensitive data masking|System configuration, audit logs, and response records mask tokens, passwords, secrets, and similar sensitive data|✅|
|Upload limits|Vulnerability attachments are restricted by type and size, currently up to 10 MB|✅|
|API key authentication|Supports long-lived `X-API-Key` tokens. The server stores only hashes, and plaintext is shown once on creation|✅|
|Vulnerability import/export|Supports admin CSV import/export. Imported data enters the unaudited workflow by default|✅|
|Data backup and restore|Supports admin JSON backup export and restore, protected by the `backup.manage` permission|✅|

### Recently Completed Capabilities

|Function|Current status|Description|
|-|-|-|
|In-site message center|Implemented|A message is created after vulnerability review; the frontend supports listing messages and marking them as read.|
|Event push messages|Implemented|After vulnerability review, DingTalk or WxWork Webhook notifications are sent according to system settings.|
|Vulnerability import/export|Implemented|The admin panel supports CSV import/export, and imported vulnerabilities enter the unaudited state by default.|
|API keys / personal access tokens|Implemented|Users can create, view, and delete personal API keys; API calls use `X-API-Key`.|
|Data backup and restore|Implemented|The admin panel supports JSON backup export and restore, protected by RBAC permissions.|

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

Support `front and backend integrated` and `front and backend separation`. It has been tested in MySQL and SQLite, it is recommended to use MySQL. If you encounter any problems, please refer to the [FAQ](FAQ.md).

When using SQLite, you need to set the environment variable `CGO-ENABLED=1`, and then rebuild. The binary files in Releases are compiled using `CGO-ENABLED=0`.

### Front and backend integrated

Start using this mode by default, the steps are as follows:
- Modify the database configuration in `config. yaml`, change the database name, and automatically create the database during initialization.
- Just run the binary files in releases or `go run main.go`.

After startup, an admin password will be randomly generated. After logging in with that password for the first time, the administrator is forced to change it.

### Front and backend separation

User front-end files and administrator front-end files are independent and can be deployed separately in different web directories. The backend service switches the startup mode through the `start_made` parameter in `config.yaml`.

- Configure the `baseURL` address in `src/api.ts` for the frontend, and then compile it.
- The compiled user frontend files are located in the `frontend/dist` directory. Simply copy the files from the directory to the web directory.
- The compiled admin frontend files are located in the `admin/dist` directory. Simply copy the files from the directory to the web directory.
- Backend configuration CORS and start_mode, modify the `start_made`, `allow_origins`, `allow_methods`, and `allow_headers` parameters in `config.yaml`, and then run it.

## ChangeLog

[CHANGELOG](CHANGELOG.md)

## API

- [HTML](API/XuanQiong.html)
- [Apifox](API/XuanQiong.apifox.json)
- [Openapi](API/XuanQiong.openapi.json)

## Star History

![](https://api.star-history.com/svg?repos=HackAllSec/XuanQiong&type=Date)

# Summary

## English | [简体中文](README.md)

**XuanQiong（玄穹）** is a high-performance open-source vulnerability library platform for small and medium-sized teams that need to build an internal vulnerability knowledge base, retain vulnerability assets, and manage review workflows. It supports vulnerability submission, review, search, rankings, action-level RBAC, operation audit logs, in-site messages, Webhook pushes, API keys, data backup/restore, and brand customization.

## Technology stack and functionality

### Frontend

- Vite
- Vue3
- Element Plus
- TypeScript

### Backend

- Gin
- GORM
- JWT

Supports databases such as MySQL, PostgreSQL, SQLite3, and SQL Server. See the GORM documentation for the full database compatibility list.
Browser application sessions use `X-Auth-Token`, automation integrations use `X-API-Key`, and `Authorization` is reserved for proxy-layer capabilities such as reverse proxy Basic Auth.
See [Score Rules](ScoreRules.md) for built-in vulnerability types, scoring rules, and point calculation rules.

Demo: [https://demo.hackall.cn](https://demo.hackall.cn)  
Administrator Account：admin/Admin@123  
Ordinary user accounts：test/123456  

### v1.1.0 Highlights

- Added action-level RBAC, dynamic permission menus, operation audit logs, and brand customization.
- Added in-site messages, Webhook event pushes, vulnerability CSV import/export, API keys, and data backup/restore.
- Hardened authentication header isolation, attachment access control, captcha rate limiting, audit log redaction, import/export, and backup/restore workflows.
- First startup now generates a random administrator password and requires password change after first login. Fixed default public credentials are no longer documented.

### User UI functions

|Function|Description|Is the front-end implemented|Is the back-end implemented|
|-|-|-|-|
|View Vulnerabilities|View vulnerability summary and details, paginated view|✅|✅|
|Register|User Register|✅|✅|
|Login/Logout|User login and logout|✅|✅|
|Forgot Password|Forgot Password and Reset Password|✅|✅|
|My Profile|Avatar, username, email, and phone number modification, points, vulnerability submission status display|✅|✅|
|Modify personal information|Change username, email, phone number, password|✅|✅|
|Message|View vulnerability review messages and point-change notifications|✅|✅|
|My vulnerabilities|View vulnerabilities submitted by oneself|✅|✅|
|API keys / personal access tokens|Create, view, and delete personal API keys for automation integrations|✅|✅|
|Get vulnerabilities|Get vulnerability summary and details|✅|✅|
|Submit vulnerabilities|Submit vulnerabilities details|✅|✅|
|Edit vulnerabilities|Edit vulnerabilities that have not been approved yet|✅|✅|
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

XuanQiong supports both integrated frontend/backend deployment and separated frontend/backend deployment. It has been tested with MySQL and SQLite; MySQL is recommended for production. If you encounter problems, refer to the [FAQ](FAQ.md).

When using SQLite, set `CGO_ENABLED=1` and rebuild. Release binaries are compiled with `CGO_ENABLED=0`.

### Integrated frontend/backend

This is the default startup mode:
- Modify the database configuration in `config.yaml`; the database is created automatically during initialization when supported by the selected database type.
- Run `go run main.go` or use a release binary.

After startup, an admin password will be randomly generated. After logging in with that password for the first time, the administrator is forced to change it.

### Separated frontend/backend

The user frontend and admin frontend are independent and can be deployed to different web directories. The backend service switches startup mode through the `start_mode` parameter in `config.yaml`.

- Configure the frontend `baseURL` in `src/api.ts`, then build the frontend.
- The compiled user frontend files are located in `frontend/dist`; copy the files to the target web directory.
- The compiled admin frontend files are located in `admin/dist`; copy the files to the target web directory.
- Configure backend CORS and startup mode by setting `start_mode`, `allow_origins`, `allow_methods`, and `allow_headers` in `config.yaml`, then start the backend.

## ChangeLog

[CHANGELOG](CHANGELOG.md)

## API

- [HTML](API/XuanQiong.html)
- [Apifox](API/XuanQiong.apifox.json)
- [Openapi](API/XuanQiong.openapi.json)

## Star History

![](https://api.star-history.com/svg?repos=HackAllSec/XuanQiong&type=Date)

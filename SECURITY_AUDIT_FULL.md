# XuanQiong 全项目安全与缺陷审计报告

## 审计目标

对当前仓库进行尽可能穷尽的源码级安全与缺陷审计，覆盖后端、前端、后台、配置、脚本、路由、权限模型、文件处理、敏感信息处理和测试基础设施。本文档统一记录已确认问题、后续新增问题、覆盖范围和最终结论。

> 说明：软件工程无法通过静态审计证明“绝对不存在任何漏洞”。本报告的目标是穷尽当前源码和可运行验证范围内的可证实问题，并明确列出尚需运行时/部署环境验证的边界。

## 当前审计状态

- 状态：本轮确认问题已完成集中修复并通过验证
- 仓库路径：`/Users/bytedance/XuanQiong`
- 当前已完成：全文件入口枚举、后端路由/控制器/模型/配置/工具复扫、admin 前端复扫、frontend 前端复扫、配置/文档/API 规范复扫、关键安全模式搜索、`go vet` 与前端类型检查复核、并行审计结果去重合并
- 当前确认发现：34 条
- 修复验证：`go test ./...`、`go vet ./...`、admin/frontend `vue-tsc --noEmit`、admin/frontend `pnpm build` 均已通过
- 边界说明：本报告覆盖当前仓库源码与配置中的可证实问题，不等价于真实部署环境渗透测试、黑盒扫描、第三方服务配置审计或生产数据权限核验。

## 已确认发现

### 1. HIGH - RBAC 迁移后旧 `role == 1` 造成后台角色越权读取完整漏洞详情

- 文件：`backend/models/vulnerabilities.go`
- 位置：`GetVulnDetailAuthed` 中 `roleid == 1` 放行逻辑
- 影响：拥有后台入口权限但不应读取漏洞敏感详情的角色，可能读取任意私有或未审核漏洞的 POC、EXP、附件 ID 等完整信息。
- 证据链：
  - 路由 `GET /api/v1/getvulndtl` 对外开放。
  - 控制器在登录态下调用 `models.GetVulnDetailAuthed(id, currentUser.ID, currentUser.Role)`。
  - 模型层使用 `vulnerabilities.UserID == userid || roleid == 1` 判定完整数据访问。
  - RBAC 同步逻辑会把任何拥有 `admin.panel.access` 的用户旧字段同步为 `role = 1`，导致“日志审计员”等后台角色被旧逻辑视为超级管理员。
- 修复建议：移除漏洞详情授权中的旧 `roleid == 1` 判断，改用动作级权限，例如 `vuln.read` / `vuln.audit.read` / `vuln.sensitive.read`，并明确区分公开详情、本人详情、审核详情和全量详情。

### 2. HIGH - RBAC 迁移后旧 `role == 1` 造成任意后台角色可下载任意附件

- 文件：`backend/models/vulnerabilities.go`
- 位置：`CanAccessAttachment` 中 `roleid == 1` 放行逻辑
- 影响：拥有后台入口权限但不应访问附件的角色，可能下载其他用户私有附件、未审核漏洞附件或敏感系统附件。
- 证据链：
  - `/download/file` 为公开路由，登录态下会把当前用户旧 `Role` 传入 `CanAccessAttachment`。
  - `CanAccessAttachment` 对 `roleid == 1` 直接返回附件内容。
  - RBAC 同步逻辑把所有带 `admin.panel.access` 的用户旧字段置为 `role = 1`。
- 修复建议：附件访问控制应基于资源关系和明确权限码，不再使用旧 `role` 字段。公开漏洞附件、本人附件、头像、品牌资源、审核附件应拆分判定。

### 3. MEDIUM - `system.config.read` 可读取明文配置密钥

- 文件：`backend/controllers/sysconfig.go`
- 位置：`GetSystemConfig`
- 影响：只具备系统配置读取权限的角色可获取 `jwt_secret`、`email_password`、通知 `secret`、Webhook 等敏感配置。
- 证据链：
  - 路由 `GET /api/v1/getsysconfig` 仅要求 `system.config.read`。
  - 控制器直接返回 `sysconf`、`emailconf`、`jwtconf`、`noticeconf`。
  - 类型中包含 `jwt_secret`、`email_password`、`secret`、`webhook`。
- 修复建议：读取接口默认脱敏敏感字段；更新接口可接受留空表示不修改，或提供单独的密钥轮换接口。

### 4. MEDIUM - 修改密码页面把密码对象输出到浏览器控制台

- 文件：
  - `admin/src/pages/Modifypasswd.vue`
  - `frontend/src/views/Modifypasswd.vue`
- 位置：`changePassword` 中 `console.log(oldpassword,newpassword,confirmpassword)`
- 影响：用户旧密码、新密码、确认密码可能暴露在浏览器控制台、远程调试、前端日志采集或录屏中。
- 修复建议：删除密码相关调试输出，并检查前端是否还有 token、密钥、完整 payload 输出。

### 5. LOW - 高级搜索 EXP 条件误用 POC 字段

- 文件：`backend/models/vulnerabilities.go`
- 位置：`SearchVulnAdv`
- 影响：用户选择 EXP 过滤时，查询条件包含 `poc IS NOT NULL AND exp <> ''`，会错误排除只有 EXP 没有 POC 的记录。
- 修复建议：改为 `exp IS NOT NULL AND exp <> ''`。

### 6. LOW - 用户漏洞列表 total 与返回列表条件不一致

- 文件：`backend/models/users.go`
- 位置：`GetUserVulnList`
- 影响：`total` 只统计 `status = 1`，但列表返回当前用户全部漏洞，分页总数与数据不一致。
- 修复建议：统计条件与列表条件保持一致，或按业务明确只返回审核通过记录。

### 7. LOW - 删除附件未检查实际删除行数

- 文件：
  - `backend/models/vulnerabilities.go`
  - `backend/controllers/vulnerabilities.go`
- 位置：`DeleteFile`
- 影响：删除不存在或非本人附件时也可能返回成功，导致 UI 和服务端状态不一致。
- 修复建议：检查 `RowsAffected`，不存在或无权限时返回明确错误。

### 8. LOW - 编辑漏洞类型漏 `await` 导致前端编辑流程异常

- 文件：`admin/src/pages/VulnType.vue`
- 位置：`editVulnType`
- 影响：`api.post` 返回 Promise，但代码直接访问 `response.data.code`，运行时会进入异常分支或无法正确刷新列表、关闭弹窗，导致漏洞类型编辑功能不稳定。
- 证据链：
  - 新增漏洞类型、删除漏洞类型等相邻流程均使用 `await api...`。
  - 编辑流程中 `const response = api.post(...)` 未等待请求完成。
  - 后续立即读取 `response.data.code`，此时 `response` 仍是 Promise。
- 修复建议：改为 `const response = await api.post(...)`，并统一业务失败分支处理。

### 9. LOW - 评分规则搜索字段错误导致搜索失效

- 文件：`admin/src/pages/ScoreRule.vue`
- 位置：`filterTableData` computed
- 影响：评分规则列表搜索使用 `item.name`，但评分规则数据字段为 `rule`，输入搜索词后可能触发 `Cannot read properties of undefined`，并被外层 `catch` 吞掉后表现为列表异常或搜索无结果。
- 修复建议：改用实际字段 `item.rule`，并对可空字段使用安全访问。

### 10. MEDIUM - 附件上传缺少大小限制且整文件读入内存和数据库

- 文件：
  - `backend/controllers/vulnerabilities.go`
  - `backend/models/vulnerabilities.go`
- 位置：`UploadFile` / `StoreFile`
- 影响：任意拥有 `attachment.upload` 权限的用户可上传超大文件。后端通过 `io.ReadAll(src)` 将文件完整读入内存，并继续写入数据库 BLOB，可能造成进程内存压力、数据库膨胀和服务不可用。
- 证据链：
  - 路由 `POST /api/v1/upload` 只要求 `attachment.upload`，普通用户默认拥有该权限。
  - 控制器未检查 `file.Size`、Content-Type 白名单或业务配额。
  - 模型层直接 `io.ReadAll(src)` 并保存到 `types.XqAttachment.Data`。
- 修复建议：在控制器和模型层双重限制单文件大小、用户总配额和允许类型；大文件应使用对象存储或流式处理，避免整文件进入内存。

### 11. MEDIUM - 后台创建用户缺少服务端空用户名/空密码校验

- 文件：
  - `backend/controllers/users.go`
  - `backend/models/users.go`
- 位置：`CreateUser` / `CreateUserWithRoles` / `CreateUser`
- 影响：绕过前端表单后，拥有 `user.create` 权限的调用方可创建空用户名或空密码账号。若创建了 `username=""` 且 `password=""` 的启用账号，登录接口可用空用户名和空密码通过校验。
- 证据链：
  - `controllers.CreateUser` 从 JSON 中读取字符串但不校验非空。
  - `models.CreateUser` 只检查用户名重复，不检查 `username == ""` 或 `password == ""`。
  - `models.CheckLogin` 按传入用户名查库并使用 `utils.IsHashEqual` 校验密码，不拒绝空用户名/空密码。
- 修复建议：服务端创建用户必须强制校验用户名、密码、邮箱/手机号格式和密码强度；登录接口也应 fail-fast 拒绝空用户名或空密码。

### 12. HIGH - `EditVuln` 使用旧 `role == 1` 放大“编辑本人漏洞”权限

- 文件：`backend/models/vulnerabilities.go`
- 位置：`EditVuln`
- 影响：拥有 `admin.panel.access` 且同时拥有 `vuln.edit` 的后台角色，会因旧字段同步为 `role = 1` 被模型层视为管理员，从而可编辑任意用户漏洞、绕过“只能编辑本人且审核通过后不可编辑”的限制，并可引用非本人附件。
- 证据链：
  - 路由 `POST /api/v1/editvuln` 只检查权限码 `vuln.edit`，该权限语义为“编辑本人提交的漏洞”。
  - 模型层 `if roleid != 1 { ... }` 只在旧角色非 1 时校验本人和审核状态。
  - RBAC 同步会把拥有 `admin.panel.access` 的用户旧字段置为 `role = 1`。
- 修复建议：废除 `EditVuln` 中的旧 `roleid` 分支，按明确权限码拆分本人编辑、审核编辑、全量编辑，并对附件引用继续做对象级校验。

### 13. LOW - 批量删除用户未防止删除当前登录用户

- 文件：
  - `backend/controllers/users.go`
  - `backend/models/sysconfig.go`
- 位置：`MultiDeleteUsers` / `MultiDelete`
- 影响：单用户删除接口显式阻止删除自己，但批量删除直接调用 `models.MultiDelete("user", ids)`，如果批量列表包含当前用户 ID，会删除当前登录用户，造成会话异常和后台管理中断。
- 修复建议：批量删除用户时过滤并拒绝当前用户 ID，同时限制系统初始管理员或最后一个超级管理员被删除。

### 14. LOW - 漏洞审核积分计算未校验评分规则存在性和类型

- 文件：`backend/models/users.go`
- 位置：`AuditVuln`
- 影响：审核接口接收 `prid`、`erid`、`irid`、`orid` 后仅按 ID 查询评分规则，不校验记录是否存在，也不校验规则类型是否与 POC/EXP/影响面/其他维度匹配。拥有审核权限的用户可以传入错误类型规则或无效 ID，导致积分计算失真。
- 修复建议：分别校验每个评分规则 ID 存在且 `type` 与当前维度匹配；无分数应使用显式哨兵值并在服务端转换，而不是信任前端约定。

### 15. LOW - 启动期自动建库 SQL 拼接未校验数据库名

- 文件：`backend/models/dbo.go`
- 位置：MySQL / PostgreSQL / SQL Server 的 `CREATE DATABASE` 分支
- 影响：数据库名来自配置文件，但启动时被直接拼入 `CREATE DATABASE` 语句。若部署配置被低权限用户、环境注入或供应链流程污染，可能造成启动期 SQL 注入、创建异常数据库名或初始化失败。
- 证据链：
  - MySQL：`fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", config.Config.Database.Connection.Name)`
  - PostgreSQL：`fmt.Sprintf("CREATE DATABASE %s", config.Config.Database.Connection.Name)`
  - SQL Server：`fmt.Sprintf("CREATE DATABASE [%s]", config.Config.Database.Connection.Name)`
- 修复建议：对数据库名做严格标识符白名单校验，例如仅允许 `[A-Za-z0-9_]+`；不同数据库使用对应的安全标识符引用方式。

### 16. LOW - API 文档仍暴露旧鉴权头和敏感示例数据

- 文件：
  - `API/XuanQiong.openapi.json`
  - `API/XuanQiong.apifox.json`
  - `README.md`
  - `README_EN.md`
  - `config.yaml`
- 影响：项目已要求应用层只使用 `X-Auth-Token`，但 API 文档仍大量声明 `Authorization` header，并包含示例 JWT、示例管理员密码、示例 `jwt_secret`、邮件密码和 Webhook 值。部署方或二次开发者可能继续按旧 `Authorization` 契约接入，触发 Basic Auth 冲突或绕开当前安全设计；公开文档中的默认口令说明也会误导部署安全基线。
- 证据链：
  - OpenAPI/Apifox 中多个接口 header 仍为 `Authorization`。
  - `README.md` / `README_EN.md` 仍写有管理员账号 `admin/Admin@123`，但当前代码已随机生成初始密码。
  - `config.yaml` 示例 CORS 头仍包含 `Authorization`，未显式列出 `X-Auth-Token`。
- 修复建议：统一更新 API 文档、README 和示例配置，删除真实形态的 JWT/密钥样例，用占位符替代；文档层面明确应用鉴权只允许 `X-Auth-Token`，`Authorization` 仅保留给代理层 Basic Auth。

### 17. MEDIUM - 后台漏洞列表权限变量作用域错误导致按钮鉴权和上传头失效

- 文件：`admin/src/pages/Vulnlist.vue`
- 位置：模板使用的 `uploadHeaders`、`canEdit`、`canDelete`
- 影响：模板引用的上传头和按钮权限变量没有在 `<script setup>` 顶层定义，而是出现在函数体局部作用域，导致上传组件无法稳定拿到 `X-Auth-Token`，编辑/删除按钮的前端权限展示也可能失效。
- 修复建议：把 `uploadHeaders = getUploadHeaders()`、`canEdit = hasPermission('vuln.edit')`、`canDelete = hasPermission('vuln.delete')` 移到顶层，并以后端 RBAC 为最终准入。

### 18. HIGH - 审核接口业务失败被前端当成审核成功

- 文件：`admin/src/pages/Unaudited.vue`
- 位置：`submitAudit`
- 影响：`/api/v1/auditvuln` 返回 `code = 3` 等业务失败时，前端仍进入成功提示、关闭详情并刷新列表，审核员会误以为漏洞已审核成功，造成审核流程和实际数据状态不一致。
- 证据链：
  - 项目约定 `code == 1` 表示业务成功。
  - 当前代码仅把 `code == 0` 当登录失效，其余所有非 0 返回都执行 `ElMessage.success(t('app.webui.audited'))`。
- 修复建议：仅 `code === 1` 进入成功分支；`0/2/3/其他` 分别处理未登录、输入错误、审核失败和未知错误。

### 19. MEDIUM - 上传成功回调只看 HTTP 2xx 不校验业务成功

- 文件：
  - `admin/src/pages/Profile.vue`
  - `admin/src/pages/Vulnlist.vue`
  - `admin/src/pages/EditVuln.vue`
  - `admin/src/pages/System.vue`
- 位置：各页面 `el-upload` 的 `on-success` 回调
- 影响：Element Plus 的上传成功回调只代表 HTTP 2xx，当前代码直接写入头像、附件 ID、Logo/Favicon 或提示上传成功，未校验后端业务 `code` 和文件 ID。若后端返回业务失败但 HTTP 状态仍为 200，前端状态会被错误更新。
- 修复建议：统一封装上传响应校验，只有 `response.code === 1` 且存在合法文件 ID 时才更新状态，否则清理临时状态并提示失败。

### 20. MEDIUM - 认证失效时手工清理 sessionStorage 不完整导致 RBAC 缓存残留

- 文件：
  - `admin/src/auth.ts`
  - `admin/src/pages/Unaudited.vue`
  - `admin/src/pages/Audited.vue`
  - `admin/src/pages/Profile.vue`
  - `admin/src/pages/Dashboard.vue`
  - `admin/src/pages/ScoreRule.vue`
- 影响：统一清理函数 `clearAuthSession()` 会删除 token、用户、头像、权限、角色和强制改密状态，但多个页面在认证失效时只手工删除 `token/username/avatar`，导致 `permissions`、`roles`、`role_ids`、`force_password_change` 可能残留，出现菜单/按钮权限状态与真实登录态不一致。
- 修复建议：所有认证失效路径统一调用 `clearAuthSession()`；在 `api.ts` 增加响应拦截器集中处理未授权业务码。

### 21. LOW - 后台漏洞等级过滤字段映射错误

- 文件：
  - `admin/src/pages/Vulnlist.vue`
  - `admin/src/pages/Unaudited.vue`
  - `admin/src/pages/Audited.vue`
- 位置：`levelfilterHandler`
- 影响：表格列和后端字段使用 `vuln_level`，过滤函数却读取 `row.level`，导致漏洞等级筛选无效。
- 修复建议：改为 `row.vuln_level === value`，并补充列表过滤用例。

### 22. LOW - 后台页面仍输出敏感业务数据到浏览器控制台

- 文件：
  - `admin/src/pages/Unaudited.vue`
  - `admin/src/pages/Audited.vue`
  - `admin/src/pages/Profile.vue`
- 位置：`console.log(data.value)`、`console.log(index, row)`、`console.log(userinfo.value)`
- 影响：后台漏洞列表元数据、审核行数据和用户个人资料会出现在浏览器控制台，可能被远程调试、录屏、前端日志采集或共享终端泄露。
- 修复建议：删除生产环境 console 输出；必要调试应受环境开关控制并脱敏。

### 23. HIGH - 公开漏洞摘要和搜索接口泄露完整 POC/EXP/附件字段

- 文件：
  - `backend/routes/routes.go`
  - `backend/models/vulnerabilities.go`
- 位置：`GetVulnAbstract` / `SearchVuln` / `SearchVulnAdv`
- 影响：`/api/v1/getvulnabs`、`/api/v1/search`、`/api/v1/advsearch` 均为公开路由，但返回模型可能包含 `poc`、`exp`、`attachment_id`、`attachment_name`、`user_id` 等敏感字段。公开详情接口和公开列表接口已有布尔化/置空处理，说明这些公开搜索/摘要接口存在遗漏。
- 证据链：
  - `GetVulnAbstract` 使用 `Select("xq_vulnerabilities.*, xq_vuln_types.name as vuln_type")`。
  - `SearchVuln` 直接 `Find(&vulnDatas)` 到完整 `XqVulnerability`。
  - `SearchVulnAdv` 拼好过滤条件后直接 `db.Where(query, values...).Find(&vulnDatas)`。
- 修复建议：所有公开接口统一使用 Public DTO，只返回 POC/EXP 布尔标记和公开展示字段，严禁返回正文、附件 ID、提交用户 ID 等敏感字段，并补充公开接口回归测试。

### 24. HIGH - 登录成功响应可能返回旧 Token 或空 Token

- 文件：
  - `backend/models/users.go`
  - `backend/controllers/users.go`
- 位置：`CheckLogin` / `Login`
- 影响：登录时后端生成新 Token 并写入数据库，但返回给前端的是更新前查询出的 `loginUser.Token`。首次登录可能返回空 Token；再次登录可能返回已被新 Token 替换的旧 Token，导致前端保存无效登录态，或出现登录成功但后续接口立即失效的认证故障。
- 证据链：
  - `CheckLogin` 中 `token, _ := utils.GenJWTToken(...)` 后执行 `db.Model(&user).Update("token", token)`。
  - 函数没有执行 `user.Token = token`，也没有检查生成或更新错误。
  - `controllers.Login` 直接返回 `"token": loginUser.Token`。
- 修复建议：检查 Token 生成和数据库更新错误；更新成功后显式把新 Token 写回返回对象，或让 `CheckLogin` 直接返回新 Token 字符串。

### 25. MEDIUM - 审计中间件对 JSON/form 请求体无上限读取

- 文件：
  - `backend/routes/middleware.go`
  - `backend/models/audit_logs.go`
- 位置：`auditLogMiddleware` / `CaptureRequestBody`
- 影响：multipart 请求已经跳过审计读取，但 JSON、form 和空 Content-Type 请求仍通过 `io.ReadAll(request.Body)` 全量读入内存；后续的截断只发生在脱敏输出阶段，不能降低读取阶段的内存压力。攻击者可用大 JSON/form 请求造成内存压力。
- 修复建议：入口层使用 `http.MaxBytesReader` 或有限 reader；审计只读取有限字节并标记 truncated；空 Content-Type 不应默认全量捕获。

### 26. LOW - 编辑漏洞时唯一性校验会被自身 CVE/NVD/CNVD 等编号误杀

- 文件：`backend/models/vulnerabilities.go`
- 位置：`EditVuln` / `checkVulnData`
- 影响：编辑漏洞复用新增漏洞的唯一性校验，查询 CVE/NVD/EDB/CNNVD/CNVD 是否存在时没有排除当前漏洞 ID。已有编号的漏洞只要编辑时保持原编号，就会命中自己并返回“漏洞已存在”。
- 修复建议：拆分新增和编辑校验，或给校验函数传入 `excludeID`，编辑时唯一性查询增加 `id <> ?`。

### 27. MEDIUM - 漏洞编号生成和审核积分更新缺少事务/并发保护

- 文件：
  - `backend/models/vulnerabilities.go`
  - `backend/models/users.go`
- 位置：`getVdbid` / `InsertVuln` / `AuditVuln`
- 影响：漏洞 ID 通过读取最新 ID 后在应用层 `sequence + 1` 生成，并发提交可能生成相同主键。审核流程先读状态再多步更新漏洞、插入积分明细、累加用户积分，没有事务和条件更新，并发审核可能重复发放积分或产生部分写入。
- 修复建议：漏洞 ID 使用数据库序列、唯一约束重试或事务锁；审核流程放入事务，状态更新使用 `WHERE status = 0` 并检查 `RowsAffected`。

### 28. HIGH - 前台 Axios 错误对象可能把 `X-Auth-Token` 输出到控制台

- 文件：
  - `frontend/src/api.ts`
  - `frontend/src/views/Submit.vue`
  - `frontend/src/views/Profile.vue`
  - `frontend/src/branding.ts`
- 影响：请求拦截器会给请求注入 `X-Auth-Token`，而多个 catch 分支直接 `console.error(error)`。Axios error 对象通常包含 `config.headers`，因此请求失败时可能把 Token 输出到浏览器控制台或前端日志采集。
- 修复建议：删除直接输出完整 error 对象的日志；统一错误处理时只记录脱敏后的 `status`、业务码和 message，显式剔除 headers/token。

### 29. MEDIUM - 前台缺少统一未授权响应处理，服务端撤销 Token 后本地登录态可能残留

- 文件：
  - `frontend/src/api.ts`
  - `frontend/src/utils.ts`
  - `frontend/src/views/Profile.vue`
- 影响：前台只有请求拦截器，没有响应拦截器统一处理 HTTP 401/403 或业务未授权码。页面多处只依赖本地 JWT exp 或局部 `code == 0` 分支，服务端撤销 Token、改密清 Token、角色变更清 Token 后，前端可能保留旧 sessionStorage 状态直到某个页面局部处理。
- 修复建议：在 axios response interceptor 中统一处理 401/403 和项目未授权业务码，调用统一会话清理函数并跳转登录。

### 30. LOW - 前台首页/提交页直接访问时 `redirectedFrom.path` 空引用崩溃

- 文件：
  - `frontend/src/views/Index.vue`
  - `frontend/src/views/Submit.vue`
- 影响：页面 mounted 流程直接读取 `router.redirectedFrom.path` 或 `route.redirectedFrom.path`。直接打开 `/#/`、`/#/submit` 或刷新页面时 `redirectedFrom` 可能为空，导致运行时异常。
- 修复建议：使用 `route.redirectedFrom?.path` 并把“从我的漏洞编辑进入”的状态改为显式 query/state。

### 31. LOW - 前台 EDB ID 字段写入 `edbid` 导致提交字段丢失

- 文件：`frontend/src/views/Submit.vue`
- 位置：EDB ID 输入框和提交表单
- 影响：表单初始化和后端字段使用 `edb`，但模板输入绑定到 `form.edbid`，用户填写的 EDB ID 可能以错误字段提交并被后端忽略，详情页也读取不到该值。
- 修复建议：统一字段名为 `form.edb`，并补充提交表单字段映射检查。

### 32. MEDIUM - 前台编辑漏洞将完整漏洞详情持久化到 localStorage

- 文件：
  - `frontend/src/views/Myvulns.vue`
  - `frontend/src/views/Submit.vue`
- 影响：编辑漏洞时把 `vulndetail.value.data` 整体写入 `localStorage`，其中可能包含 POC、EXP、附件 ID 等敏感漏洞详情。`localStorage` 持久化时间长，且同源脚本均可读取。
- 修复建议：改用内存状态、短生命周期 `sessionStorage` 或服务端重新拉取；仅传递必要 ID，避免把完整漏洞对象持久化到浏览器。

### 33. LOW - 前台提交漏洞缺少重复提交锁

- 文件：`frontend/src/views/Submit.vue`
- 影响：提交按钮和 `onSubmit` 没有 `loading/disabled` 或 in-flight guard，连续点击会并发请求 `/api/v1/addvuln` 或 `/api/v1/editvuln`，可能造成重复提交、错误提示覆盖或状态混乱。
- 修复建议：增加 `submitting` 状态并在 `finally` 中复位；后端对新增漏洞也应结合唯一约束重试或幂等保护。

### 34. LOW - 前台个人资料更新输出完整用户信息

- 文件：`frontend/src/views/Profile.vue`
- 位置：`modifyUserInfo`
- 影响：`userinfo` 来自 `/api/v1/userinfo`，包含 username、email、phone 等个人资料，提交前直接输出到控制台，可能被远程调试、录屏或日志采集泄露。
- 修复建议：删除该 console；如需调试，必须脱敏并受环境开关控制。

## 测试补齐记录

本轮已新增测试覆盖：

- 审计响应体敏感字段脱敏。
- multipart 请求不被审计中间件读取请求体。
- 注册用户默认获得普通用户角色。
- 后台创建用户无显式角色时默认获得普通用户角色。
- 历史管理员迁移为超级管理员并清空旧 token。
- 系统配置零值可以保存。
- 外部 `Authorization: Bearer` 被拒绝。
- Nginx Basic Auth 的 `Authorization: Basic` 被保留。
- `X-Auth-Token` 正常归一化为内部 Bearer。

验证命令已通过：

```bash
go test -v -gcflags='all=-l -N' ./backend/models
go test -v -gcflags='all=-l -N' ./backend/routes
go test ./...
PATH=/usr/local/bin:$PATH /usr/local/bin/corepack pnpm build
```

其中 `pnpm build` 分别在 `admin` 与 `frontend` 目录执行，均通过，仅存在 Vite 大 chunk 警告。

## 本轮审计清单

- [x] 枚举全项目源码与入口文件
- [x] 后端路由与控制器完整复扫
- [x] 后端模型/数据库/权限/文件处理完整复扫
- [x] 后端配置、初始化、日志、脚本复扫
- [x] admin 前端完整复扫
- [x] frontend 前端完整复扫
- [x] SQL/命令/路径/XSS/敏感信息/权限关键字全量模式搜索
- [x] 依赖与构建配置复核
- [x] 去重、定级、输出最终结论

## 最终结论

本轮审计共确认 34 条问题，其中 HIGH 7 条、MEDIUM 11 条、LOW 16 条。优先修复建议按以下顺序处理：

1. 认证与敏感数据泄露：公开接口泄露 POC/EXP、登录返回旧 Token、前端错误日志泄露 Token、系统配置明文密钥。
2. RBAC/对象级权限：旧 `role == 1` 在漏洞详情、附件下载、漏洞编辑中的权限放大。
3. 可用性与一致性：上传/审计无上限读取、并发漏洞编号/审核积分、审核失败前端显示成功。
4. 其余前后端字段映射、状态清理、重复提交、文档配置漂移等 LOW/MEDIUM 问题。

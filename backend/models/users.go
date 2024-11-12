package models

import (
    "fmt"
    "math"
    "strings"
    "strconv"
    "time"
    "mime/multipart"

    "xuanqiong/backend/types"
    "xuanqiong/backend/utils"
)

// 检查登录凭据
func CheckLogin(username string, password string) *types.XqUser {
    var user types.XqUser
    var jwt types.XqJwtConfig
    db.First(&jwt)
    res := db.Raw("SELECT * FROM xq_users WHERE username = ? AND status = 1", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    if user.Status != 1 {
        return nil
    }
    if utils.IsHashEqual(user.Password, password) {
        expires_in := time.Now().Add(time.Second * time.Duration(jwt.JwtExpires)).Unix()
        token, _ := utils.GenJWTToken(user.Username, user.Role, expires_in, jwt.JwtSecret)
        db.Model(&user).Update("token", token)
        return &user
    }
    return nil
}

// 根据用户名获取用户信息
func GetUserByUsername(username string) *types.XqUser {
    var user types.XqUser
    res := db.Raw("SELECT * FROM xq_users WHERE username = ? AND status = 1", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    return &user
}

// 根据token获取用户信息
func GetUserByToken(token string) *types.XqUser {
    var user types.XqUser
    var jwt types.XqJwtConfig
    db.First(&jwt)
    if token != "" {
        token = strings.TrimPrefix(token, "Bearer ")
        res := db.Raw("SELECT * FROM xq_users WHERE token = ? AND status = 1", token).Scan(&user).RowsAffected
        if res == 0 {
            return nil
        }
        claims, _ := utils.DecryptJWTToken(token, jwt.JwtSecret)
        if claims != nil {
            return &user
        }
    }
    return nil
}

// 清除token
func CleanToken(username string) error {
    var user types.XqUser
    res := db.Raw("SELECT * FROM xq_users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        return nil
    }
    err := db.Model(&user).Update("token", "").Error
    if err != nil {
        return err
    }
    return nil
}

// 创建用户
func CreateUser(username string, password string, email string, phone string, role int64) error {
    var user types.XqUser
    res := db.Raw("SELECT * FROM xq_users WHERE username = ?", username).Scan(&user).RowsAffected
    if res == 0 {
        passwdHash := utils.GenPasswordHash(password)
        userData := types.XqUser{
            Username:   username,
            Password:   passwdHash,
            Email:      email,
            Phone:      phone,
            Role:       role,
            CreateTime: time.Now(),
            Status:     1,
        }
        db.Create(&userData)
        return nil
    } else {
        return fmt.Errorf("User %s already exists.", username)
    }
}

// 删除用户
func DeleteUser(userid uint64) error {
    var user types.XqUser
    res := db.Where("id = ?", userid).Find(&user).RowsAffected
    if res == 0 {
        return fmt.Errorf("User not found.")
    }
    return db.Delete(&user).Error
}

// 修改用户信息
func UpdateUser(userid uint64, role int64, username string, password string, email string, phone string, status int64) error {
    var user types.XqUser
    res := db.Where("username = ?", username).Where("id <> ?", userid).Find(&user).RowsAffected
    if res != 0 {
        return fmt.Errorf("Username %s already exists.", username)
    }
    updates := make(map[string]interface{})
    db.Where("id = ?", userid).First(&user)
    if user.Role != role {
        updates["role"] = role
    }
    if username != "" && user.Username != username {
        updates["username"] = username
    }
    if password != "" {
        password = utils.GenPasswordHash(password)
        updates["password"] = password
    }
    if email != "" && user.Email != email && IsEmailValid(email) {
        updates["email"] = email
    }
    if phone != "" && user.Phone != phone {
        updates["phone"] = phone
    }
    if user.Status != status {
        updates["status"] = status
    }
    updates["update_time"] = time.Now()
    db.Model(&user).Where("id = ?", userid).Updates(updates)
    return nil
}

// 分页获取所有用户
func GetUsers(page string, pageSize string) (int64, []types.XqUser) {
    var users []types.XqUser
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&users).Count(&totalCount)
    db.Select("id, username, role, email, phone, create_time, status").
    Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).
    Find(&users)
    //db.Select("id, username, role, create_time, status").Find(&users)
    return totalCount, users
}

// 修改头像
func UpdateAvatar(file *multipart.FileHeader, userid uint64) string {
    var user types.XqUser
    db.Where("id = ? AND status = 1", userid).First(&user)
    fileid, err := StoreFile(file, userid)
    if err != nil {
        return ""
    }
    if user.Avatar != "" {
        DeleteFile(user.Avatar, userid)
    }
    db.Model(&user).Update("avatar", fileid)
    return fileid
}

// 获取用户漏洞情况
func GetUservulns(userid uint64) (int64, int64, []types.XqVulnerability) {
    var totalCount int64
    var authCount int64
    var vulnDatas []types.XqVulnerability
    db.Model(&vulnDatas).Where("user_id = ?", userid).Count(&totalCount)
    db.Model(&vulnDatas).Where("user_id = ? AND status = 1", userid).Count(&authCount)
    db.Model(&vulnDatas).Where("user_id = ?", userid).Order("id desc").Find(&vulnDatas)
    return totalCount, authCount, vulnDatas
}

// 修改用户个人信息
func UpdateUserInfo(userid uint64, username string, email string, phone string) error {
    var user types.XqUser
    updates := make(map[string]interface{})
    res := db.Where("id = ? AND status = 1", userid).First(&user)
    if res.RowsAffected != 1 {
        return fmt.Errorf("User is not exits.")
    }
    if username != "" {
        res = db.Where("username = ? AND id != ?", username, userid).First(&user)
        if res.RowsAffected != 0 {
            return fmt.Errorf("Username is already in use.")
        }
        updates["username"] = username
    }
    if email != "" {
        if IsEmailValid(email) == false {
            return fmt.Errorf("Invalid email address.")
        }
        res = db.Where("email = ? AND id != ?", email, userid).First(&user)
        if res.RowsAffected != 0 {
            return fmt.Errorf("Email is already in use.")
        }
        updates["email"] = email
    }
    if phone != "" {
        updates["phone"] = phone
    }
    updates["update_time"] = time.Now()
    db.Model(&user).Where("id = ?", userid).Updates(updates)
    return nil
}

// 修改用户密码
func UpdateUserPassword(userid uint64, oldpassword string, newpassword string) error {
    var user types.XqUser
    updates := make(map[string]interface{})
    res := db.Where("id = ? AND status = 1", userid).First(&user)
    if res.RowsAffected != 1 {
        return fmt.Errorf("User is not exits.")
    }
    if oldpassword == "" || newpassword == "" {
        return fmt.Errorf("oldpassword and newpassword cannot be empty")
    }
    if !utils.IsHashEqual(user.Password, oldpassword) {
        return fmt.Errorf("Old password is incorrect.")
    }
    newpassword = utils.GenPasswordHash(newpassword)
    updates["password"] = newpassword
    updates["update_time"] = time.Now()
    db.Model(&user).Where("id = ?", userid).Updates(updates)
    return nil
}

// 用户注册
func Register(username string, password string, email string, phone string) int64 {
    var user types.XqUser
    res := db.Where("username = ?", username).First(&user)
    if res.RowsAffected != 0 {
        return 2
    }
    if username == "" || password == "" {
        return 3
    }
    if email != "" {
        if IsEmailValid(email) == false {
            return 4
        }
        res = db.Where("email = ?", email).First(&user)
        if res.RowsAffected != 0 {
            return 5
        }
    }
    userData := types.XqUser{
        Username: username,
        Password: utils.GenPasswordHash(password),
        Email:    email,
        Phone:    phone,
        Role:     0,
        Status:   1,
        CreateTime: time.Now(),
    }
    db.Create(&userData)
    return 1
}

// 获取用户提交的漏洞
func GetUserVulnList(userid uint64, page string, pageSize string) (int64, []types.XqVulnerability) {
    var vulnDatas []types.XqVulnerability
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&vulnDatas).Where("user_id = ? AND status = 1", userid).Count(&totalCount)
    db.Select("id, vuln_name, vuln_type, vuln_level, cvss, is_public, status, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
    Where("user_id = ?", userid).Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Order("create_time DESC").
    Omit("user_id, attachment_id, attachment_name, update_time").Find(&vulnDatas)
    return totalCount, vulnDatas
}

// 审核漏洞-管理员
func AuditVuln(vulnid string, status int64, review string, cvss float64, prid uint64, erid uint64, irid uint64, orid uint64) error {
    var vuln types.XqVulnerability
    updates := make(map[string]interface{})
    res := db.Where("id = ?", vulnid).First(&vuln)
    if res.RowsAffected != 1 {
        return fmt.Errorf("Vuln is not exits.")
    }
    if vuln.Status != 0 {
        return fmt.Errorf("Vuln has been audited.")
    }
    if status == 1 {
        var scoreRule types.XqScoreRule
        var vulnScore int64
        updates["status"] = 1
        if cvss != vuln.CVSS {
            updates["cvss"] = cvss
        }
        fmt.Println(cvss)
        if cvss >0 && cvss <= 3.9 {
            updates["vuln_level"] = "Low"
        } else if cvss >= 4 && cvss <= 6.9 {
            updates["vuln_level"] = "Medium"
        } else if cvss >= 7 && cvss <= 8.9 {
            updates["vuln_level"] = "High"
        } else if cvss >= 9 && cvss <= 10 {
            updates["vuln_level"] = "Critical"
        } else {
            return fmt.Errorf("Invalid cvss")
        }
        vulnScore = int64(math.Round(cvss * 10))
        db.Where("id = ?", prid).First(&scoreRule)
        pocScore := int64(math.Round(scoreRule.Score * scoreRule.Coefficient))
        db.Where("id = ?", erid).First(&scoreRule)
        expScore := int64(math.Round(scoreRule.Score * scoreRule.Coefficient))
        db.Where("id = ?", irid).First(&scoreRule)
        incidenceScore := int64(math.Round(scoreRule.Score * scoreRule.Coefficient))
        db.Where("id = ?", orid).First(&scoreRule)
        otherScore := int64(math.Round(scoreRule.Score * scoreRule.Coefficient))
        totalScore := vulnScore + pocScore + expScore + incidenceScore + otherScore
        // 更新漏洞信息
        updates["update_time"] = time.Now()
        db.Model(&vuln).Where("id = ?", vulnid).Updates(updates)
        // 插入积分明细表
        rankdetail := types.XqRankingDetail{
            UserID:      vuln.UserID,
            VulnID:      vulnid,
            Ranking:     totalScore,
            CreateTime:  time.Now(),
        }
        db.Create(&rankdetail)
        // 更新用户总积分
        var user types.XqUser
        db.Where("id = ?", vuln.UserID).First(&user)
        updates = make(map[string]interface{})
        updates["ranking"] = user.Ranking + totalScore
        updates["update_time"] = time.Now()
        db.Model(&user).Where("id = ?", vuln.UserID).Updates(updates)
        return nil
    } else if status == 2 {
        updates["status"] = 2
        updates["review_comments"] = review
        updates["update_time"] = time.Now()
        db.Model(&vuln).Where("id = ?", vulnid).Updates(updates)
        return nil
    } else {
        return fmt.Errorf("Invalid status")
    }
}

// 获取用户积分Top10
func GetUserTop10() ([]map[string]interface{}, []map[string]interface{}, []map[string]interface{}) {
    year := int(time.Now().Year())
    // 构建查询条件，获取年度积分排名前10的用户ID
    ystart := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
    yend := time.Date(year + 1, time.January, 1, 0, 0, 0, 0, time.UTC)
    yres := getRankingTop(ystart, yend)
    qstart, qend := getCurrentQuarter()
    qres := getRankingTop(qstart, qend)
    mstart, mend := getCurrentMonth()
    mres := getRankingTop(mstart, mend)
    if yres == nil {
        yres = []map[string]interface{}{}
    }
    if qres == nil {
        qres = []map[string]interface{}{}
    }
    if mres == nil {
        mres = []map[string]interface{}{}
    }
    return yres, qres, mres
}

// 获取Ranking Top10
func getRankingTop(start time.Time, end time.Time) []map[string]interface{} {
    var users []map[string]interface{}
    // 使用GORM的联表查询和聚合函数
    db.Table("xq_users u").
        Select("u.username, u.avatar, COALESCE(SUM(rd.ranking), 0) as ranking").
        Joins("left join xq_ranking_details rd on u.id = rd.user_id").
        Where("rd.create_time BETWEEN ? AND ?", start, end).
        Where("u.role <> ?", 1).
        Group("u.id").
        Order("ranking DESC").
        Limit(10).
        Scan(&users)

    return users
}

// 获取评分规则
func GetScoreRules() map[string]interface{} {
    var allRules []types.XqScoreRule
    db.Find(&allRules) // 一次性查询所有规则

    pocRuleList := getRuleList(filterRulesByType(allRules, 1))
    expRuleList := getRuleList(filterRulesByType(allRules, 2))
    incidenceRuleList := getRuleList(filterRulesByType(allRules, 3))
    otherRuleList := getRuleList(filterRulesByType(allRules, 4))

    // 将所有规则列表包装在一个map[string]interface{}中
    return map[string]interface{}{
        "pocrules":    pocRuleList,
        "exprules":    expRuleList,
        "incidencerules": incidenceRuleList,
        "otherrules":  otherRuleList,
    }
}

func getRuleList(rules []types.XqScoreRule) []map[string]interface{} {
    if rules == nil {
        return []map[string]interface{}{}
    }
    var ruleList []map[string]interface{}
    for _, rule := range rules {
        ruleList = append(ruleList, map[string]interface{}{
            "id":       rule.ID,
            "rule":     rule.Rule,
            //"score":    rule.Score,
            //"coefficient": rule.Coefficient,
        })
    }
    return ruleList
}

func filterRulesByType(rules []types.XqScoreRule, ruleType int64) []types.XqScoreRule {
    var filteredRules []types.XqScoreRule
    for _, rule := range rules {
        if rule.Type == ruleType {
            filteredRules = append(filteredRules, rule)
        }
    }
    return filteredRules
}
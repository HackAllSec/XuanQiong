package controllers

import (
    "xuanqiong/backend/types"
    "xuanqiong/backend/models"
    "github.com/gin-gonic/gin"
)

// 获取漏洞类型列表
func GetVulnTypeList(c *gin.Context) {
    res := models.GetVulnTypeList()
    c.JSON(200, gin.H{"data": res})
}

// 分页获取漏洞类型列表
func GetVulnType(c *gin.Context) {
    page := c.Query("page")
    limit := c.Query("limit")
    total, data := models.GetVulnType(page, limit)
    c.JSON(200, gin.H{"total": total, "data": data})
}

// 获取漏洞摘要，无需登录
func GetVulnAbstract(c *gin.Context) {
    total, hasPoc, hasExp, affectedProduct, weeklyAdditionsVuln, weeklyAdditionsPoc, weeklyAdditionsExp, weeklyAdditionsProduct, res := models.GetVulnAbstract()
    c.JSON(200, gin.H{"total": total, "hasPoc": hasPoc, "hasExp": hasExp,
        "affectedProduct": affectedProduct, "weeklyAdditionsVuln": weeklyAdditionsVuln,
        "weeklyAdditionsPoc": weeklyAdditionsPoc, "weeklyAdditionsExp": weeklyAdditionsExp,
        "weeklyAdditionsProduct": weeklyAdditionsProduct, "data": res})
}

// 分页获取漏洞列表，无需登录
func GetVulnList(c *gin.Context) {
    page := c.Query("page")
    limit := c.Query("limit")
    total, data := models.GetVulnList(page, limit)
    c.JSON(200, gin.H{"total": total, "data": data})
}

// 分页获取待审核漏洞列表
func GetUnauditList(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        page := c.Query("page")
        limit := c.Query("limit")
        total, data := models.GetUnauditList(page, limit)
        c.JSON(200, gin.H{"code": 1, "total": total, "data": data})
    } else {
        c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
    }
}

// 分页获取已审核漏洞列表
func GetAuditedList(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        page := c.Query("page")
        limit := c.Query("limit")
        total, data := models.GetAuditedList(page, limit)
        c.JSON(200, gin.H{"code": 1, "total": total, "data": data})
    } else {
        c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
    }
}

// 获取漏洞详细信息，登录和未登录情况
func GetVulnDetail(c *gin.Context) {
    id := c.Query("id")
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil{
        res := models.GetVulnDetailAuthed(id, currentUser.ID, currentUser.Role)
        c.JSON(200, gin.H{"code": 1, "data": res})
        return
    }
    res := models.GetVulnDetail(id)
    c.JSON(200, gin.H{"code": 0, "data": res})
}

// 添加漏洞信息
func AddVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var vulnerabilities types.XqVulnerability
        if err := c.ShouldBindJSON(&vulnerabilities); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input:" + err.Error()})
            return
        }
        vulnerabilities.UserID = currentUser.ID
        err := models.InsertVuln(vulnerabilities)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Submit successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 编辑漏洞
func EditVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var vulnerabilities types.XqVulnerability
        if err := c.ShouldBindJSON(&vulnerabilities); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input:" + err.Error()})
            return
        }
        err := models.EditVuln(vulnerabilities, currentUser.ID, currentUser.Role)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Edit successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 删除漏洞
func DeleteVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        id, _ := data["id"].(string)
        err := models.DeleteVuln(id)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete type successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 批量删除漏洞
func MultiDeleteVulns(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        ids, _ := data["ids"].([]interface{})
        err := models.MultiDelete("vuln", ids)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 搜索漏洞
func SearchVuln(c *gin.Context) {
    keyword := c.Query("keyword")
    page := c.Query("page")
    limit := c.Query("limit")
    total, data := models.SearchVuln(keyword, page, limit)
    c.JSON(200, gin.H{"total": total, "data": data})
}

// 高级搜索
func SearchVulnAdv(c *gin.Context) {
    var data map[string]interface{}
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(200, gin.H{"code": 2, "msg": "Invalid Input"})
        return
    }
    total, res := models.SearchVulnAdv(data)
    c.JSON(200, gin.H{"total": total, "data": res})
}

// 上传文件
func UploadFile(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        res, err := models.StoreFile(file, currentUser.ID)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "File uploaded successfully", "file_id": res})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 获取文件内容
func DownloadFile(c *gin.Context) {
    fileId := c.Query("id")
    file, err := models.GetFileContent(fileId)
    if err != nil {
        c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
        return
    }
    c.Header("Content-Type", file.Type)
    c.Header("Content-Disposition", "attachment; filename="+file.Name)
    c.Data(200, "OK", file.Data)
}

// 管理员删除文件
func DeleteFile(c *gin.Context) {
    fileId := c.Query("id")
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        err := models.DeleteFile(fileId, currentUser.ID)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "File deleted successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 添加漏洞类型
func AddVulnType(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(200, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        name, _ := data["name"].(string)
        err := models.AddVulnType(name)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Add Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 更新漏洞类型
func UpdateVulnType(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var vulntype types.XqVulnType
        if err := c.ShouldBindJSON(&vulntype); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        err := models.UpdateVlunType(vulntype.ID, vulntype.Name)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Update Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 删除漏洞类型
func DeleteVulnType(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        id, _ := data["id"].(float64)
        err := models.DeleteVulnType(uint64(id))
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 批量删除漏洞类型
func MultiDeleteVulnTypes(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        ids, _ := data["ids"].([]interface{})
        err := models.MultiDelete("vulntype", ids)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 获取全部评分规则
func GetAllScoreRules(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        rules := models.GetAllScoreRules()
        c.JSON(200, gin.H{"code": 1, "data": rules})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 分页获取评分规则
func GetScoreRules(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        page := c.Query("page")
        limit := c.Query("limit")
        total, data := models.GetScoreRules(page, limit)
        c.JSON(200, gin.H{"code": 1, "total": total, "data": data})
    } else {
        c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
    }
}

// 添加评分规则
func AddScoreRule(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        stype, _ := data["type"].(float64)
        rule, _ := data["rule"].(string)
        score, _ := data["score"].(float64)
        coefficient, _ := data["coefficient"].(float64)
        err := models.AddScoreRule(int64(stype), rule, score, coefficient)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Add Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 编辑评分规则
func EditScoreRule(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var scorerule types.XqScoreRule
        if err := c.ShouldBindJSON(&scorerule); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
            return
        }
        err := models.EditScoreRule(scorerule)
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Edit Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 删除评分规则
func DeleteScoreRule(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid Input"})
        }
        id, _ := data["id"].(float64)
        err := models.DeleteScoreRule(uint64(id))
        if err != nil {
            c.JSON(400, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission Denied"})
}

// 批量删除评分规则
func MultiDeleteScoreRules(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil && currentUser.Role == 1 {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
            return
        }
        ids, _ := data["ids"].([]interface{})
        err := models.MultiDelete("scorerule", ids)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Delete Successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}
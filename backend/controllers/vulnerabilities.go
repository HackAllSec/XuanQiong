package controllers

import (
    "xuanqiong/types"
    "xuanqiong/models"
    "github.com/gin-gonic/gin"
)

// 获取漏洞类型列表
func GetVulnTypeList(c *gin.Context) {
    res := models.GetVulnTypeList()
    c.JSON(200, gin.H{"data": res})
}

// 获取漏洞摘要，无需登录
func GetVulnAbstract(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil{
        total, hasPoc, hasExp, affectedProduct, weeklyAdditionsVuln, weeklyAdditionsPoc, weeklyAdditionsExp, weeklyAdditionsProduct, res := models.GetVulnAbstract(true)
        c.JSON(200, gin.H{"total": total, "hasPoc": hasPoc, "hasExp": hasExp,
        "affectedProduct": affectedProduct, "weeklyAdditionsVuln": weeklyAdditionsVuln,
        "weeklyAdditionsPoc": weeklyAdditionsPoc, "weeklyAdditionsExp": weeklyAdditionsExp,
        "weeklyAdditionsProduct": weeklyAdditionsProduct, "data": res})
    } else {
        total, hasPoc, hasExp, affectedProduct, weeklyAdditionsVuln, weeklyAdditionsPoc, weeklyAdditionsExp, weeklyAdditionsProduct, res := models.GetVulnAbstract(false)
        c.JSON(200, gin.H{"total": total, "hasPoc": hasPoc, "hasExp": hasExp,
        "affectedProduct": affectedProduct, "weeklyAdditionsVuln": weeklyAdditionsVuln,
        "weeklyAdditionsPoc": weeklyAdditionsPoc, "weeklyAdditionsExp": weeklyAdditionsExp,
        "weeklyAdditionsProduct": weeklyAdditionsProduct, "data": res})
    }
}

// 分页获取漏洞列表，无需登录
func GetVulnList(c *gin.Context) {
    page := c.Query("page")
    limit := c.Query("limit")
    total, data := models.GetVulnList(page, limit)
    c.JSON(200, gin.H{"total": total, "data": data})
}

// 获取漏洞详细信息，登录和未登录情况
func GetVulnDetail(c *gin.Context) {
    id := c.Query("id")
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil{
        res := models.GetVulnDetailAuthed(id, currentUser.ID)
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
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input:" + err.Error()})
            return
        }
        vulnerabilities.UserID = currentUser.ID
        vulnerabilities.Submitter = currentUser.Username
        err := models.InsertVuln(vulnerabilities)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Submit successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 编辑漏洞
func EditVuln(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        var vulnerabilities types.XqVulnerability
        if err := c.ShouldBindJSON(&vulnerabilities); err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input:" + err.Error()})
            return
        }
        err := models.EditVuln(vulnerabilities, currentUser.ID)
        if err != nil {
            c.JSON(200, gin.H{"code": 3, "msg": err.Error()})
            return
        }
        c.JSON(200, gin.H{"code": 1, "msg": "Edit successfully"})
        return
    }
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

// 搜索漏洞
func SearchVuln(c *gin.Context) {
    keyword := c.Query("keyword")
    res := models.SearchVuln(keyword)
    c.JSON(200, gin.H{"msg": res})
}

// 高级搜索
func SearchVulnAdv(c *gin.Context) {
    var data map[string]interface{}
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(200, gin.H{"code": 2, "msg": "Invalid input"})
        return
    }
    res := models.SearchVulnAdv(data)
    c.JSON(200, gin.H{"code": 1, "msg": res})
}

// 上传文件
func UploadFile(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(400, gin.H{"code": 2, "msg": "Invalid input"})
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
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
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
    c.JSON(200, gin.H{"code": 0, "msg": "Permission denied"})
}

package models

import (
    "io"
    "mime/multipart"
    "fmt"
    "strings"
    "strconv"
    "time"
    "xuanqiong/types"
    "xuanqiong/utils"
)

// 获取漏洞摘要
func GetVulnAbstract() (int64, int64, int64, int64, int64, int64, int64, int64, []types.XqVulnerability) {
    var vulnDatas []types.XqVulnerability
    var totalCount int64
    var pocCount int64
    var expCount int64
    var affectedProductCount int64
    var weeklyCount int64
    var weeklyPocCount int64
    var weeklyExpCount int64
    var weeklyAffectedProductCount int64
    db.Model(&vulnDatas).Where("status = 1").Count(&totalCount)
    db.Raw("SELECT COUNT(*) FROM xq_vulnerabilities WHERE poc <> '' AND status = 1").Scan(&pocCount)
    db.Raw("SELECT COUNT(*) FROM xq_vulnerabilities WHERE exp <> '' AND status = 1").Scan(&expCount)
    db.Model(&vulnDatas).Where("affected_product <> '' AND status = 1").Group("affected_product").Count(&affectedProductCount)
    thisWeek := time.Now().UTC().Truncate(24 * 7 * time.Hour)
    db.Model(&vulnDatas).Where("create_time >= ? AND status = 1", thisWeek).Count(&weeklyCount)
    db.Model(&vulnDatas).Where("create_time >= ? AND poc <> '' AND status = 1", thisWeek).Count(&weeklyPocCount)
    db.Model(&vulnDatas).Where("create_time >= ? AND exp <> '' AND status = 1", thisWeek).Count(&weeklyExpCount)
    subQuery := db.Model(&vulnDatas).Where("create_time < ? AND affected_product = v.affected_product AND status = 1", thisWeek)
    db.Table("xq_vulnerabilities as v").
        Select("COUNT(DISTINCT v.affected_product)").
        Where("v.create_time >= ? AND v.affected_product <> '' AND v.status = 1", thisWeek).
        Where("NOT EXISTS (?)", subQuery).
        Count(&weeklyAffectedProductCount)
        /*
    db.Select("id, vuln_name, vuln_type, cvss, vuln_level, is_public, status, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
        Where("status = 1").
        Order("create_time DESC").
        Limit(10).
        Find(&vulnDatas)
    for i := range vulnDatas {
        var vulntype types.XqVulnType
        db.Where("id = ?", vulnDatas[i].VulnTypeID).First(&vulntype)
        vulnDatas[i].VulnType = vulntype.Name
    }
        */
    //db.Preload("VulnType").Where("status = 1").Order("create_time DESC").Limit(10).Find(&vulnDatas)
    db.Table("xq_vulnerabilities").
    Select("xq_vulnerabilities.*, xq_vuln_types.name as vuln_type").
    Joins("left join xq_vuln_types on xq_vulnerabilities.vuln_type_id = xq_vuln_types.id").
    Order("create_time DESC").
    Limit(10).
    Find(&vulnDatas)
    return totalCount, pocCount, expCount, affectedProductCount, weeklyCount, weeklyPocCount, weeklyExpCount,weeklyAffectedProductCount, vulnDatas
}

// 分页获取审核通过的漏洞列表
func GetVulnList(page string, pageSize string) (int64, []types.XqVulnerability) {
    var vulnDatas []types.XqVulnerability
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&vulnDatas).Where("status = 1").Count(&totalCount)
    /*
    db.Select("id, vuln_name, vuln_type, vuln_level, cvss, is_public, status, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
    Where("status = 1").Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Order("create_time DESC").
    Omit("user_id, attachment_id, attachment_name, update_time").Find(&vulnDatas)
    */
    db.Table("xq_vulnerabilities").
    Select("xq_vulnerabilities.id, xq_vulnerabilities.vuln_name, xq_vuln_types.name as vuln_type, xq_vulnerabilities.vuln_level, xq_vulnerabilities.cvss, xq_vulnerabilities.is_public, xq_vulnerabilities.status, CASE WHEN xq_vulnerabilities.poc <> '' THEN true ELSE false END AS poc, CASE WHEN xq_vulnerabilities.exp <> '' THEN true ELSE false END AS exp, xq_vulnerabilities.create_time").
    Joins("left join xq_vuln_types on xq_vulnerabilities.vuln_type_id = xq_vuln_types.id").
    Where("xq_vulnerabilities.status = 1").
    Limit(pageSizeNum).
    Offset((pageNum - 1) * pageSizeNum).
    Order("xq_vulnerabilities.create_time DESC").
    Omit("xq_vulnerabilities.user_id, xq_vulnerabilities.attachment_id, xq_vulnerabilities.attachment_name, xq_vulnerabilities.update_time").
    Find(&vulnDatas)
    return totalCount, vulnDatas
}

// 分页获取待审核漏洞列表
func GetUnauditList(page string, pageSize string) (int64, []types.XqVulnerability) {
    var vulnDatas []types.XqVulnerability
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&vulnDatas).Where("status = 0").Count(&totalCount)
    /*
    db.Select("id, vuln_name, vuln_type, vuln_level, cvss, is_public, status, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
    Where("status = 0").Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Order("create_time DESC").
    Omit("user_id, attachment_id, attachment_name, update_time").Find(&vulnDatas)
    */
    db.Table("xq_vulnerabilities").
    Select("xq_vulnerabilities.id, xq_vulnerabilities.vuln_name, xq_vuln_types.name as vuln_type, xq_vulnerabilities.vuln_level, xq_vulnerabilities.cvss, xq_vulnerabilities.is_public, xq_vulnerabilities.status, CASE WHEN xq_vulnerabilities.poc <> '' THEN true ELSE false END AS poc, CASE WHEN xq_vulnerabilities.exp <> '' THEN true ELSE false END AS exp, xq_vulnerabilities.create_time").
    Joins("left join xq_vuln_types on xq_vulnerabilities.vuln_type_id = xq_vuln_types.id").
    Where("xq_vulnerabilities.status = 0").
    Limit(pageSizeNum).
    Offset((pageNum - 1) * pageSizeNum).
    Order("xq_vulnerabilities.create_time DESC").
    Omit("xq_vulnerabilities.user_id, xq_vulnerabilities.attachment_id, xq_vulnerabilities.attachment_name, xq_vulnerabilities.update_time").
    Find(&vulnDatas)
    return totalCount, vulnDatas
}

// 分页获取已审核漏洞列表
func GetAuditedList(page string, pageSize string) (int64, []types.XqVulnerability) {
    var vulnDatas []types.XqVulnerability
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&vulnDatas).Where("status <> 0").Count(&totalCount)
    /*
    db.Select("id, vuln_name, vuln_type, vuln_level, cvss, is_public, status, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
    Where("status <> 0").Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Order("create_time DESC").
    Omit("user_id, attachment_id, attachment_name, update_time").Find(&vulnDatas)
    */
    db.Table("xq_vulnerabilities").
    Select("xq_vulnerabilities.id, xq_vulnerabilities.vuln_name, xq_vuln_types.name as vuln_type, xq_vulnerabilities.vuln_level, xq_vulnerabilities.cvss, xq_vulnerabilities.is_public, xq_vulnerabilities.status, CASE WHEN xq_vulnerabilities.poc <> '' THEN true ELSE false END AS poc, CASE WHEN xq_vulnerabilities.exp <> '' THEN true ELSE false END AS exp, xq_vulnerabilities.create_time").
    Joins("left join xq_vuln_types on xq_vulnerabilities.vuln_type_id = xq_vuln_types.id").
    Where("xq_vulnerabilities.status <> 0").
    Limit(pageSizeNum).
    Offset((pageNum - 1) * pageSizeNum).
    Order("xq_vulnerabilities.create_time DESC").
    Omit("xq_vulnerabilities.user_id, xq_vulnerabilities.attachment_id, xq_vulnerabilities.attachment_name, xq_vulnerabilities.update_time").
    Find(&vulnDatas)
    return totalCount, vulnDatas
}

// 获取漏洞详情，未登录时，返回不包含poc、exp和附件信息
func GetVulnDetail(id string) (types.XqVulnerability) {
    var vulnerabilities types.XqVulnerability
    //var user types.XqUser
    //var vulntype types.XqVulnType
    //db.Where("id = ? AND is_public = true AND status = 1", id).Omit("exp, attachment_id, attachment_name").First(&vulnerabilities)
    //db.Where("id = ?", vulnerabilities.UserID).First(&user)
    //db.Where("id = ?", vulnerabilities.VulnTypeID).First(&vulntype)
    //vulnerabilities.Submitter = user.Username
    //vulnerabilities.VulnType = vulntype.Name
    db.Table("xq_vulnerabilities as v").
        Select("v.*, '' as poc, '' as exp, '' as attachment_id, '' as attachment_name, u.id as user_id, u.username as submitter,vt.id as vuln_type_id, vt.name as vuln_type").
        Joins("left join xq_users as u on v.user_id = u.id").
        Joins("left join xq_vuln_types as vt on v.vuln_type_id = vt.id").
        Where("v.id = ? AND v.is_public = true AND v.status = 1", id).
        First(&vulnerabilities)
    if vulnerabilities.Submitter == "" {
        var sysconfig types.XqSystemConfig
        db.First(&sysconfig)
        vulnerabilities.Submitter = sysconfig.UserDisplay
    }
    return vulnerabilities
}

// 获取漏洞详情，已登录时，返回漏洞全部信息
func GetVulnDetailAuthed(id string, userid uint64, roleid int64) (types.XqVulnerability) {
    var vulnerabilities types.XqVulnerability
    //var user types.XqUser
    //var vulntype types.XqVulnType
    //db.Where("id = ?", id).First(&vulnerabilities)
    //db.Where("id = ?", vulnerabilities.UserID).First(&user)
    //db.Where("id = ?", vulnerabilities.VulnTypeID).First(&vulntype)
    //vulnerabilities.Submitter = user.Username
    //vulnerabilities.VulnType = vulntype.Name
    db.Table("xq_vulnerabilities as v").
        Select("v.*,u.id as user_id, u.username as submitter,vt.id as vuln_type_id, vt.name as vuln_type").
        Joins("left join xq_users as u on v.user_id = u.id").
        Joins("left join xq_vuln_types as vt on v.vuln_type_id = vt.id").
        Where("v.id = ?", id).
        First(&vulnerabilities)
    if vulnerabilities.UserID == userid || roleid == 1 {
        return vulnerabilities
    }
    if vulnerabilities.IsPublic && vulnerabilities.Status == 1 {
        return vulnerabilities
    }
    return types.XqVulnerability{}
}

// 检查漏洞数据
func checkVulnData(vuln types.XqVulnerability) error {
    var vulnerability types.XqVulnerability
    if vuln.VulnName == "" {
        return fmt.Errorf("漏洞名称不能为空")
    }
    if vuln.CVSS < 0 || vuln.CVSS > 10 {
        return fmt.Errorf("CVSS值必须在0到10之间")
    }
    if vuln.VulnLevel == "" {
        return fmt.Errorf("漏洞等级不能为空")
    } else {
        if vuln.VulnLevel != "Critical" && vuln.VulnLevel != "High" && vuln.VulnLevel != "Medium" && vuln.VulnLevel != "Low" {
            return fmt.Errorf("漏洞等级只能为 Critical, High, Medium, Low 中的一个")
        }
    }
    if vuln.Description == "" {
        return fmt.Errorf("漏洞描述不能为空")
    }
    if vuln.AffectedProduct == "" {
        return fmt.Errorf("受影响产品不能为空")
    }
    if vuln.AffectedProductVersion == "" {
        return fmt.Errorf("受影响产品版本不能为空")
    }
    if vuln.RepairSuggestion == "" {
        return fmt.Errorf("修复建议不能为空")
    }
    // 检查CVE格式
    if vuln.CVE != "" {
        if parts := strings.Split(vuln.CVE, "-"); len(parts) != 3 || parts[0] != "CVE" || len(parts[1]) != 4 || len(parts[2]) < 4 {
            return fmt.Errorf("CVE格式不正确")
        }
        result := db.Where("cve = ?", vuln.CVE).First(&vulnerability)
        if result.RowsAffected != 0 {
            return fmt.Errorf("漏洞已存在")
        }
    }
    // 检查NVD格式
    if vuln.NVD != "" {
        if parts := strings.Split(vuln.NVD, "-"); len(parts) != 3 || parts[0] != "NVD" || len(parts[1]) != 4 || len(parts[2]) < 4 {
            return fmt.Errorf("NVD格式不正确")
        }
        result := db.Where("nvd = ?", vuln.NVD).First(&vulnerability)
        if result.RowsAffected != 0 {
            return fmt.Errorf("漏洞已存在")
        }
    }
    // 检查EDBID格式，纯数字，如：50280
    if vuln.EDB != "" {
        if _, err := strconv.Atoi(vuln.EDB); err != nil {
            return fmt.Errorf("EDBID格式不正确")
        }
        result := db.Where("edb = ?", vuln.EDB).First(&vulnerability)
        if result.RowsAffected != 0 {
            return fmt.Errorf("漏洞已存在")
        }
    }
    // 检查CNNVD格式
    if vuln.CNNVD != "" {
        if parts := strings.Split(vuln.CNNVD, "-"); len(parts) != 3 || parts[0] != "CNNVD" || len(parts[1]) != 6 || len(parts[2]) < 4 {
            return fmt.Errorf("CNNVD格式不正确")
        }
        result := db.Where("cnnvd = ?", vuln.CNNVD).First(&vulnerability)
        if result.RowsAffected != 0 {
            return fmt.Errorf("漏洞已存在")
        }
    }
    // 检查CNVD格式
    if vuln.CNVD != "" {
        if parts := strings.Split(vuln.CNVD, "-"); len(parts) != 3 || parts[0] != "CNVD" || len(parts[1]) != 4 || len(parts[2]) < 4 {
            return fmt.Errorf("CNVD格式不正确")
        }
        result := db.Where("cnvd = ?", vuln.CNVD).First(&vulnerability)
        if result.RowsAffected != 0 {
            return fmt.Errorf("漏洞已存在")
        }
    }
    return nil
}

// 生成漏洞ID
func getVdbid() string {
    var hvdid string
    var vulnerabilities types.XqVulnerability
    currentYear := int64(time.Now().Year())
    result := db.Order("id desc").Last(&vulnerabilities)
    if result.RowsAffected == 0 {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, 1)
        return hvdid
    }
    parts := strings.Split(vulnerabilities.ID, "-")
    if len(parts) != 3 {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, 1)
        return hvdid
    }

    year, err := strconv.ParseInt(parts[1], 10, 64)
    if err != nil {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, 1)
        return hvdid
    }
    sequence, err := strconv.ParseInt(parts[2], 10, 64)
    if err != nil {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, 1)
        return hvdid
    }
    if year == currentYear {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, sequence + 1)
    } else {
        hvdid = fmt.Sprintf("HVD-%d-%04d", currentYear, 1)
    }
    return hvdid
}

// 插入漏洞信息
func InsertVuln(vuln types.XqVulnerability) error {
    var attachment types.XqAttachment
    err := checkVulnData(vuln)
    if vuln.AttachmentID != "" {
        db.Where("id = ?", vuln.AttachmentID).First(&attachment)
        vuln.AttachmentName = attachment.Name
    }
    vuln.VulnType = ""
    vuln.Submitter = ""
    hvdid := getVdbid()
    vuln.ID = hvdid
    vuln.Status = 0
    vuln.CreateTime = time.Now()
    err = db.Create(&vuln).Error
    if err != nil {
        return err
    }
    
    return nil
}

// 编辑漏洞
func EditVuln(vuln types.XqVulnerability, userid uint64, roleid int64) error {
    var vulnerability types.XqVulnerability
    res := db.Where("id = ?", vuln.ID).First(&vulnerability)
    if res.RowsAffected == 0 {
        return fmt.Errorf("漏洞不存在")
    }
    if roleid != 1 {
        if vulnerability.UserID != userid {
            return fmt.Errorf("你没有权限编辑该漏洞")
        }
        if vulnerability.Status == 1 {
            return fmt.Errorf("漏洞已审核通过，无法编辑")
        }
    }
    err := checkVulnData(vuln)
    if err != nil {
        return err
    }
    vuln.ReviewComments = vulnerability.ReviewComments
    vuln.UserID = vulnerability.UserID
    vuln.Status = 0
    vuln.UpdateTime = time.Now()
    err = db.Save(&vuln).Error
    return err
}

// 搜索漏洞信息
func SearchVuln(keyword string) []types.XqVulnerability {
    var vulnDatas []types.XqVulnerability
    db.Where("status = 1").Where("id LIKE ? OR cve LIKE ? OR cnnvd LIKE ? OR cnvd LIKE ? OR vuln_name LIKE ? OR description LIKE ? OR affected_product LIKE ?",
        "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&vulnDatas)
    return vulnDatas
}

// 高级搜索
func SearchVulnAdv(data map[string]interface{}) []types.XqVulnerability {
    var vulnDatas []types.XqVulnerability
    var conditions []string
    var values []interface{}
    conditions = append(conditions, "status = 1")
    vulnName, _ := data["vuln_name"].(string)
    if vulnName != "" {
        conditions = append(conditions, "vuln_name LIKE ?")
        values = append(values, "%"+vulnName+"%")
    }
    vulnType, _ := data["vuln_type"].(string)
    if vulnType != "" {
        conditions = append(conditions, "vuln_type LIKE ?")
        values = append(values, "%"+vulnType+"%")
    }
    vulnLevel, _ := data["vuln_level"].(string)
    if vulnLevel != "" {
        conditions = append(conditions, "vuln_level LIKE ?")
        values = append(values, "%"+vulnLevel+"%")
    }
    description, _ := data["description"].(string)
    if description != "" {
        conditions = append(conditions, "description LIKE ?")
        values = append(values, "%"+description+"%")
    }
    affected_product, _ := data["affected_product"].(string)
    if affected_product != "" {
        conditions = append(conditions, "affected_product LIKE ?")
        values = append(values, "%"+affected_product+"%")
    }
    poc, _ := data["poc"].(string)
    if poc != "" {
        conditions = append(conditions, "poc IS NOT NULL AND poc <> ''")
    }
    exp, _ := data["exp"].(string)
    if exp != "" {
        conditions = append(conditions, "poc IS NOT NULL AND poc <> ''")
    }
    submitter, _ := data["submitter"].(string)
    if submitter != "" {
        conditions = append(conditions, "submitter LIKE ?")
        values = append(values, "%"+submitter+"%")
    }
    query := strings.Join(conditions, " AND ")
    db.Where(query, values...).Find(&vulnDatas)
    return vulnDatas
}

// 存储文件到数据库
func StoreFile(file *multipart.FileHeader, userid uint64) (string, error) {
    // 打开文件
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // 读取文件内容
    bytes, err := io.ReadAll(src)
    if err != nil {
        return "", err
    }

    // 生成唯一的文件ID
    attachmentID := utils.GenerateUniqueID()

    // 创建一个新的 Attachment 实例
    attachment := types.XqAttachment{
        ID:          attachmentID,
        UserID:      userid,
        Name:        file.Filename,
        Type:        file.Header.Get("Content-Type"),
        Data:        bytes,
        Status:      1,
        CreateTime:  time.Now(),
        UpdateTime:  time.Now(),
    }

    // 将文件保存到数据库
    if err := db.Create(&attachment).Error; err != nil {
        return "", err
    }

    return attachmentID, nil
}

// 获取文件内容
func GetFileContent(attachmentID string) (types.XqAttachment, error) {
    var attachment types.XqAttachment
    err := db.Where("id = ?", attachmentID).First(&attachment).Error
    return attachment, err
}

// 删除文件
func DeleteFile(attachmentID string, userid uint64) error {
    return db.Where("id = ? AND user_id = ?", attachmentID, userid).Delete(&types.XqAttachment{}).Error
}

// 获取漏洞类型列表
func GetVulnTypeList() []types.XqVulnType {
    var vulnTypes []types.XqVulnType
    db.Find(&vulnTypes)
    return vulnTypes
}

// 分页获取漏洞类型列表
func GetVulnType(page string, pageSize string) (int64, []types.XqVulnType) {
    var vulntype []types.XqVulnType
    var totalCount int64
    pageNum, _ := strconv.Atoi(page)
    pageSizeNum, _ := strconv.Atoi(pageSize)
    db.Model(&vulntype).Count(&totalCount)
    db.Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Find(&vulntype)
    return totalCount, vulntype
}

// 添加漏洞类型
func AddVulnType(name string) error {
    var vulnType types.XqVulnType
    res := db.Where("name = ?", name).First(&vulnType).RowsAffected
    if res != 0 {
        return fmt.Errorf("漏洞类型已存在")
    }
    vulnType.Name = name
    vulnType.CreateTime = time.Now()
    return db.Create(&vulnType).Error
}

// 更新漏洞类型
func UpdateVlunType(id uint64, name string) error {
    var vulnType types.XqVulnType
    res := db.Where("id = ?", id).First(&vulnType).RowsAffected
    if res == 0 {
        return fmt.Errorf("漏洞类型不存在")
    }
    vulnType.Name = name
    vulnType.UpdateTime = time.Now()
    return db.Save(&vulnType).Error
}
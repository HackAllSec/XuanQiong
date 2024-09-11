package models

import (
    "fmt"
    "strings"
    "strconv"
    "time"
    "xuanqiong/types"
)

// 获取漏洞信息
func GetVulnAbstract() ([]types.Vulnerability, int64, int64, int64) {
    var vulnDatas []types.Vulnerability
    var totalCount int64
    var pocCount int64
    var expCount int64
    db.Model(&vulnDatas).Count(&totalCount)
    db.Raw("SELECT COUNT(*) FROM vulnerabilities WHERE poc <> ''").Scan(&pocCount)
    db.Raw("SELECT COUNT(*) FROM vulnerabilities WHERE exp <> ''").Scan(&expCount)
    db.Select("id, vuln_name, vuln_type, vuln_level, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").
    Order("create_time DESC").
    Limit(20).Find(&vulnDatas)
    return vulnDatas, totalCount, pocCount, expCount
}

// 获取漏洞详情，未登录时，返回不包含poc和exp
func GetVulnDetail(id string) (types.Vulnerability) {
    var vulnerabilities types.Vulnerability
    db.Where("id = ?", id).Omit("poc, exp").First(&vulnerabilities)
    return vulnerabilities
}

// 获取漏洞详情，已登录时，返回漏洞全部信息
func GetVulnDetailAuthed(id string) (types.Vulnerability) {
    var vulnerabilities types.Vulnerability
    db.Where("id = ?", id).First(&vulnerabilities)
    return vulnerabilities
}

// 检查漏洞数据
func checkVulnData(vuln types.Vulnerability) error {
    if vuln.VulnName == "" {
        return fmt.Errorf("漏洞名称不能为空")
    }
    if vuln.VulnType == "" {
        return fmt.Errorf("漏洞类型不能为空")
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
    if vuln.RepairSuggestion == "" {
        return fmt.Errorf("修复建议不能为空")
    }
    // 检查CVE格式
    if vuln.CVE != "" {
        if parts := strings.Split(vuln.CVE, "-"); len(parts) != 3 || parts[0] != "CVE" || len(parts[1]) != 4 || len(parts[2]) < 4 {
            return fmt.Errorf("CVE格式不正确")
        }
    }
    // 检查CNNVD格式
    if vuln.CNNVD != "" {
        if parts := strings.Split(vuln.CNNVD, "-"); len(parts) != 3 || parts[0] != "CNNVD" || len(parts[1]) != 6 || len(parts[2]) < 4 {
            return fmt.Errorf("CNNVD格式不正确")
        }
    }
    // 检查CNVD格式
    if vuln.CNVD != "" {
        if parts := strings.Split(vuln.CNVD, "-"); len(parts) != 3 || parts[0] != "CNVD" || len(parts[1]) != 4 || len(parts[2]) < 4 {
            return fmt.Errorf("CNVD格式不正确")
        }
    }
    return nil
}

// 生成VDBID
func getVdbid() string {
    var hvbid string
    var vulnerabilities types.Vulnerability
    currentYear := int64(time.Now().Year())
    result := db.Order("id desc").Last(&vulnerabilities)
    if result.RowsAffected == 0 {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, 1)
        return hvbid
    }
    parts := strings.Split(vulnerabilities.ID, "-")
    if len(parts) != 3 {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, 1)
        return hvbid
    }

    year, err := strconv.ParseInt(parts[1], 10, 64)
    if err != nil {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, 1)
        return hvbid
    }
    sequence, err := strconv.ParseInt(parts[2], 10, 64)
    if err != nil {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, 1)
        return hvbid
    }
    if year == currentYear {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, sequence + 1)
    } else {
        hvbid = fmt.Sprintf("HVB-%d-%04d", currentYear, 1)
    }
    return hvbid
}

// 插入漏洞信息
func InsertVuln(vuln types.Vulnerability) (*types.Vulnerability, error) {
    err := checkVulnData(vuln)
    if err != nil {
        return nil, err
    }
    hvbid := getVdbid()
    vuln.ID = hvbid
    vuln.CreateTime = time.Now()
    err = db.Create(&vuln).Error
    if err != nil {
        return nil, err
    }
    return &vuln, nil
}

// 搜索漏洞信息
func SearchVuln(keyword string) []types.Vulnerability {
    var vulnDatas []types.Vulnerability
    db.Where("id LIKE ? OR cve LIKE ? OR cnnvd LIKE ? OR cnvd LIKE ? OR vuln_name LIKE ? OR description LIKE ? OR affected_product LIKE ?",
        "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&vulnDatas)
    return vulnDatas
}

// 高级搜索
func SearchVulnAdv(data map[string]interface{}) []types.Vulnerability {
    var vulnDatas []types.Vulnerability
    var conditions []string
    var values []interface{}
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
    submit, _ := data["submit"].(string)
    if submit != "" {
        conditions = append(conditions, "submit LIKE ?")
        values = append(values, "%"+submit+"%")
    }
    query := strings.Join(conditions, " AND ")
    db.Where(query, values...).Find(&vulnDatas)
    return vulnDatas
}

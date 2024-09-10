package models

import (
    "xuanqiong/types"
)

// 获取漏洞摘要
func GetVulnAbstract() ([]types.Vulnerability) {
    var vulnerabilities []types.Vulnerability
    db.Select("id, vuln_name, vuln_type, vuln_level, CASE WHEN poc <> '' THEN true ELSE false END AS poc, CASE WHEN exp <> '' THEN true ELSE false END AS exp, create_time").Find(&vulnerabilities)
    return vulnerabilities
}

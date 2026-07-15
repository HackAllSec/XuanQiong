package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"xuanqiong/backend/types"
)

var vulnCSVHeaders = []string{
	"vuln_name", "vuln_type_id", "vuln_level", "cvss", "description",
	"affected_product", "affected_product_version", "repair_suggestion",
	"cve", "nvd", "edb", "cnnvd", "cnvd", "fofa_query", "zoomeye_query",
	"quake_query", "hunter_query", "google_query", "shodan_query", "censys_query",
	"greynoise_query", "poc", "poc_type", "exp", "exp_type", "is_public",
}

func ExportVulnsCSV(writer io.Writer) error {
	csvWriter := csv.NewWriter(writer)
	if err := csvWriter.Write(vulnCSVHeaders); err != nil {
		return err
	}

	var vulns []types.XqVulnerability
	if err := db.Order("create_time desc").Find(&vulns).Error; err != nil {
		return err
	}
	for _, vuln := range vulns {
		row := []string{
			vuln.VulnName,
			strconv.FormatUint(vuln.VulnTypeID, 10),
			vuln.VulnLevel,
			strconv.FormatFloat(vuln.CVSS, 'f', -1, 64),
			vuln.Description,
			vuln.AffectedProduct,
			vuln.AffectedProductVersion,
			vuln.RepairSuggestion,
			vuln.CVE,
			vuln.NVD,
			vuln.EDB,
			vuln.CNNVD,
			vuln.CNVD,
			vuln.FofaQuery,
			vuln.ZoomEyeQuery,
			vuln.QuakeQuery,
			vuln.HunterQuery,
			vuln.GoogleQuery,
			vuln.ShodanQuery,
			vuln.CensysQuery,
			vuln.GreynoiseQuery,
			vuln.Poc,
			vuln.PocType,
			vuln.Exp,
			vuln.ExpType,
			strconv.FormatBool(vuln.IsPublic),
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return csvWriter.Error()
}

func ImportVulnsCSV(reader io.Reader, userID uint64) (int, []string) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	rows, err := csvReader.ReadAll()
	if err != nil {
		return 0, []string{err.Error()}
	}
	if len(rows) < 2 {
		return 0, []string{"no vulnerability rows found"}
	}

	headerMap := map[string]int{}
	for index, header := range rows[0] {
		headerMap[strings.TrimSpace(header)] = index
	}

	imported := 0
	errors := []string{}
	for rowIndex, row := range rows[1:] {
		vuln, err := vulnerabilityFromCSVRow(headerMap, row)
		if err != nil {
			errors = append(errors, fmt.Sprintf("row %d: %s", rowIndex+2, err.Error()))
			continue
		}
		vuln.UserID = userID
		if err := InsertVuln(vuln); err != nil {
			errors = append(errors, fmt.Sprintf("row %d: %s", rowIndex+2, err.Error()))
			continue
		}
		imported++
	}
	return imported, errors
}

func vulnerabilityFromCSVRow(headerMap map[string]int, row []string) (types.XqVulnerability, error) {
	get := func(name string) string {
		index, ok := headerMap[name]
		if !ok || index >= len(row) {
			return ""
		}
		return strings.TrimSpace(row[index])
	}
	vulnTypeID, _ := strconv.ParseUint(get("vuln_type_id"), 10, 64)
	cvss, _ := strconv.ParseFloat(get("cvss"), 64)
	isPublic, _ := strconv.ParseBool(get("is_public"))
	return types.XqVulnerability{
		VulnName:               get("vuln_name"),
		VulnTypeID:             vulnTypeID,
		VulnLevel:              get("vuln_level"),
		CVSS:                   cvss,
		Description:            get("description"),
		AffectedProduct:        get("affected_product"),
		AffectedProductVersion: get("affected_product_version"),
		RepairSuggestion:       get("repair_suggestion"),
		CVE:                    get("cve"),
		NVD:                    get("nvd"),
		EDB:                    get("edb"),
		CNNVD:                  get("cnnvd"),
		CNVD:                   get("cnvd"),
		FofaQuery:              get("fofa_query"),
		ZoomEyeQuery:           get("zoomeye_query"),
		QuakeQuery:             get("quake_query"),
		HunterQuery:            get("hunter_query"),
		GoogleQuery:            get("google_query"),
		ShodanQuery:            get("shodan_query"),
		CensysQuery:            get("censys_query"),
		GreynoiseQuery:         get("greynoise_query"),
		Poc:                    get("poc"),
		PocType:                get("poc_type"),
		Exp:                    get("exp"),
		ExpType:                get("exp_type"),
		IsPublic:               isPublic,
	}, nil
}

type SystemBackup struct {
	Version         int                      `json:"version"`
	CreatedAt       time.Time                `json:"created_at"`
	SystemConfigs   []types.XqSystemConfig   `json:"system_configs"`
	JwtConfigs      []types.XqJwtConfig      `json:"jwt_configs"`
	EmailConfigs    []types.XqEmailConfig    `json:"email_configs"`
	NoticeConfigs   []types.XqNoticeConfig   `json:"notice_configs"`
	Users           []types.XqUser           `json:"users"`
	Roles           []types.XqRole           `json:"roles"`
	Permissions     []types.XqPermission     `json:"permissions"`
	RolePermissions []types.XqRolePermission `json:"role_permissions"`
	UserRoles       []types.XqUserRole       `json:"user_roles"`
	VulnTypes       []types.XqVulnType       `json:"vuln_types"`
	Vulnerabilities []types.XqVulnerability  `json:"vulnerabilities"`
	Attachments     []types.XqAttachment     `json:"attachments"`
	ScoreRules      []types.XqScoreRule      `json:"score_rules"`
	RankingDetails  []types.XqRankingDetail  `json:"ranking_details"`
	Messages        []types.XqMessage        `json:"messages"`
	AuditLogs       []types.XqAuditLog       `json:"audit_logs"`
	APIKeys         []BackupAPIKey           `json:"api_keys"`
}

type BackupAPIKey struct {
	ID         uint64     `json:"id"`
	UserID     uint64     `json:"user_id"`
	Name       string     `json:"name"`
	KeyPrefix  string     `json:"key_prefix"`
	KeyHash    string     `json:"key_hash"`
	Scopes     string     `json:"scopes"`
	Status     int64      `json:"status"`
	LastUsedAt *time.Time `json:"last_used_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
}

func CreateSystemBackup(writer io.Writer) error {
	backup := SystemBackup{Version: 1, CreatedAt: time.Now()}
	db.Find(&backup.SystemConfigs)
	db.Find(&backup.JwtConfigs)
	db.Find(&backup.EmailConfigs)
	db.Find(&backup.NoticeConfigs)
	db.Find(&backup.Users)
	db.Find(&backup.Roles)
	db.Find(&backup.Permissions)
	db.Find(&backup.RolePermissions)
	db.Find(&backup.UserRoles)
	db.Find(&backup.VulnTypes)
	db.Find(&backup.Vulnerabilities)
	db.Find(&backup.Attachments)
	db.Find(&backup.ScoreRules)
	db.Find(&backup.RankingDetails)
	db.Find(&backup.Messages)
	db.Find(&backup.AuditLogs)
	var apiKeys []types.XqAPIKey
	db.Find(&apiKeys)
	backup.APIKeys = make([]BackupAPIKey, 0, len(apiKeys))
	for _, apiKey := range apiKeys {
		backup.APIKeys = append(backup.APIKeys, BackupAPIKey{
			ID:         apiKey.ID,
			UserID:     apiKey.UserID,
			Name:       apiKey.Name,
			KeyPrefix:  apiKey.KeyPrefix,
			KeyHash:    apiKey.KeyHash,
			Scopes:     apiKey.Scopes,
			Status:     apiKey.Status,
			LastUsedAt: apiKey.LastUsedAt,
			ExpiresAt:  apiKey.ExpiresAt,
			CreateTime: apiKey.CreateTime,
			UpdateTime: apiKey.UpdateTime,
		})
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(backup)
}

func RestoreSystemBackup(reader io.Reader) error {
	var backup SystemBackup
	if err := json.NewDecoder(reader).Decode(&backup); err != nil {
		return err
	}
	if backup.Version != 1 {
		return fmt.Errorf("unsupported backup version")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		tables := []interface{}{
			&types.XqAuditLog{},
			&types.XqMessage{},
			&types.XqAPIKey{},
			&types.XqRankingDetail{},
			&types.XqScoreRule{},
			&types.XqAttachment{},
			&types.XqVulnerability{},
			&types.XqVulnType{},
			&types.XqUserRole{},
			&types.XqRolePermission{},
			&types.XqPermission{},
			&types.XqRole{},
			&types.XqUser{},
			&types.XqNoticeConfig{},
			&types.XqEmailConfig{},
			&types.XqJwtConfig{},
			&types.XqSystemConfig{},
		}
		for _, table := range tables {
			if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error; err != nil {
				return err
			}
		}
		createAll := []interface{}{
			backup.SystemConfigs,
			backup.JwtConfigs,
			backup.EmailConfigs,
			backup.NoticeConfigs,
			backup.Users,
			backup.Roles,
			backup.Permissions,
			backup.RolePermissions,
			backup.UserRoles,
			backup.VulnTypes,
			backup.Vulnerabilities,
			backup.Attachments,
			backup.ScoreRules,
			backup.RankingDetails,
			backup.Messages,
			backup.AuditLogs,
			apiKeyBackupRecords(backup.APIKeys),
		}
		for _, records := range createAll {
			if reflect.ValueOf(records).Len() == 0 {
				continue
			}
			if err := tx.Create(records).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func apiKeyBackupRecords(apiKeys []BackupAPIKey) []types.XqAPIKey {
	records := make([]types.XqAPIKey, 0, len(apiKeys))
	for _, apiKey := range apiKeys {
		records = append(records, types.XqAPIKey{
			ID:         apiKey.ID,
			UserID:     apiKey.UserID,
			Name:       apiKey.Name,
			KeyPrefix:  apiKey.KeyPrefix,
			KeyHash:    apiKey.KeyHash,
			Scopes:     apiKey.Scopes,
			Status:     apiKey.Status,
			LastUsedAt: apiKey.LastUsedAt,
			ExpiresAt:  apiKey.ExpiresAt,
			CreateTime: apiKey.CreateTime,
			UpdateTime: apiKey.UpdateTime,
		})
	}
	return records
}

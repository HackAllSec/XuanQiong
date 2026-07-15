package models

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"xuanqiong/backend/types"
)

const auditPayloadLimit = 2048
const auditReadLimit = 4096

func maskSensitiveValue(key string, value string) string {
	lowerKey := strings.ToLower(key)
	sensitiveKeywords := []string{"password", "token", "secret", "authorization", "jwt", "captcha", "cookie", "key"}
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lowerKey, keyword) {
			return "***"
		}
	}
	return value
}

func truncateAuditText(value string) string {
	if len(value) <= auditPayloadLimit {
		return value
	}
	return value[:auditPayloadLimit] + "...(truncated)"
}

func sanitizeJSONBody(raw []byte) string {
	if len(raw) == 0 {
		return ""
	}
	var parsed interface{}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return truncateAuditText(string(raw))
	}
	sanitized := sanitizeJSONValue(parsed)
	data, err := json.Marshal(sanitized)
	if err != nil {
		return truncateAuditText(string(raw))
	}
	return truncateAuditText(string(data))
}

func sanitizeJSONValue(value interface{}) interface{} {
	switch current := value.(type) {
	case map[string]interface{}:
		result := map[string]interface{}{}
		for key, item := range current {
			if text, ok := item.(string); ok {
				result[key] = maskSensitiveValue(key, text)
				continue
			}
			result[key] = sanitizeJSONValue(item)
		}
		return result
	case []interface{}:
		result := make([]interface{}, 0, len(current))
		for _, item := range current {
			result = append(result, sanitizeJSONValue(item))
		}
		return result
	default:
		return value
	}
}

func sanitizeFormValues(values url.Values) string {
	if len(values) == 0 {
		return ""
	}
	result := map[string][]string{}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		valueList := values[key]
		maskedValues := make([]string, 0, len(valueList))
		for _, value := range valueList {
			maskedValues = append(maskedValues, maskSensitiveValue(key, value))
		}
		result[key] = maskedValues
	}
	data, _ := json.Marshal(result)
	return truncateAuditText(string(data))
}

func sanitizeMultipartBody(form *multipart.Form) string {
	if form == nil {
		return ""
	}
	payload := map[string]interface{}{}
	if len(form.Value) > 0 {
		payload["fields"] = sanitizeFormValues(form.Value)
	}
	if len(form.File) > 0 {
		files := map[string][]string{}
		for key, headers := range form.File {
			names := make([]string, 0, len(headers))
			for _, header := range headers {
				names = append(names, header.Filename)
			}
			files[key] = names
		}
		payload["files"] = files
	}
	data, _ := json.Marshal(payload)
	return truncateAuditText(string(data))
}

func SanitizeRequestBody(request *http.Request, rawBody []byte) string {
	contentType := request.Header.Get("Content-Type")
	switch {
	case strings.Contains(contentType, "application/json"):
		return sanitizeJSONBody(rawBody)
	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		values, err := url.ParseQuery(string(rawBody))
		if err != nil {
			return truncateAuditText(string(rawBody))
		}
		return sanitizeFormValues(values)
	case strings.Contains(contentType, "multipart/form-data"):
		if err := request.ParseMultipartForm(10 << 20); err == nil {
			return sanitizeMultipartBody(request.MultipartForm)
		}
		return ""
	default:
		return truncateAuditText(string(rawBody))
	}
}

func SanitizeResponseBody(rawBody []byte) string {
	return sanitizeJSONBody(rawBody)
}

func ShouldCaptureRequestBody(request *http.Request) bool {
	contentType := request.Header.Get("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		return false
	}
	if request.ContentLength < 0 || request.ContentLength > auditReadLimit {
		return false
	}
	return strings.Contains(contentType, "application/json") ||
		strings.Contains(contentType, "application/x-www-form-urlencoded")
}

func CaptureRequestBody(request *http.Request) ([]byte, error) {
	if request.Body == nil {
		return nil, nil
	}
	rawBody, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	request.Body = io.NopCloser(strings.NewReader(string(rawBody)))
	return rawBody, nil
}

func CreateAuditLog(logEntry types.XqAuditLog) error {
	logEntry.RequestBody = truncateAuditText(logEntry.RequestBody)
	logEntry.ResponseBody = truncateAuditText(logEntry.ResponseBody)
	if logEntry.CreateTime.IsZero() {
		logEntry.CreateTime = time.Now()
	}
	return db.Create(&logEntry).Error
}

func GetAuditLogs(page string, pageSize string, keyword string, action string) (int64, []types.XqAuditLog) {
	var logs []types.XqAuditLog
	var totalCount int64

	pageNum := 1
	pageSizeNum := 20
	if page != "" {
		if parsed, err := strconv.Atoi(page); err == nil && parsed > 0 {
			pageNum = parsed
		}
	}
	if pageSize != "" {
		if parsed, err := strconv.Atoi(pageSize); err == nil && parsed > 0 {
			pageSizeNum = parsed
		}
	}

	query := db.Model(&types.XqAuditLog{})
	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("username LIKE ? OR path LIKE ? OR client_ip LIKE ?", like, like, like)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}

	query.Count(&totalCount)
	query.Order("id desc").Limit(pageSizeNum).Offset((pageNum - 1) * pageSizeNum).Find(&logs)
	return totalCount, logs
}

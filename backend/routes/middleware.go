package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"xuanqiong/backend/models"
	"xuanqiong/backend/types"
)

const currentUserContextKey = "current_user"
const auditSkipResponseBodyKey = "audit_skip_response_body"
const apiKeyScopesContextKey = "api_key_scopes"

type auditResponseWriter struct {
	gin.ResponseWriter
	body    *bytes.Buffer
	context *gin.Context
}

func (writer auditResponseWriter) Write(data []byte) (int, error) {
	if skipResponseBody, _ := writer.context.Get(auditSkipResponseBodyKey); skipResponseBody != true {
		writer.body.Write(data)
	}
	return writer.ResponseWriter.Write(data)
}

func normalizeAccessTokenHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		xAuthToken := strings.TrimSpace(c.GetHeader("X-Auth-Token"))
		if authHeader != "" && !strings.HasPrefix(strings.ToLower(authHeader), "basic ") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": "Authorization header is not allowed. Use X-Auth-Token."})
			return
		}
		if xAuthToken != "" {
			c.Request.Header.Set("Authorization", "Bearer "+xAuthToken)
		}
		c.Next()
	}
}

func currentUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			if currentUser := models.GetUserByToken(token); currentUser != nil {
				c.Set(currentUserContextKey, currentUser)
				c.Next()
				return
			}
		}
		if currentUser, scopes := models.GetUserByAPIKeyWithScopes(c.GetHeader("X-API-Key")); currentUser != nil {
			c.Set(currentUserContextKey, currentUser)
			c.Set(apiKeyScopesContextKey, scopes)
		}
		c.Next()
	}
}

func currentUserFromContext(c *gin.Context) *types.XqUser {
	value, exists := c.Get(currentUserContextKey)
	if !exists {
		return nil
	}
	currentUser, ok := value.(*types.XqUser)
	if !ok {
		return nil
	}
	return currentUser
}

func requireAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if currentUserFromContext(c) == nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": "Permission denied"})
			return
		}
		c.Next()
	}
}

func currentRequestAPIKeyScopes(c *gin.Context) ([]string, bool) {
	value, exists := c.Get(apiKeyScopesContextKey)
	if !exists {
		return nil, false
	}
	scopes, ok := value.([]string)
	return scopes, ok
}

func apiKeyHasAnyPermissionScope(c *gin.Context, permissionCodes ...string) bool {
	scopes, exists := currentRequestAPIKeyScopes(c)
	if !exists {
		return true
	}
	if len(permissionCodes) == 0 {
		return false
	}
	scopeSet := map[string]bool{}
	for _, scope := range scopes {
		scopeSet[scope] = true
	}
	for _, code := range permissionCodes {
		if scopeSet[code] {
			return true
		}
	}
	return false
}

func requireBrowserSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := currentRequestAPIKeyScopes(c); exists {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": "API key is not allowed for this operation"})
			return
		}
		c.Next()
	}
}

func requirePermissionMiddleware(permissionCodes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := currentUserFromContext(c)
		if currentUser == nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": "Permission denied"})
			return
		}
		if len(permissionCodes) == 0 || models.UserHasAnyPermission(currentUser.ID, permissionCodes...) {
			if apiKeyHasAnyPermissionScope(c, permissionCodes...) {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": "Permission denied"})
	}
}

func markAuditActionMiddleware(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("audit_action", action)
		c.Next()
	}
}

func auditLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rawBody []byte
		if shouldCaptureAudit(c.Request.Method) && models.ShouldCaptureRequestBody(c.Request) {
			if capturedBody, err := models.CaptureRequestBody(c.Request); err == nil {
				rawBody = capturedBody
			}
		}

		recorder := &auditResponseWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil), context: c}
		c.Writer = recorder
		startedAt := time.Now()
		c.Next()

		actionValue, shouldAudit := c.Get("audit_action")
		if !shouldAudit {
			return
		}

		action, _ := actionValue.(string)
		currentUser := currentUserFromContext(c)
		skipResponseBody, _ := c.Get(auditSkipResponseBodyKey)
		responseBody := ""
		if skipResponseBody != true {
			responseBody = models.SanitizeResponseBody(recorder.body.Bytes())
		}
		auditLog := types.XqAuditLog{
			Action:       action,
			Method:       c.Request.Method,
			Path:         c.FullPath(),
			StatusCode:   c.Writer.Status(),
			ClientIP:     c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			RequestBody:  models.SanitizeRequestBody(c.Request, rawBody),
			ResponseBody: responseBody,
			CreateTime:   startedAt,
		}
		if currentUser != nil {
			auditLog.UserID = currentUser.ID
			auditLog.Username = currentUser.Username
		} else if auditLog.RequestBody != "" {
			var payload map[string]interface{}
			if err := json.Unmarshal(rawBody, &payload); err == nil {
				if username, ok := payload["username"].(string); ok {
					auditLog.Username = username
				}
			}
		}

		if skipResponseBody != true {
			var responsePayload map[string]interface{}
			if err := json.Unmarshal(recorder.body.Bytes(), &responsePayload); err == nil {
				if code, ok := responsePayload["code"].(float64); ok {
					auditLog.ResultCode = int64(code)
				}
				if message, ok := responsePayload["msg"].(string); ok {
					auditLog.ResultMessage = message
				}
			}
		}
		_ = models.CreateAuditLog(auditLog)
	}
}

func shouldCaptureAudit(method string) bool {
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}

func protectedRoute(permissionCodes []string, auditAction string, handler gin.HandlerFunc) []gin.HandlerFunc {
	handlers := []gin.HandlerFunc{
		markAuditActionMiddleware(auditAction),
		requireAuthMiddleware(),
	}
	if len(permissionCodes) > 0 {
		handlers = append(handlers, requirePermissionMiddleware(permissionCodes...))
	}
	handlers = append(handlers, handler)
	return handlers
}

func protectedSessionRoute(permissionCodes []string, auditAction string, handler gin.HandlerFunc) []gin.HandlerFunc {
	handlers := protectedRoute(permissionCodes, auditAction, handler)
	return append(handlers[:2], append([]gin.HandlerFunc{requireBrowserSessionMiddleware()}, handlers[2:]...)...)
}

func auditedPublicRoute(auditAction string, handler gin.HandlerFunc) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		markAuditActionMiddleware(auditAction),
		handler,
	}
}

package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNormalizeAccessTokenHeaderMiddleware_BitsUT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name              string
		authorization     string
		xAuthToken        string
		wantStatus        int
		wantAuthorization string
		wantHandlerCalled bool
	}{
		{
			name:              "reject external bearer authorization",
			authorization:     "Bearer raw-token",
			wantStatus:        http.StatusOK,
			wantHandlerCalled: false,
		},
		{
			name:              "reject external api key authorization",
			authorization:     "ApiKey xqk_raw-token",
			wantStatus:        http.StatusOK,
			wantHandlerCalled: false,
		},
		{
			name:              "allow nginx basic authorization without app auth",
			authorization:     "Basic dXNlcjpwYXNz",
			wantStatus:        http.StatusOK,
			wantAuthorization: "Basic dXNlcjpwYXNz",
			wantHandlerCalled: true,
		},
		{
			name:              "normalize x auth token to internal bearer",
			xAuthToken:        "app-token",
			wantStatus:        http.StatusOK,
			wantAuthorization: "Bearer app-token",
			wantHandlerCalled: true,
		},
		{
			name:              "prefer x auth token when basic auth is also present",
			authorization:     "Basic dXNlcjpwYXNz",
			xAuthToken:        "app-token",
			wantStatus:        http.StatusOK,
			wantAuthorization: "Bearer app-token",
			wantHandlerCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlerCalled := false
			router := gin.New()
			router.Use(normalizeAccessTokenHeaderMiddleware())
			router.GET("/probe", func(c *gin.Context) {
				handlerCalled = true
				if got := c.Request.Header.Get("Authorization"); got != tt.wantAuthorization {
					t.Fatalf("Authorization header = %q, want %q", got, tt.wantAuthorization)
				}
				c.JSON(http.StatusOK, gin.H{"code": 1})
			})

			request := httptest.NewRequest(http.MethodGet, "/probe", nil)
			if tt.authorization != "" {
				request.Header.Set("Authorization", tt.authorization)
			}
			if tt.xAuthToken != "" {
				request.Header.Set("X-Auth-Token", tt.xAuthToken)
			}
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body=%s", response.Code, tt.wantStatus, response.Body.String())
			}
			if handlerCalled != tt.wantHandlerCalled {
				t.Fatalf("handlerCalled = %v, want %v", handlerCalled, tt.wantHandlerCalled)
			}
		})
	}
}

func TestAuditResponseWriterSkipsConfiguredResponseBody_BitsUT(t *testing.T) {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Set(auditSkipResponseBodyKey, true)
	buffer := bytes.NewBuffer(nil)
	writer := auditResponseWriter{ResponseWriter: context.Writer, body: buffer, context: context}

	if _, err := writer.Write([]byte("sensitive export body")); err != nil {
		t.Fatalf("write response: %v", err)
	}
	if buffer.Len() != 0 {
		t.Fatalf("audit buffer should stay empty when response body capture is skipped, got %q", buffer.String())
	}
	if response.Body.String() != "sensitive export body" {
		t.Fatalf("client response body was not written")
	}
}

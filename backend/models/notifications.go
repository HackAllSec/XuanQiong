package models

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type EventNotice struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func SendEventNotice(event EventNotice) error {
	_, _, _, noticeConf := GetSystemConfig()
	webhook := strings.TrimSpace(noticeConf.Webhook)
	if noticeConf.Type == 0 || webhook == "" {
		return nil
	}

	switch noticeConf.Type {
	case 1:
		return sendDingTalkNotice(webhook, noticeConf.Secret, event)
	case 2:
		return sendWxWorkNotice(webhook, event)
	default:
		return nil
	}
}

func sendDingTalkNotice(webhook string, secret string, event EventNotice) error {
	if secret != "" {
		timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
		signPayload := timestamp + "\n" + secret
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write([]byte(signPayload))
		sign := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
		separator := "?"
		if strings.Contains(webhook, "?") {
			separator = "&"
		}
		webhook = webhook + separator + "timestamp=" + timestamp + "&sign=" + sign
	}
	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": event.Title,
			"text":  "### " + event.Title + "\n\n" + event.Content,
		},
	}
	return postWebhookJSON(webhook, payload)
}

func sendWxWorkNotice(webhook string, event EventNotice) error {
	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"content": "### " + event.Title + "\n" + event.Content,
		},
	}
	return postWebhookJSON(webhook, payload)
}

func postWebhookJSON(webhook string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	client := http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}
	return nil
}

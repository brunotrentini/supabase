package mailer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/supabase/auth/internal/models"
)

func (m *TemplateMailer) SendAuthWebhook(eventType string, user *models.User, link string) {
	webhookURL := os.Getenv("AUTH_WEBHOOK_URL")

	if webhookURL == "" {
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		payload := map[string]interface{}{
			"event":         eventType,
			"user_id":       user.ID,
			"email":         user.Email,
			"user_metadata": user.UserMetaData,
			"link":          link,
			"timestamp":     time.Now().UTC(),
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			logrus.Errorf("Error marshaling webhook payload: %v", err)
			return
		}

		req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(jsonData))
		if err != nil {
			logrus.Errorf("Error creating webhook request: %v", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logrus.Errorf("Error sending auth webhook: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			logrus.Warnf("Webhook returned error status: %d", resp.StatusCode)
		}
	}()
}

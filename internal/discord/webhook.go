package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Webhook struct {
	Content string `json:"content"`
}

func SendWebhook(url string, change int, id int) {
	var webhookData Webhook
	webhookData.Content = fmt.Sprintf("New updated spotted for App %d. \nChange Number: `%d`", id, change)
	webhookJson, err := json.Marshal(webhookData)
	r, err := http.Post(url, "application/json", bytes.NewBuffer(webhookJson))
	if err != nil {
		log.Error().
			Err(err).
			Msg("There was an error sending the POST request to discord. Skipping webhook execution.")
		return
	}
	log.Debug().Str("status", r.Status).Msg(fmt.Sprintf("Webhook execution for App %d finished.", id))
}

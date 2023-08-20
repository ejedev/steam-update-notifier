package steamcmd

import (
	"fmt"

	"github.com/ejedev/steam-update-notifier/internal/discord"
	"github.com/rs/zerolog/log"
)

func CheckChanges(apps map[int]int, webhook string, enabled bool) map[int]int {
	for key := range apps {
		newChangeNumber := GetChangeNumber(key)
		if apps[key] != newChangeNumber {
			apps[key] = newChangeNumber
			log.Info().
				Int("old", apps[key]).
				Int("new", newChangeNumber).
				Msg(fmt.Sprintf("New update spotted for App %d", key))
			if enabled {
				discord.SendWebhook(webhook, newChangeNumber, key)
			}
		} else {
			log.Debug().Msg(fmt.Sprintf("No update for App %d.", key))
		}
	}

	return apps
}

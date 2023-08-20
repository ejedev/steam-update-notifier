package steamcmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func InitDataTracking(apps []int) map[int]int {
	history := make(map[int]int)
	for _, app := range apps {
		history[app] = GetChangeNumber(app)
		log.Debug().Msg(fmt.Sprintf("App %d initialized with change number %d.", app, history[app]))
	}
	return history
}

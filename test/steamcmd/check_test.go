package steamcmd

import (
	"testing"

	"github.com/ejedev/steam-update-notifier/internal/steamcmd"
	"github.com/rs/zerolog"
)

func TestCheckChanges(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.NoLevel)
	apps := map[int]int{570: 0}
	apps[570] = 0
	result := steamcmd.CheckChanges(apps, "", false)
	if result[570] == 0 {
		t.Errorf("CheckChanges(map[int]int{570: 0}, \"\", false) == %d; want != 0", result[570])
	}
}

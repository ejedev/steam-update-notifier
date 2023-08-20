package steamcmd

import (
	"testing"

	"github.com/ejedev/steam-update-notifier/internal/steamcmd"
)

func TestInitDataTracking(t *testing.T) {
	apps := []int{570}
	result := steamcmd.InitDataTracking(apps)
	if result[570] <= 0 {
		t.Errorf("InitDataTracking([]int{570}) = %d; want >0", result[570])
	}
}

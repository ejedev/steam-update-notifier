package steamcmd

import (
	"testing"

	"github.com/ejedev/steam-update-notifier/internal/steamcmd"
)

func TestGetChangeNumber(t *testing.T) {
	id := 570
	result := steamcmd.GetChangeNumber(id)
	if result <= 0 {
		t.Errorf("GetChangeNumber(570) = %d; want >0", result)
	}
}

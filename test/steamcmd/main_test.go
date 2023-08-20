package steamcmd

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.NoLevel)
	m.Run()
}

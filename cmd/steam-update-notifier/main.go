package main

import (
	"flag"
	"os"
	"time"

	"github.com/ejedev/steam-update-notifier/internal/config"
	"github.com/ejedev/steam-update-notifier/internal/steamcmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	configPath := flag.String("config", "", "path to configuration file")
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	if len(*configPath) == 0 {
		log.Error().
			Msg("Please provide a config with --config.")
		os.Exit(1)
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	var config config.AppConfig
	config.GetConfig(*configPath)
	versionTracking := steamcmd.InitDataTracking(config.Apps)
	var interval int
	if config.Interval < 300 {
		interval = 300
	} else {
		interval = config.Interval
	}
	interval = 10
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		versionTracking = steamcmd.CheckChanges(versionTracking, config.Discord.Webhook, config.Discord.Enabled)
	}
}

# steam-update-notifier

Monitors Steam Apps and sends out Discord notifications when an update is spotted.

## Usage

```
$ ./steam-update-notifier -help

  -config string
    	path to configuration file
  -debug
    	sets log level to debug
```

An example configuration file can be found at `configs/config.yaml.example`.

Run the program with `./steam-update-notifier -config path/to/config.yaml`

The minimum probe time is set at 300 seconds (5 minutes.) This is because the API that this program relies on (steamcmd) caches its data for that long, making more frequent requests pointless.

## Install

### Binary
If a release has been made, you can download a binary [here.](https://github.com/ejedev/steam-update-notifier/releases/latest)

## Building
```
$ git clone https://github.com/ejedev/steam-update-notifier.git
$ cd steam-update-notifier
$ make
```

Your binary will be inside the `bin` folder.

## Acknowledgements

[SteamCMD API](https://www.steamcmd.net/) for the great API they provide.
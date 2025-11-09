package internal

import (
	"os"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type Config struct {
	configDir string
}

func ReadConfig() Config {
	c := Config{}

	checkConfigDirectory()

	return c
}

// checkConfigDirectory looks for a default directory to write the config file to
func checkConfigDirectory() {
	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	xdgConfigExists := !helpers.EqualsString("", xdgConfig)
	if !xdgConfigExists {
		files.MakeDirectory("~/.config/dragon/nasa/")
	}

	configFileExists, _ := files.DirExists("~/.config/dragon/nasa")
	if configFileExists {
		loadConfig()
	} else {
		writeDefaultConfig()
	}
}

func loadConfig() {
	// READ TOML File
}

func writeDefaultConfig() {
	// Default Config:
	// API:
	// API_KEY: demo -unset
}

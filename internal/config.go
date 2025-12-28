package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var (
	configDirectory   string
	screenshotsConfig string
	c                 *Config = &Config{}
)

// sets default config directory for the program
func checkConfigDirectory() {
	home, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	configDirectory = filepath.Join(home, "dragon/nasa")
}

func loadConfig() {
	fmt.Println("Loading Config")
	// READ TOML File
	//
	_, err := toml.DecodeFile(screenshotsConfig, c)
	if err != nil {
		panic(err)
	}
}

// writeDefaultConfig currently does nothing, other than print it's purpose
func writeDefaultConfig() {
	fmt.Println("Default config NOT YET copied to ", configDirectory)
	// Default Config:
	// API:
	// API_KEY: demo -unset
}

type (
	Config struct {
		Api map[string]API

		Hidef bool
	}
)

// ReadConfig reads the config from the config directory,
// by default on linux checks to "~/.config/dragon/nasa"
func ReadConfig() *Config {
	checkConfigDirectory() // sets configDirectory
	screenshotsConfig = configDirectory + "/wallpaper.toml"

	_, err := os.ReadFile(screenshotsConfig)

	// if screenshot.conf exists: load
	if err != nil {
		print(err)
		writeDefaultConfig()
	} else {
		loadConfig()
	}

	return c
}

// APIs returns the keys that can be used to reference each API
// in the map
func (c *Config) APIs() []string {
	keys := make([]string, len(c.Api))
	i := 0

	for k := range c.Api {
		keys[i] = k
		i++
	}

	return keys
}

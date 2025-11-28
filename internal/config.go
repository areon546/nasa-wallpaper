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

type (
	Config struct {
		Api map[string]API

		Hidef bool
	}
)

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

func writeDefaultConfig() {
	fmt.Println("Default config printed to ", configDirectory)
	// Default Config:
	// API:
	// API_KEY: demo -unset
}

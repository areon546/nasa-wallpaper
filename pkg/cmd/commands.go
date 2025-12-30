package cmd

import (
	"fmt"
	"os/exec"
)

func Notify(args ...string) {
	RunCommand("notify-send", args...)
}

func SetBackground(uri string) {
	// alternative bg args : see man feh /--bg-
	// --bg-scale
	// --bg-max
	// --bg-tile
	// --bg-fill
	RunCommand("/usr/bin/feh", "--bg-max", uri)
}

func RunCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)

	fmt.Println("Command ran: ", cmd.String())

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

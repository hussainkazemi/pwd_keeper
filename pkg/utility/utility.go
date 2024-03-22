package utility

import (
	"os"
	"os/exec"
	"runtime"
)

type OsType string

const (
	WINDOWS = "windows"
	MAC     = "darwin"
	LINUX   = "linux"
)

func ClearScreen() {
	userOs := runtime.GOOS
	var cmd *exec.Cmd
	switch userOs {
	case WINDOWS:
		cmd = exec.Command("cls")
	case MAC:
		cmd = exec.Command("clear")
	case LINUX:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

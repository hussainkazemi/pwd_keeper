package utility

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"pwsd_keeper/config"
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

func LoadFromEnv(key string) (string, error) {
	err := godotenv.Load(config.GetEnvFilePath())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return os.Getenv(key), nil
}

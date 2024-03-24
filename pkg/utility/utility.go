package utility

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"runtime"
)

type OsType string

const (
	WINDOWS = "windows"
	MAC     = "darwin"
	LINUX   = "linux"
	//read env file path form config file
	EnvPath = ".env"
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
	err := godotenv.Load(EnvPath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return os.Getenv(key), nil
}

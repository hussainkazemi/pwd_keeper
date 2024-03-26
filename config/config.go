package config

const (
	appVersion = "0.0.1"
	envPath    = ".env"
)

func GetAppVersion() string {

	return appVersion
}

func GetEnvFilePath() string {

	return envPath
}

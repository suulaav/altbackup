package main

import (
	"github.com/suulaav/altbackup/config"
)

func main() {
	appConfig := config.GetConfig()
	config.StartBackupService(appConfig)
}

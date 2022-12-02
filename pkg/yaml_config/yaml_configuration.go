package yaml_config

import (
	"fmt"
	"quick_passive_recon/pkg/args_proc"
)

type Settings struct {
	PingCount int
	MaxPort int
}

func GetRunConfig(command *args_proc.Command) Settings {
	fmt.Println("Retrieved using config: ", command)
	
	settings := Settings{}
	mode := command.Mode
	
	getModeSettings(mode, settings)

	//placeholder - remove
	settings.PingCount = 4
	settings.MaxPort = 1000

	return settings
}

func getModeSettings(mode string, settings Settings) {
	//TODO:
	//fetch settings for specified mode in yaml config
	fmt.Println(mode, settings)
}
package yaml_config

import (
	"fmt"
	"log"
	"io/ioutil"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/pkg/run_settings"
	"gopkg.in/yaml.v3"
)

/* TODO: Consider yaml mapping using struct instead of make() */
// type Config struct {
// 	Modes struct {
// 		Name struct {
// 			Description string `yaml:"description"`
// 			PingCount int `yaml:"pingCount"`
// 			MaxPort int	`yaml:"maxPort"`
// 		}
// 	}
// }

func GetRunConfig(command *args_proc.Command) run_settings.Settings {
	settings := run_settings.Settings{}
	mode := command.Mode
	
	getModeSettings(mode, &settings)

	return settings
}

func ShowConfig() {
	fmt.Println("settings.yaml content:")

	file, err := ioutil.ReadFile("settings.yaml")
	if err == nil {
		fmt.Println(string(file))
	} else {
		log.Fatal(err)
	}
}

func getModeSettings(mode string, settings *run_settings.Settings) {
	file, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Fatal(err)
	}

	config := make(map[interface{}]map[interface{}]map[interface{}]interface{})

	error := yaml.Unmarshal([]byte(file), &config)
	if error != nil {
		log.Fatal(error)
	}

	modes := config["modes"]
	activeMode := modes[mode]
	settings.Mode.PingCount = activeMode["PingCount"].(int)
	settings.Mode.MaxPort = activeMode["MaxPort"].(int)
}
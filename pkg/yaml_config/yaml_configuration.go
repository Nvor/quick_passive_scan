package yaml_config

import (
	"fmt"
	"log"
	"io/ioutil"
	"quick_passive_recon/pkg/args_proc"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Modes struct {
		Name struct {
			Description string `yaml:"description"`
			PingCount int `yaml:"pingCount"`
			MaxPort int	`yaml:"maxPort"`
		}
	}
}

type Mode struct {
	Description string
	PingCount int
	MaxPort int
}

type Settings struct {
	PingCount int
	MaxPort int
}

func GetRunConfig(command *args_proc.Command) Settings {
	fmt.Println("Retrieved using config: ", command)
	
	settings := Settings{}
	mode := command.Mode
	
	getModeSettings(mode, &settings)

	return settings
}

func getModeSettings(mode string, settings *Settings) {
	//TODO:
	//fetch settings for specified mode in yaml config
	fmt.Println(mode, settings)
	fmt.Println(settings)

	file, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))

	config := make(map[interface{}]map[interface{}]map[interface{}]interface{})

	error := yaml.Unmarshal([]byte(file), &config)
	if error != nil {
		log.Fatal(error)
	}

	modes := config["modes"]
	activeMode := modes[mode]
	settings.PingCount = activeMode["PingCount"].(int)
	settings.MaxPort = activeMode["MaxPort"].(int)

	fmt.Println("Obtained settings: ", settings)
}
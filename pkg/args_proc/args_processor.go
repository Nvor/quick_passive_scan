package args_proc

import (
	"fmt"
	"log"
	"os"
)

type Command struct {
	Url string
	Mode string
	Options []string
}

func ProcessArgs(args []string) *Command {
	if len(args) < 1 || len(args) > 2 {
		log.Println("Error: Invalid number of arguments")
		os.Exit(3)
	}

	command := Command{}
	for i := 0; i < len(args); i++ {
		switch i {
		case 0:
			ProcessUrl(&command, args[i])
		case 1:
			ProcessMode(&command, args[i])
		}
	}

	printCommand(&command)

	return &command
}

func ProcessUrl(command *Command, urlArg string) {
	if urlArg != "" {
		command.Url = urlArg
	} else {
		fmt.Println("Invalid url argument. Exiting..")
		os.Exit(3)
	}
}

func ProcessMode(command *Command, modeArg string) {
	if modeArg != "" {
		command.Mode = modeArg
	} else {
		fmt.Println("Invalid mode, using default.")
	}
}

func printCommand(command *Command) {
	fmt.Println("URL:", command.Url)
	fmt.Println("MODE:", command.Mode)
}

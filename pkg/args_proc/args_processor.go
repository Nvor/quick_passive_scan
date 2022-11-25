package args_proc

import (
	"log"
	//"os"
)

type Command struct {
	url string
	mode string
	options []string
}

func ProcessArgs(args []string) string {
	log.Println("Processing arguments:", args)

	log.Println("size", len(args))
	if len(args) < 1 || len(args) > 2 {
		log.Println("Error: Invalid number of arguments")
		return "replace with error - exit program"
	}

	log.Println("DS", args[0])

	for i := 0; i < len(args); i++ {
		if i == 0 {
			processUrlArg(args[i])
		}
	}

	var url string

	if args[0] != "" {
		url = args[0]
	} else {
		log.Fatal("Error: url not specified as an argument")
		os.Exit(3)
	}

	return url
}

func processUrlArg(urlArg string) string {
	log.Println("process url")
}

func processModeArg(modeArg string) {
	log.Println("process mode")
}
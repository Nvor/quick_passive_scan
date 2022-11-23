package args_proc

import (
	"log"
	"os"
)

func ProcessArgs(args []string) string {
	log.Println("Processing arguments:", args)
	var url string

	if args[0] != "" {
		url = args[0]
	} else {
		log.Fatal("Error: url not specified as an argument")
		os.Exit(3)
	}

	return url
}
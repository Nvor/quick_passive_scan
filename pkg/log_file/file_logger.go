package log_file

import (
	"fmt"
	"log"
	"os"
)

func Log(message string) {
	file, err := os.OpenFile("/tmp/qpc.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println(message)
}

func LogCommand(cmdArgs ...string) {
	cmdStr := "Executing command: "
	for i := 0; i < len(cmdArgs); i++ {
		cmdStr += fmt.Sprintf(" %s ", cmdArgs[i])
	}

	Log(cmdStr)
}

func LogError(message string) {
	log.Println("Error:", message)
}
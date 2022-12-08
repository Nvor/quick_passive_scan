package log_file

import (
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

func LogError(message string) {
	log.Println("Error:", message)
}
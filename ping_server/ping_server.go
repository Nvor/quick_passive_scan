package ping_server

import (
	"fmt"
	"log"
	"os/exec"
)

func PingServer(url string) {
	fmt.Println("Pinging url", url)
	var pcount int = 4

	arg1 := "ping"
	arg2 := fmt.Sprintf("-c %d", pcount)
	arg3 := url

	log.Println("Executing command:", arg1, arg2, arg3)

	cmd := exec.Command(arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	fmt.Println(string(stdout))
}
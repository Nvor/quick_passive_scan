package ping_server

import (
	"fmt"
	"log"
	"os/exec"
	"quick_passive_recon/pkg/args_proc"
)

func PingServer(command *args_proc.Command) {
	fmt.Println("Pinging url", command.Url)
	var pcount int = 4

	arg1 := "ping"
	arg2 := fmt.Sprintf("-c %d", pcount)
	arg3 := command.Url

	log.Println("Executing command:", arg1, arg2, arg3)

	cmd := exec.Command(arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	fmt.Println(string(stdout))
}
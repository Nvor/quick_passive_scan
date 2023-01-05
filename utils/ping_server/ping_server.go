package ping_server

import (
	"fmt"
	"log"
	"os/exec"
	"quick_passive_recon/pkg/log_file"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/pkg/run_settings"
)

func PingServer(command *args_proc.Command, settings *run_settings.Settings) {
	fmt.Println("Pinging url", command.Url)
	pcount := settings.Mode.PingCount

	arg1 := "ping"
	arg2 := fmt.Sprintf("-c %d", pcount)
	arg3 := command.Url

	log.Println("Executing command:", arg1, arg2, arg3)
	log_file.LogCommand(arg1, arg2, arg3)

	cmd := exec.Command(arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	fmt.Println(string(stdout))
}

package port_scan

import (
	"fmt"
	"log"
	"os/exec"
	"quick_passive_recon/pkg/log_file"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/pkg/run_settings"
)

func ScanWithNmap(command *args_proc.Command, settings *run_settings.Settings) {
	fmt.Println("nmap: scanning ports 0 -", settings.Mode.MaxPort)
	maxPort := settings.Mode.MaxPort

	arg1 := "nmap"
	arg2 := fmt.Sprintf("-p 0-%d", maxPort)
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
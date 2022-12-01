package port_scan

import (
	"fmt"
	"log"
	"os/exec"
	"quick_passive_recon/pkg/args_proc"
)

func ScanWithNmap(command *args_proc.Command, maxPort int) {
	fmt.Println("nmap: scanning ports 0 -", maxPort)

	arg1 := "nmap"
	arg2 := fmt.Sprintf("-p 0-%d", maxPort)
	arg3 := command.Url

	log.Println("Executing command:", arg1, arg2, arg3)

	cmd := exec.Command(arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	log.Println(string(stdout))
}
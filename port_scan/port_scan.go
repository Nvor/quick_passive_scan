package port_scan

import (
	"fmt"
	"log"
	"os/exec"
)

func ScanWithNmap(url string, maxPort int) {
	fmt.Println("nmap: scanning ports 0 -", maxPort)

	arg1 := "nmap"
	arg2 := fmt.Sprintf("-p 0-%d", maxPort)
	arg3 := url

	log.Println("Executing command:", arg1, arg2, arg3)

	cmd := exec.Command(arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	log.Println(string(stdout))
}
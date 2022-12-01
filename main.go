package main

import (
	"fmt"
	"os"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/ping_server"
	"quick_passive_recon/port_scan"
)

func main() {
	fmt.Println("Executing passive recon on url..")

	//Get arguments after executable
	args := os.Args[1:]
	command := args_proc.ProcessArgs(args)
	
	ping_server.PingServer(command)
	port_scan.ScanWithNmap(command, 5000)
}
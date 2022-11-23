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

	var url string
	args := os.Args[1:]

	url = args_proc.ProcessArgs(args)
	
	ping_server.PingServer(url)
	port_scan.ScanWithNmap(url, 5000)
}
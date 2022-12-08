package main

import (
	"fmt"
	"os"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/pkg/yaml_config"
	"quick_passive_recon/pkg/log_file"
	"quick_passive_recon/utils/ping_server"
	"quick_passive_recon/utils/port_scan"
)

func main() {
	fmt.Println("Executing passive recon on url..")

	//Get arguments and process them
	args := os.Args[1:]
	command := args_proc.ProcessArgs(args)

	//Set run config based on command and settings
	settings := yaml_config.GetRunConfig(command)

	log_file.Log("Starting scan with parameters:" + command.Url)

	//Start scan
	ping_server.PingServer(command, &settings)
	port_scan.ScanWithNmap(command, &settings)
}
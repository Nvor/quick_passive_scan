package main

import (
	"fmt"
	"quick_passive_recon/port_scan"
)

func main() {
	fmt.Println("Executing passive recon on url..")

	port_scan.ScanWithNmap("nvor.net", 5000)
}
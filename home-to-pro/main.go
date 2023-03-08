package main

import (
	"os/exec"
)

func main() {

	// convert win 10/11 home to pro
	cmd1 := exec.Command("cmd", "/C", "sc", "config", "LicenseManager", "start=", "auto", "&", "net", "start", "LicenseManager")
	cmd2 := exec.Command("cmd", "/C", "sc", "config", "wuauserv", "start=", "auto", "&", "net", "start", "wuauserv")
	cmd3 := exec.Command("cmd", "/C", "changepk.exe", "/productkey", "VK7JG-NPHTM-C97JM-9MPGT-3V66T")
	cmd4 := exec.Command("cmd", "/C", "exit")

	cmd1.Run()
	cmd2.Run()
	cmd3.Run()
	cmd4.Run()
}

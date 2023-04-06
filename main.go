package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Get product key from command-line argument
	var k1 string
	fmt.Print("Nhập key: ")
	fmt.Scanln(&k1)

	// Change directory to %windir%\system32
	err := exec.Command("cmd", "/c", "cd", "%windir%\\system32").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Clear screen
	err = exec.Command("cmd", "/c", "cls").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Reset license information
	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "/rilc").Run()
	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "/upk").Run()
	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "/cpky").Run()
	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "/ckms").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Start necessary services
	err = exec.Command("cmd", "/c", "sc", "config", "Winmgmt", "start=demand", "&", "net", "start", "Winmgmt").Run()
	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("cmd", "/c", "sc", "config", "LicenseManager", "start=auto", "&", "net", "start", "LicenseManager").Run()
	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("cmd", "/c", "sc", "config", "wuauserv", "start=auto", "&", "net", "start", "wuauserv").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Activate Windows with the provided key
	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "/ipk", k1).Run()
	if err != nil {
		fmt.Println(err)
	}

	// Copy updated licensing information to clipboard
	err = exec.Command("cmd", "/c", "clipup", "-v", "-o", "-altto", "c:\\").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Activate Windows
	err = exec.Command("cmd", "/c", "cscript", "slmgr.vbs", "-ato").Run()
	if err != nil {
		fmt.Println(err)
	}

	// Open Windows activation settings
	err = exec.Command("cmd", "/c", "start", "ms-settings:activation").Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Đã kích hoạt Windows thành công! Bạn có thể đóng cửa sổ này.")
	// dừng chương trình
	fmt.Scanln()
}

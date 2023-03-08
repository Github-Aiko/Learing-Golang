package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	// nhập key từ bàn phím
	var key string
	fmt.Print("Nhập key: ")
	fmt.Scanln(&key)
	match, _ := regexp.MatchString(`^[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}$`, key)
	if match {
		fmt.Println("Key format is correct")
	} else {
		fmt.Println("Key format is incorrect")
	}

	// xử lý key nhập vào
	cmd1 := exec.Command("cmd", "/C", "cd %windir%\\system32")
	cmd2 := exec.Command("cmd", "/C", "set k1="+key)
	cmd3 := exec.Command("cmd", "/C", "cls")
	cmd4 := exec.Command("cmd", "/C", "cscript slmgr.vbs /rilc")
	cmd5 := exec.Command("cmd", "/C", "cscript slmgr.vbs /upk")
	cmd6 := exec.Command("cmd", "/C", "cscript slmgr.vbs /cpky")
	cmd7 := exec.Command("cmd", "/C", "cscript slmgr.vbs /ckms")
	cmd8 := exec.Command("cmd", "/C", "sc config Winmgmt start=demand & net start Winmgmt")
	cmd9 := exec.Command("cmd", "/C", "sc config LicenseManager start= auto & net start LicenseManager")
	cmd10 := exec.Command("cmd", "/C", "sc config wuauserv start= auto & net start wuauserv")
	cmd11 := exec.Command("cmd", "/C", "@echo on&mode con: cols=20 lines=2")
	cmd12 := exec.Command("cmd", "/C", fmt.Sprintf("cscript slmgr.vbs /ipk %s", key))
	cmd13 := exec.Command("cmd", "/C", "@mode con: cols=100 lines=30")
	cmd14 := exec.Command("cmd", "/C", "clipup -v -o -altto c:\\")
	cmd15 := exec.Command("cmd", "/C", "cscript slmgr.vbs -ato")
	cmd16 := exec.Command("cmd", "/C", "start ms-settings:activation")

	cmd1.Run()
	cmd2.Run()
	cmd3.Run()
	cmd4.Run()
	cmd5.Run()
	cmd6.Run()
	cmd7.Run()
	cmd8.Run()
	cmd9.Run()
	cmd10.Run()
	cmd11.Run()
	cmd12.Run()
	cmd13.Run()
	cmd14.Run()
	cmd15.Run()
	cmd16.Run()
}

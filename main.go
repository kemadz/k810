package main

import (
	"fmt"
	"runtime"

	"github.com/kemadz/hid"
)

func main() {
	cmd := []byte{0x10, 0xff, 0x06, 0x15, 0x00, 0x00, 0x00}
	// cmd := []byte{0x10, 0xff, 0x06, 0x15, 0x01, 0x00, 0x00} // default
	// dis := hid.Enumerate(1133, 45847) // k811
	dis := hid.Enumerate(1133, 45849)
	var di *hid.DeviceInfo
	if len(dis) > 0 {
		if runtime.GOOS == "windows" {
			for _, v := range dis {
				if v.Usage == 1 && v.UsagePage == 65280 {
					di = &v
					break
				}
			}
		} else {
			di = &(dis[0])
		}
		dev, err := (*di).Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		cnt, err := dev.Write(cmd)
		if err != nil {
			fmt.Println(err)
		}
		if cnt > 0 {
			fmt.Println("Function key behavior is now on.")
		}
	} else {
		fmt.Println("No device found.")
	}
}

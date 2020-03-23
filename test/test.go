package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("10.0.0.0"))
	ip, ipmask, err := net.ParseCIDR("10.0.0.0/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip)
	fmt.Println(ipmask.String())
	fmt.Println(ipmask.IP)
	fmt.Println(ipmask.Mask)
}

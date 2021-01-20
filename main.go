package main

import (
	"github.com/vishvananda/netlink"
	"log"
	"net"
)

func main() {
	link, err := netlink.LinkByName("eth0")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Got link type: %s, struct: %+v", link.Type(), link.Attrs())

	ipConfig := &netlink.Addr{IPNet: &net.IPNet{
		IP: net.ParseIP("192.168.1.3"),
		Mask: net.CIDRMask(24, 32),
	}}
	if err = netlink.AddrAdd(link, ipConfig); err != nil {
		log.Fatalln(err)
	}
	log.Println("Configured eth0 interface")

	if err = netlink.LinkSetUp(link); err != nil {
		log.Fatalln(err)
	}
	log.Println("Set up link")
}

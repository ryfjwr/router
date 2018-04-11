package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

//
// SockAddrLinklayer test

func main() {
	interf, err := net.InterfaceByName("enp0s3")
	sock, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)

	if err != nil {
		fmt.Println("Socket Error")
	}
	// sockaddr ll
	var sall syscall.SockaddrLinklayer

	sall.Protocol = syscall.ETH_P_ARP
	sall.Ifindex = interf.Index
	sall.Hatype = syscall.ARPHRD_ETHER

	berr := syscall.Bind(sock, &sall)
	if berr != nil {
		fmt.Println("bind Error")
	}

	for {
		fmt.Println("hogehoge")
		time.Sleep(1 * time.Second)

		var buf []byte

		size, from, err := syscall.Recvfrom(sock, buf, 0)

		if size <= 0 {
			fmt.Println("no packet")
		}
		if err != nil {
			fmt.Println("recvfrom error")
		}

		fmt.Println(from)
	}

	defer syscall.Close(sock)
}

package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

func main() {
	// imported
	// http://www.geekpage.jp/programming/linux-network/tcp-1.php
	server := syscall.SockaddrInet4{Port: 12345}
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	if err != nil {
		fmt.Println("sock init error")
		fmt.Println(err.Error())
	}

	copy(server.Addr[:], net.ParseIP("0.0.0.0").To4())
	err = syscall.Connect(sock, &server)

	if err != nil {
		fmt.Println("sock connect Error")
		fmt.Println(err.Error())
	}

	var buf []byte
	_, err = syscall.Read(sock, buf)

	if err != nil {
		fmt.Println("Read Error")
		fmt.Println(err.Error())
	}

	fmt.Println(buf)

	time.Sleep(2 * time.Second)

	syscall.Close(sock)
}

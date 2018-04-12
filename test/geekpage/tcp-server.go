package main

import (
	"fmt"
	"net"
	"syscall"
)

// http://www.geekpage.jp/programming/linux-network/tcp-1.php

func main() {
	addr := syscall.SockaddrInet4{Port: 12345}
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	if err != nil {
		fmt.Println("sock init error")
	}

	// https://gist.github.com/tevino/3a4f4ec4ea9d0ca66d4f
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())

	syscall.Bind(sock, &addr)

	err = syscall.Listen(sock, 5)

	if err != nil {
		fmt.Println("listen Error")
		fmt.Println(err.Error())
	}

	sock1, _, err := syscall.Accept(sock)

	if err != nil {
		fmt.Println("error occured")
	}

	_str := "HELLO"
	str := []byte(_str)

	fmt.Println("byte order")
	fmt.Println(str)

	ret, err := syscall.Write(sock, str)

	if err != nil {
		fmt.Println("Write Error")
		fmt.Println(ret)
		fmt.Println(err.Error())
	}

	fmt.Println("exit")
	syscall.Close(sock)
	syscall.Close(sock1)
}

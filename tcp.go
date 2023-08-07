package main

import (
	"io"
	"log"
	"net"
)

func startTCPForwarder(target string) {
	service := ":22"
	remote := target + ":22"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Println("Error resolving TCP address:", err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("Error setting up TCP listener:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn, remote)
	}
}

func handleClient(conn net.Conn, remote string) {
	defer conn.Close()
	remoteConn, err := net.Dial("tcp", remote)
	if err != nil {
		log.Println("Error dialing remote address:", err)
		return
	}
	defer remoteConn.Close()

	go func() {
		_, err := io.Copy(conn, remoteConn)
		if err != nil {
			log.Println("Error copying data from remote to local:", err)
		}
	}()

	_, err = io.Copy(remoteConn, conn)
	if err != nil {
		log.Println("Error copying data from local to remote:", err)
	}
}

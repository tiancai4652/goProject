package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port string
}

func NewServer(ip string, port string) *Server {
	server := &Server{ip, port}
	return server
}

func (This *Server) handle(conn net.Conn) {
	fmt.Println("链接建立成功")
}

func (this *Server) Start() {
	//Create listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accepterr:", err)
			continue
		}

		//利用conn做事情
		go this.handle(conn)

	}
}

package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip        string
	Port      string
	Message   chan string
	OnlineMap map[string]*User
	mplock    sync.RWMutex
}

func NewServer(ip string, port string) *Server {
	server := &Server{ip, port, make(chan string), make(map[string]*User), sync.RWMutex{}}
	return server
}

func (this *Server) Broadcast(user *User, msg string) {
	message := "[" + user.Address + "]" + user.Name + ":" + msg
	this.Message <- message
}

func (This *Server) handle(conn net.Conn) {
	fmt.Println("链接建立成功")
	user := NewUser(conn)

	//添加到在线用户map
	This.mplock.Lock()
	This.OnlineMap[user.Name] = user
	This.mplock.Unlock()

	//通过带缓存chaneel用户上线
	This.Broadcast(user, "用户上线了")
}

func (this *Server) xx() {
	for {
		msg := <-this.Message
		this.mplock.Lock()

		for _, user := range this.OnlineMap {
			user.C <- msg
		}

		this.mplock.Unlock()
	}
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

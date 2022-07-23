package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	Conn    net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddress := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddress,
		Address: userAddress,
		C:       make(chan string),
		Conn:    conn,
	}

	go user.ListenMessage()

	return user
}

func (this *User) ListenMessage() {

	for {
		msg := <-this.C
		this.Conn.Write([]byte(msg + "\n"))
	}
}

package main

import (
	"errors"
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) (*Server, error) {
	if ip == "" || port == 0 {
		fmt.Println("ip 或者端口错误")
		return &Server{}, errors.New("ip 或者端口错误")
	}
	return &Server{
		ip,
		port,
	}, nil
}
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")
}
func (this *Server) Start() {
	// socket listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("Start failed, err:", err)
		return
	}

	// close listener
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		// handler
		go this.Handler(conn)
	}

}

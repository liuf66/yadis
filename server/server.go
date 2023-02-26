package yadis

import (
	"net"
)

type Server struct {
}

func (s *Server) Serve() {
	var (
		listener net.Listener
		conn     net.Conn
		err      error
	)
	listener, err = net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		ylog("net listen, err=%v", err)
		return
	}
	for {
		conn, err = listener.Accept()
		if err != nil {
			ylog("listenner accept got err=%v", err)
			continue
		}

		ylog("new conn, addr=%v", conn.RemoteAddr())
		_ = conn.Close()
	}
}

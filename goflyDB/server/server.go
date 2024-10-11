package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn   net.Conn
	addr   net.Addr
	server *Server
}
type Server struct {
	addr      net.Addr
	clientMap map[net.Addr]*Client
	msgBus    chan Command
}

func NewServer(addr net.Addr) *Server {
	return &Server{
		addr:      addr,
		clientMap: make(map[net.Addr]*Client),
		msgBus:    make(chan Command, 1024),
	}
}

func (s *Server) start() {
	fmt.Println("server is starting ")

	listner, err := net.Listen(s.addr.Network(), s.addr.String())
	if err != nil {
		fmt.Println("failed to start the server : ", err)
		return
	}
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("error while accepting conneciton", err)
			continue
		}
		client := s.newClient(conn)
		s.clientMap[conn.LocalAddr()] = client
		go client.handleConnection()
	}

}
func (s *Server) newClient(conn net.Conn) *Client {
	return &Client{
		conn:   conn,
		addr:   conn.LocalAddr(),
		server: s,
	}
}
func (s *Server) pushCommand(command Command) {

	s.msgBus <- command
}
func (c *Client) handleConnection() {
	// handler incoming request
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			fmt.Println("send message to client for err")
			continue
		}
		if strings.HasPrefix(msg, "get") {
			key, found := strings.CutPrefix(msg, "get")
			if found {

				cmd := Command{
					id:     1,
					client: c,
					args:   []string{strings.TrimSpace(key)},
				}
				c.server.pushCommand(cmd)

			}
		}
		if strings.HasPrefix(msg, "put") {

		}
		if strings.HasPrefix(msg, "exit") {

		}
	}

}

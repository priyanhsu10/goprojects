package server

import (
	"fmt"
	"net"
)


type Client struct {
  conn  net.Conn,
  addr net.Addr,
}
type Server struct {
	addr      net.Addr
	clientMap map[net.Addr]*Client
	msgBus    chan string
}

func NewServer(addr net.Addr) *Server {
	return &Server{
		addr:      addr,
		clientMap: make(map[net.Addr]net.Conn),
		msgBus:    make(chan string, 1024),
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
    client:=s.newClient(conn)
    s.clientMap[conn.LocalAddr()]=client
    go client.handleConnection(conn)
	}

}
func (s *Server) newClient(conn net.Conn) *Client{
  return &Client{
    conn: conn,
    addr: conn.LocalAddr().String(),
  }
}
func (s *Client) handleConnection() {
//handler incoming request
}


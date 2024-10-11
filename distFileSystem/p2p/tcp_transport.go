package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddr string
	listener   net.Listener
	mu         sync.RWMutex
	peers      map[net.Addr]Peer
	decoder    Decoder
	shakeHands HandleshakeFunc
}

// TCPPeer represent the remote node over a tcp establish connection
type TCPPeer struct {
	//conn is underlying connection
	conn net.Conn
	//if we dial a connection => outbout  == true
	//fi we accept the connet => outbound == false
	outbound bool
}

func NewTcpPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}
func NewTCPTransport(listnerAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listnerAddr,
		shakeHands: NOPHandShakeFunc,
	}
}
func (t *TCPTransport) ListnerAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	go t.acceptLoop()
	return nil
}
func (t *TCPTransport) acceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error : %s\n", err)
		}
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer := NewTcpPeer(conn, true)
	if err := t.shakeHands(peer); err != nil {

	}
	msg := &Temp{}
	for {
		t.decoder.Decode(conn, msg)
	}
	fmt.Print("new incomming connection ..   %+v\n", peer)
}

package p2p

type Peer interface {
}

// trnaport thtat handle the comumnicaotin l
// between the nodes in the netewrok
// form of (tcop , udp ,websocket...)
type Transport interface {
	ListenAndAccept() error
}

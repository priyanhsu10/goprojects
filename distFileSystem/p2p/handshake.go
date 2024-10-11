package p2p

// handshake func ....?
type HandleshakeFunc func(Peer) error

func NOPHandShakeFunc(Peer) error {
	return nil
}

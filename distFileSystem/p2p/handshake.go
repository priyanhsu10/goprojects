package p2p

type HandleshakeFunc func(any) error

func NOPHandShakeFunc(any) error {
	return nil
}

package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTrasport(t *testing.T) {
	listenerAddr := ":4000"
	tr := NewTCPTransport(listenerAddr)
	assert.Equal(t, tr.listenAddr, listenerAddr)
	assert.Nil(t, tr.ListnerAndAccept())
}

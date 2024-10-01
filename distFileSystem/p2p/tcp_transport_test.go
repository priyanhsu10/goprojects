package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTrasport(t *testing.T) {
	listenerAddr := ":4000"
	tr := NewTCPTransport(listenerAddr)
	assert.Equal(t, tr.listenAddr, listenerAddr)
	assert.Nil(t, tr.ListnerAndAccept())
}

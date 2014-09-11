package bitswap

import (
	context "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/go.net/context"
	pubsub "github.com/jbenet/go-ipfs/Godeps/_workspace/src/github.com/tuxychandru/pubsub"

	blocks "github.com/jbenet/go-ipfs/blocks"
	u "github.com/jbenet/go-ipfs/util"
)

type notifications struct {
	wrapped *pubsub.PubSub
}

func newNotifications() *notifications {
	const bufferSize = 16
	return &notifications{pubsub.New(bufferSize)}
}

func (ps *notifications) Publish(block *blocks.Block) {
	topic := string(block.Key())
	ps.wrapped.Pub(block, topic)
}

// Sub returns a one-time use |blockChannel|. |blockChannel| returns nil if the
// |ctx| times out or is cancelled
func (ps *notifications) Subscribe(ctx context.Context, k u.Key) <-chan *blocks.Block {
	topic := string(k)
	subChan := ps.wrapped.Sub(topic)
	blockChannel := make(chan *blocks.Block)
	go func() {
		defer close(blockChannel)
		select {
		case val := <-subChan:
			block, ok := val.(*blocks.Block)
			if ok {
				blockChannel <- block
			}
		case <-ctx.Done():
			ps.wrapped.Unsub(subChan, topic)
		}
	}()
	return blockChannel
}

func (ps *notifications) Shutdown() {
	ps.wrapped.Shutdown()
}

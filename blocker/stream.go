package blocker

import (
	"context"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
)

func (b *Blocker) Watch(ctx context.Context) {
	stream := b.tw.PublicStreamSample(nil)
	defer stream.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-stream.C:
			if !ok {
				log.Fatalf("Stream chan %v is closed.", stream.C)
			}
			switch tweet := msg.(type) {
			case anaconda.Tweet:
				b.judge(&tweet)
			case anaconda.StatusDeletionNotice:
				// pass
			default:
				fmt.Printf("unknown type(%T) : %v \n", msg, msg)
			}
		}
	}

}

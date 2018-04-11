package blocker

import (
	"context"
	"github.com/ChimeraCoder/anaconda"
)

type Blocker struct {
	tw      *anaconda.TwitterApi
	friends map[int64]bool
}

func NewBlocker(tw *anaconda.TwitterApi) *Blocker {
	return &Blocker{
		tw:      tw,
		friends: make(map[int64]bool),
	}
}

func (b *Blocker) Prepare(ctx context.Context) <-chan bool {
	result := make(chan bool)
	go func() {
		defer close(result)
		if ok := <-b.getFollowers(ctx); !ok {
			result <- false
		}
		if ok := <-b.getFollowings(ctx); !ok {
			result <- false
		}
		result <- true
	}()
	return result
}

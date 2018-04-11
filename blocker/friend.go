package blocker

import (
	"context"
	log "github.com/Sirupsen/logrus"
)

func (b *Blocker) getFollowers(ctx context.Context) <-chan bool {
	result := make(chan bool)
	ch := b.tw.GetFollowersIdsAll(nil)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				result <- false
				return
			case page, ok := <-ch:
				if !ok {
					result <- true
					return
				}
				if page.Error != nil {
					log.Fatal(page.Error)
				}
				for _, user := range page.Ids {
					b.friends[user] = true
				}
			}
		}
	}()
	return result
}

func (b *Blocker) getFollowings(ctx context.Context) <-chan bool {
	result := make(chan bool)
	ch := b.tw.GetFriendsIdsAll(nil)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				result <- false
				return
			case page, ok := <-ch:
				if !ok {
					result <- true
					return
				}
				if page.Error != nil {
					log.Fatal(page.Error)
				}
				for _, user := range page.Ids {
					b.friends[user] = true
				}
			}
		}
	}()
	return result
}

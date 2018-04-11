package blocker

import (
	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
	"github.com/ledyba/go-block-them-all/conf"
	"strings"
)

func (b *Blocker) judge(tweet *anaconda.Tweet) {
	user := &tweet.User
	if b.friends[user.Id] {
		log.Printf("(%s, %s, %s) wan't blocked: following or follower", user.Name, user.ScreenName, user.IdStr)
		return
	}
	if user.Verified {
		b.block(user, "Verified")
		return
	}
	if tweet.PossiblySensitive {
		b.block(user, "Sensitive")
		return
	}
	text := tweet.Text
	// specific ng words.
	for _, kw := range conf.NGText {
		if strings.Contains(text, kw) {
			b.block(user, "by tweet.")
			return
		}
	}
	for _, kw := range conf.NGScreenName {
		if strings.Contains(user.ScreenName, kw) {
			b.block(user, "by screen name.")
			return
		}
	}
	for _, kw := range conf.NGDescription {
		if strings.Contains(user.Description, kw) {
			b.block(user, "by description.")
			return
		}
	}
	for _, kw := range conf.NGURL {
		if strings.Contains(user.URL, kw) {
			b.block(user, "by url.")
			return
		}
	}
	//common
	for _, kw := range conf.NGCommon {
		if strings.Contains(text, kw) {
			b.block(user, "by tweet.")
			return
		} else if strings.Contains(user.ScreenName, kw) {
			b.block(user, "by screen name")
			return
		} else if strings.Contains(user.Description, kw) {
			b.block(user, "by description.")
			return
		}
	}
}

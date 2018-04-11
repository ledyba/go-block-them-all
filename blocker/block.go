package blocker

import (
	"github.com/ChimeraCoder/anaconda"
	log "github.com/Sirupsen/logrus"
)

func (b *Blocker) block(user *anaconda.User, cause string) {
	_, err := b.tw.BlockUserId(user.Id, nil)
	if err != nil {
		log.Errorf("Failed to block (err=\"%v\"): (%s, %s, %s)", err, user.Name, user.ScreenName, user.IdStr)
	}
	log.Printf("Blocking (%s, %s, %s): %s", user.Name, user.ScreenName, user.IdStr, cause)
}

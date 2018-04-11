package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"context"
	"github.com/ChimeraCoder/anaconda"
	"github.com/fatih/color"
	"github.com/ledyba/go-block-them-all/blocker"
	"github.com/ledyba/go-block-them-all/conf"
)

//go:generate bash geninfo.sh

func mainLoop(sig <-chan os.Signal) os.Signal {
	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	api := anaconda.NewTwitterApi(conf.OAuthToken, conf.OAuthSecret)
	defer api.Close()
	ctx, cancel := context.WithCancel(context.Background())
	b := blocker.NewBlocker(api)
	go func() {
		if ok:= <-b.Prepare(ctx); !ok {
			return
		}
		b.Watch(ctx)
	}()
	select {
	case s := <-sig:
		cancel()
		return s
	}
}

func printLogo() {
	log.Info("****************************************")
	log.Info(color.BlueString("  block-them-all  "))
	log.Info("****************************************")
	log.Infof("Build at: %s", color.MagentaString("%s", buildAt()))
	log.Infof("Git Revision: \n%s", color.MagentaString("%s", gitRev()))
}

func main() {
	//var err error

	printLogo()
	flag.Parse()
	log.Info("----------------------------------------")
	log.Info("Initializing...")
	log.Info("----------------------------------------")

	log.Info(color.GreenString("                                    [OK]"))

	log.Info("----------------------------------------")
	log.Info("Initialized.")
	log.Info("----------------------------------------")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := mainLoop(sig)
	log.Fatalf("Signal (%v) received, stopping\n", s)
}

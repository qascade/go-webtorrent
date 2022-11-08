package main

import (
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func main() {
	logger.WithFields(log.Fields{
		"starting": "Webtorrent",
	}).Info("this is the start of webtorrent main")
}

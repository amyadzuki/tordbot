package main

import (
	"strconv"
	"time"
)

var PlayingFrequent = []string{
	"Prefix is tord",
	"Try tord help",
	"with a bunch of cuties",
}

func cyclePlayingStatus() {
	var servers string
	var guilds int64
	for {
		guilds = 0
		guilds += int64(len(session.State.Guilds))
		servers = strconv.FormatInt(guilds, 10) + " Servers"
		now := time.Now()
		month := int(now.Month()) - 1

		session.UpdateStatus(0, servers)
		time.Sleep(10 * time.Second)

		for _, playing := range Playing {
			session.UpdateStatus(0, playing)
			time.Sleep(10 * time.Second)
		}
	}
}

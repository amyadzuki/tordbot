package main

import (
	"strconv"
	"time"
)

var Playing = []string{
	"Prefix is tord",
	"Try tord help",
	"with a bunch of cuties",
}

func cyclePlayingStatus() {
	var servers string
	var guilds int64
	for {
		guilds = 0
		guilds += int64(len(Session.State.Guilds))
		servers = strconv.FormatInt(guilds, 10) + " Servers"

		Session.UpdateStatus(0, servers)
		time.Sleep(10 * time.Second)

		for _, playing := range Playing {
			Session.UpdateStatus(0, playing)
			time.Sleep(10 * time.Second)
		}
	}
}

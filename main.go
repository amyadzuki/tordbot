package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var Session *discordgo.Session

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "/path/to/token.dat /path/to/database.db")
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", os.Args[2])
	if err != nil {
		panic(err)
	}
	DB = db
	Session, err = discordgo.New("Bot " + strings.TrimSpace(string(b)))
	if err != nil {
		panic(err)
	}
	Session.AddHandler(onMessageCreate)
	Session.AddHandler(onMessageReactionAdd)
	Session.AddHandler(onMessageReactionRemove)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	if err := Session.Open(); err != nil {
		panic(err)
	}
	defer Session.Close()
	go cyclePlayingStatus()
	<-sc
}

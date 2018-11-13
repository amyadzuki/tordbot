package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func main(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage:", args[0], "/path/to/token.dat /path/to/database.db")
	}
	b, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", args[2])
	if err != nil {
		panic(err)
	}
	session, err := discordgo.New("Bot " + strings.TrimSpace(string(b)))
	if err != nil {
		panic(err)
	}
	session.AddHandler(onMessageCreate)
	session.AddHandler(onMessageReactionAdd)
	session.AddHandler(onMessageReactionRemove)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
}

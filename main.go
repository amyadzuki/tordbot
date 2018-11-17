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
		fmt.Println("Usage:", os.Args[0], "init /path/to/database.db")
		fmt.Println("Usage:", os.Args[0], "/path/to/token.dat /path/to/database.db")
	}
	var err error
	DB, err = sql.Open("sqlite3", os.Args[2])
	if err != nil {
		panic(err)
	}
	if os.Args[1] == "init" {
		initDB()
		return
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	Session, err = discordgo.New("Bot " + strings.TrimSpace(string(b)))
	if err != nil {
		panic(err)
	}
	Session.AddHandler(onMessageCreate)
	Session.AddHandler(onMessageReactionAdd)
	Session.AddHandler(onMessageReactionRemove)
	Session.AddHandler(onVoiceStateUpdate)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	if err := Session.Open(); err != nil {
		panic(err)
	}
	defer Session.Close()
	go cyclePlayingStatus()
	<-sc
}

func initDB() {
	ExecDB(
		`CREATE TABLE "Prompts" (` +
		"\n\t" + `"guild"  INTEGER NOT NULL,` +
		"\n\t" + `"blame"  INTEGER NOT NULL,` +
		"\n\t" + `"dare"   INTEGER NOT NULL,` +
		"\n\t" + `"nsfw"   INTEGER NOT NULL,` +
		"\n\t" + `"at"     INTEGER NOT NULL,` +
		"\n\t" + `"score"  INTEGER NOT NULL,` +
		"\n\t" + `"flags"  INTEGER NOT NULL,` +
		"\n\t" + `"prompt" TEXT    NOT NULL,` +
		"\n\t" + `"z" INTEGER NOT NULL DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))` +
		"\n" + `);`,
		`CREATE TABLE "Channels" (` +
		"\n\t" + `"vid" INTEGER PRIMARY KEY NOT NULL,` +
		"\n\t" + `"cid" INTEGER NOT NULL,` +
		"\n\t" + `"z" INTEGER NOT NULL DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))` +
		"\n" + `);`,
		`CREATE TABLE "Users" (` +
		"\n\t" + `"uid"   INTEGER PRIMARY KEY NOT NULL,` +
		"\n\t" + `"cid"   INTEGER NOT NULL,` +
		"\n\t" + `"score" INTEGER NOT NULL,` +
		"\n\t" + `"hp"    INTEGER NOT NULL,` +
		"\n\t" + `"mp"    INTEGER NOT NULL,` +
		"\n\t" + `"nsfw"  INTEGER NOT NULL,` +
		"\n\t" + `"at"    INTEGER NOT NULL,` +
		"\n\t" + `"items" TEXT    NOT NULL,` +
		"\n\t" + `"z" INTEGER NOT NULL DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))` +
		"\n" + `);`,
		// Truths
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 0, 0xffff,    5, 0x0000, 'What is your name?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 0, 0xffff,    5, 0x0000, 'How old are you?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 0, 0xffff,   10, 0x0000, 'What exactly are you wearing?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 0, 0xffff,   25, 0x0000, 'What is the last thing you texted?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 0, 0xffff,   50, 0x0000, 'The previous person to go gets to ask you any SFW truth and you must answer it.  If this is the first round of the game then this question is a freebie.')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 1, 0xffff,   10, 0x0000, 'When is the last time you kissed someone?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 1, 0xffff,   25, 0x0000, 'Do you have a crush on anyone in this server?  If so, who is it?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 1, 0xffff,   50, 0x0000, 'What color are your underwear?')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   25, 0x0000, 'What is one of your kinks?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'When is the last time you did erp or sexted with someone?')`,

		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'Have you ever performed a strip tease?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'Do you like being bitten?  List all the places you have been bitten before.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'Have you ever tried rope play?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'Have you ever tried pee play?')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,   50, 0x0000, 'Have you ever masturbated live on web cam?')`,

		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,  100, 0x0000, 'Describe in detail the kinkiest thing you have ever done.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 2, 0xffff,  100, 0x0000, 'The previous person to go gets to ask you any NSFW truth and you must answer it.  If this is the first round of the game then this question is a freebie.')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 3, 0xffff,   100, 0x0000, 'Describe in detail a NSFW event in your past.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 3, 0xffff,   100, 0x0000, 'Describe in detail a NSFW fantasy of yours.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 0, 3, 0xffff,   100, 0x0000, 'Describe in detail the last NSFW dream you had.')`,
		// Dares
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 0, 0xffff,    5, 0x0000, 'Take a drink of water (hydration is important).')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 0, 0xffff,   10, 0x0000, 'Screenshot the top of your DM list and post it to this channel.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 0, 0xffff,   25, 0x0000, 'Take a selfie right now and post it to this channel.')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 1, 0xffff,   10, 0x0000, 'Write a flirty message to someone in this server.  Post it in a general chat channel in this server and tag them.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 1, 0xffff,   25, 0x0000, 'Take a SFW but seductive selfie right now and post it to this channel.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 1, 0xffff,   50, 0x0000, 'Take a picture of your tummy right now and post it to this channel.')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,   25, 0x0000, 'Change your nickname in this server to "slut".  Don't change it back until the game ends.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,   25, 0x0000, 'Take a picture of your thighs right now and post it to this channel.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,   50, 0x0000, 'Name and remove one article of clothing.  You must not put on any more clothing until the end of the game.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,  100, 0x0000, 'Unmute voice if muted.  Then moan audibly.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,  100, 0x0000, 'Walk outside for 15 seconds in whatever you are wearing.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,  250, 0x0000, 'Name and remove one article of clothing.  You must not put on any more clothing until the end of the game.  Then, take a full-body selfie and post it to this channel.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff,  250, 0x0000, 'Unmute voice if muted.  Then masturbate rapidly for 60 seconds.')`,
		//
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff, 500, 0x0000, 'Send a NSFW selfie to someone playing.')`,
		`INSERT INTO "Prompts" ("guild", "blame", "dare", "nsfw", "at", "score", "flags", "prompt") ` +
		`VALUES (0, 0, 1, 2, 0xffff, 500, 0x0000, 'Attempt to satisfy the last kink mentioned by someone during this game.  If none has been mentioned, take requests.')`,
	)
}

func ExecDB(sqls ...string) error {
	for _, sql := range sqls {
		stmt, err := DB.Prepare(sql)
		if err != nil {
			return err
		}
		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}


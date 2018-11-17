package main

import (
	"github.com/bwmarrin/discordgo"
)

const (
	AT_UNSPECIFIED = uint32(1) << iota
	AT_HOMEALONE
	AT_HOME
	AT_WORK
	AT_SCHOOL
	AT_CAR
)

const AT_ANYWHERE = uint32(0xffff)

func addPrompt(guildID, channelID, author string, dare int, nsfw, at uint32, prompt string) {
	stmt, err := DB.Prepare(`INSERT INTO "Prompts" ("guild", "dare", "nsfw", "at", "prompt", "blame") ` +
		`VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	_, err = stmt.Exec(guildID, dare, nsfw, at, prompt, author)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Exec: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	Session.ChannelMessageSend(channelID,
		"Got it!  Thanks for your contribution and please add some more :)")
	return
}

func givePrompt(guildID, channelID, author string, dare int, nsfw, at uint32, ent uint64) {
	stmt, err := DB.Prepare(`SELECT "prompt", "blame" FROM "Prompts" WHERE `+
		`("guild" = 0 OR "guild" = ?) AND "dare" = ? AND "nsfw" = ? AND ` +
		`(("at" & ?) <> 0) ORDER BY random() LIMIT 1`)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error querying prompt during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	rows, err := stmt.Query(guildID, dare, nsfw, at)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error querying prompt during SQL Query: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	defer rows.Close()
	if !rows.Next() {
		Session.ChannelMessageSend(channelID, "No prompts matched the criteria.")
		return
	}
	var prompt, blame string
	err = rows.Scan(&prompt, &blame)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error querying prompt during SQL Scan: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	Session.ChannelMessageSend(channelID, "<@!" + author + ">, " + prompt + " (added by " + blame + ")")
	return
}

func install(c *discordgo.Channel, vis string) {
}

func deinstall(c *discordgo.Channel) {
}

func voiceStateUpdate(vs *discordgo.VoiceState) {
	stmt, err := DB.Prepare(`SELECT "cid" FROM "Channels" WHERE "vid" = ? LIMIT 1`)
	if err != nil {
		return
	}
	rows, err := stmt.Query(vs.ChannelID)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		voiceStateLeft(vs)
		return
	}
	var vid string
	err = rows.Scan(&vid)
	if err != nil {
		return
	}
	Session.ChannelMessageSend(vid, "<@!" + vs.UserID + "> joined the game!")
}

func voiceStateLeft(vs *discordgo.VoiceState) {
	stmt, err := DB.Prepare(`SELECT "cid" FROM "Users" WHERE "uid" = ? LIMIT 1`)
	if err != nil {
		return
	}
	rows, err := stmt.Query(vs.ChannelID)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		return
	}
	var cid string
	err = rows.Scan(&cid)
	if err != nil {
		return
	}
	removeFromGame(vs.UserID, cid)
}

func removeFromGame(uid, cid string) {
	stmt, err := DB.Prepare(`DELETE FROM "Users" WHERE "uid" = ?`)
	if err != nil {
		if len(cid) > 0 {
			Session.ChannelMessageSend(cid,
				"Error deleting user during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(guildID, dare, nsfw, at, prompt, author)
	if err != nil {
		if len(cid) > 0 {
			Session.ChannelMessageSend(cid,
				"Error deleting user during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	if len(cid) > 0 {
		Session.ChannelMessageSend(cid, "<@!" + vs.UserID + "> left the game!")
	}
	return
}

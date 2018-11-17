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

func install(c *discordgo.Channel, tail string) {
	for len(tail) > 0 && (tail[0] < '0' || tail[0] > '9') {
		tail = tail[1:]
	}
	for len(tail) > 0 && (tail[len(tail) - 1] < '0' || tail[len(tail) - 1] > '9') {
		tail = tail[:len(tail) - 1]
	}
	channelIDTmp := tail
	voice, err := Session.State.Channel(channelIDTmp)
	if err != nil {
		voice, err = Session.Channel(channelIDTmp)
		if err != nil {
			Session.ChannelMessageSend(c.ID, "Oops, " + channelIDTmp +
				" does not look like a channel ID number.")
			return
		}
	}

	cleanForInstall(voice.ID, c.ID)

	stmt, err := DB.Prepare(`INSERT INTO "Channels" ("vid", "cid") VALUES (?, ?)`)
	if err != nil {
		Session.ChannelMessageSend(c.ID,
			"Error installing during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	_, err = stmt.Exec(voice.ID, c.ID)
	if err != nil {
		Session.ChannelMessageSend(c.ID,
			"Error installing during SQL Exec: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	Session.ChannelMessageSend(c.ID,
		"Installation successful!")
	return
}

func cleanForInstall(vid, cid string) {
	stmt, err := DB.Prepare(`DELETE FROM "Channels" WHERE "vid" = ? OR "cid" = ?`)
	if err != nil {
		Session.ChannelMessageSend(cid,
			"Error cleaning for install during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	_, err = stmt.Exec(vid, cid)
	if err != nil {
		Session.ChannelMessageSend(cid,
			"Error cleaning for install during SQL Exec: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
}

func deinstall(c *discordgo.Channel) {
	stmt, err := DB.Prepare(`DELETE FROM "Channels" WHERE "cid" = ?`)
	if err != nil {
		Session.ChannelMessageSend(c.ID,
			"Error deinstalling during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	_, err = stmt.Exec(c.ID)
	if err != nil {
		Session.ChannelMessageSend(c.ID,
			"Error deinstalling during SQL Exec: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	Session.ChannelMessageSend(c.ID,
		"Deinstallation successful!")
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
	var cid string
	err = rows.Scan(&cid)
	if err != nil {
		return
	}

	update(cid, vs.UserID)
}

func update(cid, uid string) {
	stmt, err := DB.Prepare(`SELECT "cid" FROM "Players" WHERE "uid" = ? LIMIT 1`)
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
	var oldcid string
	err = rows.Scan(&oldcid)
	if err != nil {
		return
	}
	if cid == oldcid {
		return
	}
	leave(oldcid, uid)
	join(cid, uid)
}

func leave(cidOptional, uid string) {
	stmt, err := DB.Prepare(`DELETE FROM "Players" WHERE "uid" = ?`)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error deleting player during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(uid)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error deleting player during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	if len(cidOptional) > 0 {
		Session.ChannelMessageSend(cidOptional, "<@!" + uid + "> left the game!")
	}
	return
}

func join(gid, cid, uid string) {
	stmt, err := DB.Prepare(`INSERT INTO "Players" ` +
		`("uid", "gid", "cid", "score", "hp", "mp") ` +
		`VALUES (?, ?, ?, 0, 12, 0)`)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error inserting player during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(uid, gid, cid)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error inserting player during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	if len(cidOptional) > 0 {
		Session.ChannelMessageSend(cidOptional, "<@!" + uid + "> joined the game!")
	}
	return
}

func usGet(cidOptional, uid string) (maxnsfw, at uint32) {
	// Set the defaults in case there's a problem
	maxnsfw = 3
	at = AT_HOME
	stmt, err := DB.Prepare(`SELECT "maxnsfw", "at" FROM "UserSettings" WHERE "uid" = ? LIMIT 1`)
	if err != nil {
		if len(cidOptional) > = {
			Session.ChannelMessageSend(channelID,
				"Error querying user settings during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		if len(cidOptional) > = {
			Session.ChannelMessageSend(channelID,
				"Error querying user settings during SQL Query: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	defer rows.Close()
	if !rows.Next() {
		// no settings yet; use the defaults
		return
	}
	var maxnsfwTmp, atTmp uint32
	err = rows.Scan(&maxnsfwTmp, &atTmp)
	if err != nil {
		if len(cidOptional) > = {
			Session.ChannelMessageSend(channelID,
				"Error querying user settings during SQL Scan: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	maxnsfw = maxnsfwTmp
	at = atTmp
	return
}

func usInit(uid string) {
	stmt, err := DB.Prepare(`INSERT INTO "UserSettings" ` +
		`("uid", "maxnsfw", "at", "items") ` +
		`VALUES (?, 3, ?, '')`)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error inserting user settings during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(uid, AT_HOME)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error inserting user settings during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	return
}

func usUpdateLocation(uid string, at uint32) {
	stmt, err := DB.Prepare(`UPDATE "UserSettings" SET "at" = ? WHERE "uid" = ?`)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error updating location in user settings during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(at, uid)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error updating location in user settings during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	return
}

func usUpdateMaxNSFW(uid string, maxnsfw uint32) {
	stmt, err := DB.Prepare(`UPDATE "UserSettings" SET "maxnsfw" = ? WHERE "uid" = ?`)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error updating max NSFW in user settings during SQL Prepare: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	_, err = stmt.Exec(maxnsfw, uid)
	if err != nil {
		if len(cidOptional) > 0 {
			Session.ChannelMessageSend(cidOptional,
				"Error updating max NSFW in user settings during SQL Exec: ``" +
				err.Error() + "\u00b4\u00b4.")
		}
		return
	}
	return
}

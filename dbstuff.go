package main

const (
	AT_UNSPECIFIED = uint32(1) << iota,
	AT_HOME_ALONE,
	AT_HOME,
	AT_WORK,
	AT_SCHOOL,
)

const AT_ANYWHERE = uint32(0xffff)

func addPrompt(channelID, author string, dare int, nsfw, at uint32, prompt string) {
	stmt, err := DB.Prepare(`INSERT INTO "Prompts" SET "dare" = ?, "nsfw" = ?, "at" = ?, "prompt" = ?, "blame" = ?`)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	err = DB.Exec(dare, nsfw, at, prompt, author)
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

func givePrompt(channelID, author string, dare int, nsfw, at uint32) {
	stmt, err := DB.Prepare(`SELECT "prompt", "blame" FROM "Prompts" WHERE "dare" = ? AND "nsfw" = ? AND ` +
		`(("at" & ?) <> 0) ORDER BY RANDOM LIMIT 1`)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	rows, err = DB.Query(table, nsfw, at, prompt, author)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Query: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	defer rows.Close()
	if !rows.Next() {
		Session.ChannelMessageSend(channelID, "No prompts matched the criteria.")
		return
	}
	var prompt, blame string
	err := rows.Scan(&prompt, &blame)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Scan: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	Session.ChannelMessageSend(channelID, "<@!" + author + ">, " + prompt + " (added by " + blame + ")")
	return
}

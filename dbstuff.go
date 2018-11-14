package main

const (
	AT_UNSPECIFIED = uint32(1) << iota,
	AT_HOME_ALONE,
	AT_HOME,
	AT_WORK,
	AT_SCHOOL,
)

const AT_ANYWHERE = uint32(0xffff)

func addPrompt(channelID, blame, table string, nsfw, at uint32, prompt string) {
	if at < 1 {
		at = AT_UNSPECIFIED
	}
	stmt, err := DB.Prepare(`INSERT INTO ? SET "nsfw" = ?, "at" = ?, "prompt" = ?, "blame" = ?`)
	if err != nil {
		Session.ChannelMessageSend(channelID,
			"Error adding prompt during SQL Prepare: ``" +
			err.Error() + "\u00b4\u00b4.")
		return
	}
	err = DB.Exec(table, nsfw, at, prompt, blame)
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

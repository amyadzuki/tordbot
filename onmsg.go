package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	message := mc.Message
	if message.Author.Bot {
		return
	}
	payl := message.Content
	if len(payl) < 3 {
		return // attachment-only post, or just a comma
	}
	lowr := strings.ToLower(payl)
	if !strings.HasPrefix(lowr, "tod") {
		return
	}
	payl = payl[3:]
	lowr = lowr[3:]
	for len(payl) > 0 && payl[0] == ' ' {
		payl = payl[1:]
		lowr = lowr[1:]
	}
	channelIDTmp := mc.Message.ChannelID
	channel, err := Session.State.Channel(channelIDTmp)
	if err != nil {
		channel, err = Session.Channel(channelIDTmp)
		if err != nil {
			return
		}
	}
	very := 1
	for strings.HasPrefix(lowr, "very") {
		payl = payl[4:]
		lowr = lowr[4:]
		very++
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	}

	if strings.HasPrefix(lowr, "sfw") {
		payl = payl[3:]
		lowr = lowr[3:]
		very = -very
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	} else if strings.HasPrefix(lowr, "nsfw") {
		payl = payl[4:]
		lowr = lowr[4:]
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	} else {
		very = 0
	}
	nsfw := 1 + very
	nsfwadd := uint32(nsfw)
	if nsfwadd < 0 {
		nsfwadd = 0
	}
	if nsfwadd > 3 {
		nsfwadd = 3
	}
	if channel.NSFW {
		nsfw++
	}
	entropy := PRG.Uint32()
	if (entropy & 0x1) == 0 {
		if entropy & 0x2 == 0 {
			nsfw--
		} else {
			nsfw++
		}
	}
	if nsfw < 0 {
		nsfw = 0
	} else if nsfw > 1 && !channel.NSFW {
		nsfw = 1
	} else if nsfw > 3 {
		nsfw = 3
	}

	u32nsfw := uint32(nsfw)

	at := uint32(0)
	cmdline := strings.Fields(payl)
	for len(cmdline) >= 2 && strings.ToLower(cmdline[0]) == "at" {
		switch strings.ToLower(cmdline[1]) {
		case "anywhere":
			at |= AT_ANYWHERE
		case "homealone":
			at |= AT_HOME | AT_HOME_ALONE
		case "home":
			at |= AT_HOME
			if len(cmdline) >= 3 && strings.ToLower(cmdline[2]) == "alone" {
				at |= AT_HOME_ALONE
			}
		case "work":
			at |= AT_WORK
		case "school":
			at |= AT_SCHOOL
		}
		cmdline = cmdline[2:]
	}
	if at < 1 {
		at = AT_UNSPECIFIED
	}
	if len(cmdline) < 1 {
		return
	}
	author := mc.Message.Author.ID
	switch strings.ToLower(cmdline[0]) {
	case "add":
		if len(cmdline) < 2 {
			Session.ChannelMessageSend(channel.ID,
				"Join our public Google Doc here to suggest stuff:\n" +
				"<https://docs.google.com/document/d/" +
				"1NsD_0fASVaixXJAtWIF4tUVRG9vBiSyyiqM_Sb1Hl2c/edit?usp=sharing>")
		}
		prompt := strings.Join(cmdline[2:], " ")
		switch strings.ToLower(cmdline[1]) {
		case "dare":
			addPrompt(channel.GuildID, channel.ID, author, 1, nsfwadd, at, prompt)
		case "truth":
			addPrompt(channel.GuildID, channel.ID, author, 0, nsfwadd, at, prompt)
		}
	case "dare":
		givePrompt(channel.GuildID, channel.ID, author, 1, nsfwadd, at)
	case "truth":
		givePrompt(channel.GuildID, channel.ID, author, 0, nsfwadd, at)
	case "help":
	case "invite":
		Session.ChannelMessageSend(channel.ID,
			"<https://discordapp.com/oauth2/authorize?client_id=" +
			"512117311" + "415648275&scope=bot&permissions=378" + "944>")
	}
}

func onDare(s *discordgo.Session, mc *discordgo.MessageCreate) {
}

func onTruth(s *discordgo.Session, mc *discordgo.MessageCreate) {
}

func onMessageDelete(s *discordgo.Session, md *discordgo.MessageDelete) {
}

func onMessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
}

func onMessageReactionRemove(s *discordgo.Session, mrr *discordgo.MessageReactionRemove) {
}

func onMessageUpdate(s *discordgo.Session, mu *discordgo.MessageUpdate) {
}

package main

import (
	"strings"
)

func chp(s, prefix string) (bool, string) {
	if len(s) < len(prefix) {
		return false, s
	}
	len_prefix := len(prefix)
	if strings.ToLower(s[0:len_prefix]) == prefix {
		return true, s[len_prefix:]
	}
	return false, s
}

func lex(s *discordgo.Session, m *discordgo.Message) {
	tail := m.Content
	if b, t := chp(tail, "tord"); b {
		tail = t
	} else if b, t := chp(tail, "tod"); false && b {
		tail = t
	} else {
		return
	}

	author := mc.Message.Author.ID
	channelIDTmp := mc.Message.ChannelID
	channel, err := Session.State.Channel(channelIDTmp)
	if err != nil {
		channel, err = Session.Channel(channelIDTmp)
		if err != nil {
			return
		}
	}

	var at, dare uint32
	very, nsfwi := 1, 1

	for {
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		if false {
		} else if b, t := chp(tail, "alone "); b {
			if (at & AT_HOME) != 0 {
				at |= AT_HOMEALONE
			}
		} else if b, t := chp(tail, "anywhere "); b {
			at |= AT_ANYWHERE
		} else if b, t := chp(tail, "at "); b {
			tail = t
		} else if b, t := chp(tail, "home "); b {
			at |= AT_HOME
		} else if b, t := chp(tail, "homealone "); b {
			at |= AT_HOMEALONE
		} else if b, t := chp(tail, "nsfw "); b {
			nsfwi += very
			very = 1
		} else if b, t := chp(tail, "school "); b {
			at |= AT_SCHOOL
		} else if b, t := chp(tail, "sfw "); b {
			nsfwi -= very
			very = 1
		} else if b, t := chp(tail, "very "); b {
			very++
		} else if b, t := chp(tail, "work "); b {
			at |= AT_WORK
		} else {
			break
		}
	}
	if at < 1 {
		at = AT_UNSPECIFIED
	}
	nsfwaddi := nsfwi
	ent := PRG.Uint64()
	if (ent & 1) == 0 {
		if (ent & 2) == 0 {
			nsfwi++
		} else {
			nsfwi--
		}
		ent >>= 1
	}
	ent >>= 1
	if channel.NSFW {
		nsfwi++
	} else if nsfwi > 1 {
		nsfw = 1
	}
	if nsfwi < 0 {
		nsfwi = 0
	} else if nsfwi > 3 {
		nsfwi = 3
	}
	nsfw32 := uint32(nsfw)
	if b, t := chp(tail, "add"); b {
		if nsfwaddi < 0 {
			nsfwaddi = 0
		} else if nsfwaddi > 3 {
			nsfwaddi = 3
		}
		nsfwadd32 := uint32(nsfwaddi)
		return
	}
	switch tail {
	case "truth":
		givePrompt(channel.GuildID, channel.ID, author, 0, nsfw32, at, ent)
	case "dare":
		givePrompt(channel.GuildID, channel.ID, author, 1, nsfw32, at, ent)
	case "either", "go":
		d := uint32(ent) & 1
		ent >>= 1
		givePrompt(channel.GuildID, channel.ID, author, d, nsfw32, at, ent)
	case "fix":
		Session.ChannelMessageSend(channel.ID,
			"fix command coming soon")
	case "help":
		Session.ChannelMessageSend(channel.ID,
			"Prefix is 'tord' and will eventually also respond to 'tod'." +
			"\n" + "Command list:" +
			"\n" + "'truth'          - get a truth prompt" +
			"\n" + "'dare'           - get a dare prompt" +
			"\n" + "'go' or 'either' - get a prompt of a random type" +
			"\n" + "'fix'            - fix common problems automatically" +
			"\n" + "'invite'         - the invite link" +
			"\n" + "'pass' or 'skip' - skip your turn" +
			"")
	case "invite":
		Session.ChannelMessageSend(channel.ID,
			"<https://discordapp.com/oauth2/authorize?client_id=" +
			"512117311" + "415648275&scope=bot&permissions=378" + "944>")
	case "pass", "skip":
		Session.ChannelMessageSend(channel.ID,
			"pass/skip command coming soon")
	default:
		Session.ChannelMessageSend(channel.ID,
			"Unknown command ``" + tail + "\u00b4\u00b4.")
	}
}

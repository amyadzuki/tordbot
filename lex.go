package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
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
	tail := m.Content + " "
	if b, t := chp(tail, "tord"); b {
		tail = t
	} else if b, t := chp(tail, "tod"); false && b {
		tail = t
	} else {
		return
	}

	author := m.Author.ID
	channelIDTmp := m.ChannelID
	channel, err := Session.State.Channel(channelIDTmp)
	if err != nil {
		channel, err = Session.Channel(channelIDTmp)
		if err != nil {
			return
		}
	}

	if b, t := chp(tail, "install"); b {
		tail = t
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		install(channel, tail)
	} else if b, t := chp(tail, "deinstall"); b {
		tail = t
		deinstall(channel)
	} else if b, t := chp(tail, "uninstall"); b {
		tail = t
		deinstall(channel)
	}

	var at uint32
	very, nsfwi := 1, 1

	for {
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		if b, t := chp(tail, "a "); b {
			tail = t
		} else if b, t := chp(tail, "an "); b {
			tail = t
		} else if b, t := chp(tail, "alone "); b {
			tail = t
			if (at & AT_HOME) != 0 {
				at |= AT_HOMEALONE
			}
		} else if b, t := chp(tail, "anywhere "); b {
			tail = t
			at |= AT_ANYWHERE
		} else if b, t := chp(tail, "at "); b {
			tail = t
		} else if b, t := chp(tail, "home "); b {
			tail = t
			at |= AT_HOME
		} else if b, t := chp(tail, "homealone "); b {
			tail = t
			at |= AT_HOMEALONE
		} else if b, t := chp(tail, "i'm "); b {
			tail = t
		} else if b, t := chp(tail, "im "); b {
			tail = t
		} else if b, t := chp(tail, "in "); b {
			tail = t
		} else if b, t := chp(tail, "nsfw "); b {
			tail = t
			nsfwi += very
			very = 1
		} else if b, t := chp(tail, "school "); b {
			tail = t
			at |= AT_SCHOOL
		} else if b, t := chp(tail, "sfw "); b {
			tail = t
			nsfwi -= very
			very = 1
		} else if b, t := chp(tail, "very "); b {
			tail = t
			very++
		} else if b, t := chp(tail, "work "); b {
			tail = t
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
		nsfwi = 1
	}
	if nsfwi < 0 {
		nsfwi = 0
	} else if nsfwi > 3 {
		nsfwi = 3
	}
	nsfw32 := uint32(nsfwi)
	if b, t := chp(tail, "add"); b {
		tail = t
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		if nsfwaddi < 0 {
			nsfwaddi = 0
		} else if nsfwaddi > 3 {
			nsfwaddi = 3
		}
		nsfwadd32 := uint32(nsfwaddi)
		var dare int
		if b, t := chp(tail, "truth "); b {
			tail = t
			dare = 0
		} else if b, t := chp(tail, "dare "); b {
			tail = t
			dare = 1
		} else {
			Session.ChannelMessageSend(channel.ID,
				"Unknown add sub-command ``" + tail + "\u00b4\u00b4.")
			return
		}
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		addPrompt(channel.GuildID, channel.ID, author, dare, nsfwadd32, at, tail)
		return
	}
	switch tail {
	case " ":
		Session.ChannelMessageSend(channel.ID,
			"empty command to update settings coming soon")
	case "truth ":
		givePrompt(channel.GuildID, channel.ID, author, 0, nsfw32, at, ent)
	case "dare ":
		givePrompt(channel.GuildID, channel.ID, author, 1, nsfw32, at, ent)
	case "either ", "go ":
		d := int(uint32(ent) & 1)
		ent >>= 1
		givePrompt(channel.GuildID, channel.ID, author, d, nsfw32, at, ent)
	//
	case "fix ":
		Session.ChannelMessageSend(channel.ID,
			"fix command coming soon")
	case "help ":
		Session.ChannelMessageSend(channel.ID,
			"Prefix is 'tord' and will eventually also respond to 'tod'." +
			"\n" +
			"\n" + "Installation:" +
			"\n" + "```prolog" +
			"\n" + "'install #voice-channel'   - install to the current channel" +
			"\n" + "'deinstall' or 'uninstall' - deinstall from the current channel" +
			"\n" + "```" +
			"\n" + "Pre-command modifiers:" +
			"\n" + "```prolog" +
			"\n" + "'car'            - mark yourself as a passenger in a vehicle" +
			"\n" + "'home'           - mark yourself at home" +
			"\n" + "'home alone'     - mark yourself home alone" +
			"\n" + "'nsfw'           - request a more NSFW prompt" +
			"\n" + "'school'         - mark yourself at school" +
			"\n" + "'sfw'            - request a less NSFW prompt" +
			"\n" + "'very nsfw'      - request a much more NSFW prompt" +
			"\n" + "'very sfw'       - request a much less NSFW prompt" +
			"\n" + "'work'           - mark yourself at work" +
			"\n" + "```" +
			"\n" + "Command list:" +
			"\n" + "```prolog" +
			"\n" + "'truth'          - get a truth prompt" +
			"\n" + "'dare'           - get a dare prompt" +
			"\n" + "'go' or 'either' - get a prompt of a random type" +
			"\n" +
			"\n" + "'fix'            - fix common problems automatically" +
			"\n" + "'invite'         - the invite link" +
			"\n" + "'pass' or 'skip' - skip your turn" +
			"\n" + "'score'          - check your score" +
			"\n" + "'scores'         - view the score of everyone playing" +
			"\n" + "'suggest'        - get a link to the suggestion doc" +
			"\n" + "'turns'          - view the turn order" +
			"\n" + "```" +
			"")
	case "invite ":
		Session.ChannelMessageSend(channel.ID,
			"<https://discordapp.com/oauth2/authorize?client_id=" +
			"512117311" + "415648275&scope=bot&permissions=378" + "944>")
	case "pass ", "skip ":
		Session.ChannelMessageSend(channel.ID,
			"pass/skip command coming soon")
	case "score ":
		Session.ChannelMessageSend(channel.ID,
			"score command coming soon")
	case "scores ":
		Session.ChannelMessageSend(channel.ID,
			"scores command coming soon")
	case "suggest ":
		Session.ChannelMessageSend(channel.ID,
			"Join our public Google Doc here to suggest stuff:\n" +
			"<https://docs.google.com/document/d/" +
			"1NsD_0fASVaixXJAtWIF4tUVRG9vBiSyyiqM_Sb1Hl2c/edit?usp=sharing>")
	case "turns ":
		Session.ChannelMessageSend(channel.ID,
			"turns command coming soon")
	default:
		Session.ChannelMessageSend(channel.ID,
			"Unknown command ``" + tail + "\u00b4\u00b4.")
	}
}

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

	var at, dare uint32
	very := 1
	var sfw, nsfw bool

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
			nsfw = true
		} else if b, t := chp(tail, "school "); b {
			at |= AT_SCHOOL
		} else if b, t := chp(tail, "sfw "); b {
			sfw = true
		} else if b, t := chp(tail, "very "); b {
			very++
		} else if b, t := chp(tail, "work "); b {
			at |= AT_WORK
		} else {
			break
		}
	}
	ent := PRG.Uint64()
	switch tail {
	case "truth":
		givePrompt(channel.GuildID, channel.ID, author, 0, u32nsfw, at, ent)
	case "dare":
		givePrompt(channel.GuildID, channel.ID, author, 1, u32nsfw, at, ent)
	case "go":
		d := uint32(ent) & 1
		ent >>= 1
		givePrompt(channel.GuildID, channel.ID, author, d, u32nsfw, at, ent)
	case "fix":
	case "help":
	case "invite":
	case "pass", "skip":
	}
}

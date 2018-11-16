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

	for {
		for len(tail) > 0 && tail[0] == ' ' {
			tail = tail[1:]
		}
		if b, t := chp(tail, "at "); b {
			tail = t
		} else if b, t := chp(tail, "anywhere "); b {
			at |= AT_ANYWHERE
		}
	}
}

package main

import (
	"strings"
)

func caseHasLcPrefixSkip(s, prefix string) (bool, string) {
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
}

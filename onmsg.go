package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	message := mc.Message
	lex(s, message)
}

func onMessageDelete(s *discordgo.Session, md *discordgo.MessageDelete) {
}

func onMessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
}

func onMessageReactionRemove(s *discordgo.Session, mrr *discordgo.MessageReactionRemove) {
}

func onMessageUpdate(s *discordgo.Session, mu *discordgo.MessageUpdate) {
}

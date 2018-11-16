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
		default:
			Session.ChannelMessageSend(channel.ID,
				"Unknown add command ``" + cmdline[1] + "\u00b4\u00b4.")
		}
}

func onMessageDelete(s *discordgo.Session, md *discordgo.MessageDelete) {
}

func onMessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
}

func onMessageReactionRemove(s *discordgo.Session, mrr *discordgo.MessageReactionRemove) {
}

func onMessageUpdate(s *discordgo.Session, mu *discordgo.MessageUpdate) {
}

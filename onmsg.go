package bot

import (
	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	message := mc.Message
	if message.Author.Bot {
		return
	}
	payl := message.Content
	if len(payl) < 2 {
		return // attachment-only post, or just a comma
	}
	if payl[0] != ',' {
		return
	}
	com, err := NewCom(s, message, payl[1:])
	if err != nil {
		return
	}
	com.Lex()
}

func onMessageDelete(s *discordgo.Session, md *discordgo.MessageDelete) {
}

func onMessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
}

func onMessageReactionRemove(s *discordgo.Session, mrr *discordgo.MessageReactionRemove) {
}

func onMessageUpdate(s *discordgo.Session, mu *discordgo.MessageUpdate) {
}

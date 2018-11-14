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
	if len(payl) < 4 {
		return // attachment-only post, or just a comma
	}
	payl = strings.ToLower(payl)
	if !strings.HasPrefix(payl, "tod") {
		return
	}
	payl = payl[4:]
	for len(payl) > 0 && payl[0] == ' ' {
		payl = payl[1:]
	}
	channelID := mc.Message.ChannelID
	channel, err := session.State.Channel(channelID)
	if err != nil {
		channel, err = session.Channel(channelID)
		if err != nil {
			return
		}
	}
	//
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

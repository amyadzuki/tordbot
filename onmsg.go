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
	if len(payl) < 3 {
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
	nsfw := 1
	if channel.NSFW {
		nsfw = 5
	}
	very := false
	if strings.HasPrefix(payl, "very") {
		payl = payl[4:]
		very = true
	}
	if strings.HasPrefix(payl, "sfw") {
		payl = payl[3:]
		if very {
			nsfw /= 2 // 0 or 2
		} // else 1 or 5
	} else if strings.HasPrefix(payl, "nsfw") {
		payl = payl[4:]
		if very {
			nsfw += 2 // 3 or 7
		} else {
			nsfw += 4 // 5 or 9
		}
	}
	if nsfw < 0 || nsfw > 9 {
		panic("nsfw < 0 || nsfw > 9")
	}

	cmdline := strings.Split(payl)
	if len(cmdline) < 1 {
		return
	}
	switch cmdline[0] {
	case "dare":
	case "help":
	case "truth":
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

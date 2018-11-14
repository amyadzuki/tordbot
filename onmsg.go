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
	lowr = strings.ToLower(payl)
	if !strings.HasPrefix(lowr, "tod") {
		return
	}
	payl = payl[3:]
	lowr = lowr[3:]
	for len(payl) > 0 && payl[0] == ' ' {
		payl = payl[1:]
		lowr = lowr[1:]
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
	if strings.HasPrefix(lowr, "very") {
		payl = payl[4:]
		lowr = lowr[4:]
		very = true
	}
	if strings.HasPrefix(lowr, "sfw") {
		payl = payl[3:]
		lowr = lowr[3:]
		if very {
			nsfw /= 2 // 0 or 2
		} // else 1 or 5
	} else if strings.HasPrefix(lowr, "nsfw") {
		payl = payl[4:]
		lowr = lowr[4:]
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
	embed := new(discordgo.Embed)
	switch cmdline[0] {
	case "add":
		if len(cmdline) < 2 {
			Session.ChannelMessageSend(channelID,
				"Join our public Google Doc here to suggest stuff:\n" +
				"https://docs.google.com/document/d/" +
				"1NsD_0fASVaixXJAtWIF4tUVRG9vBiSyyiqM_Sb1Hl2c/edit?usp=sharing")
		}
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

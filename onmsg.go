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
	very := 0
	for strings.HasPrefix(lowr, "very") {
		payl = payl[4:]
		lowr = lowr[4:]
		very++
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	}

	if strings.HasPrefix(lowr, "sfw") {
		payl = payl[3:]
		lowr = lowr[3:]
		very = -very
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	} else if strings.HasPrefix(lowr, "nsfw") {
		payl = payl[4:]
		lowr = lowr[4:]
		for len(payl) > 0 && payl[0] == ' ' {
			payl = payl[1:]
			lowr = lowr[1:]
		}
	} else {
		very = 0
	}
	nsfw := 5 + 2 * very
	if nsfw < 0 {
		nsfw = 0
	}
	nsfwadd := uint32(nsfw)
	if channel.NSFW {
		nsfw += nsfw
	}
	u32nsfw := uint32(nsfw)
	u32nsfw = prg.Uint32() % u32nsfw // first step
	u32nsfw = (prg.Uint32() % u32nsfw + u32nsfw) / 2 // second step
	if u32nsfw > 5 && !channel.NSFW {
		u32nsfw = 5
	}

	cmdline := strings.Split(payl)
	if len(cmdline) < 1 {
		return
	}
	switch strings.ToLower(cmdline[0]) {
	case "add":
		if len(cmdline) < 2 {
			Session.ChannelMessageSend(channelID,
				"Join our public Google Doc here to suggest stuff:\n" +
				"https://docs.google.com/document/d/" +
				"1NsD_0fASVaixXJAtWIF4tUVRG9vBiSyyiqM_Sb1Hl2c/edit?usp=sharing")
		}
		prompt := strings.Join(cmdline[2:], " ")
		switch strings.ToLower(cmdline[1]) {
		case "dare":
			addPrompt("Dares", nsfwadd, prompt)
		case "truth":
			addPrompt("Truths", nsfwadd, prompt)
		}
	case "dare":
	case "help":
	case "invite":
		Session.ChannelMessageSend(channelID,
			"https://discordapp.com/oauth2/authorize?client_id=" +
			"512117311" + "415648275&scope=bot&permissions=378" + "944")
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

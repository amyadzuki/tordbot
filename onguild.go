package main

import (
	"github.com/bwmarrin/discordgo"
)

func onVoiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	voiceStateUpdate(vsu.VoiceState)
}

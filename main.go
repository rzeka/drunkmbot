package drunkmbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
			)

type Bot struct {
	Discord *discordgo.Session
	CommandPrefix string
}

func New(token string) Bot {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalln("Error while creating Discord session ", err)
		os.Exit(1)
	}

	return Bot{
		Discord: discord,
		CommandPrefix: "!",
	}
}

func (bot *Bot) Start() {
	bot.initCommands()
	bot.Discord.Open()
}

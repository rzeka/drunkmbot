package drunkmbot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandCallable func(args string, s *discordgo.Session, m *discordgo.MessageCreate)

type command struct {
	trigger string
	callable CommandCallable
}
var commands []command

func (bot *Bot) AddCommand(trigger string, callable CommandCallable) {
	trigger = bot.CommandPrefix + trigger

	commands = append(
		commands,
		command{
			trigger,
			callable,
		},
	)
}

func (bot *Bot) initCommands() {
	var messageHandler = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		for _, command := range commands {
			if strings.HasPrefix(m.Content, command.trigger) {
				command.callable(m.Content[len(command.trigger) + 1:], s, m)
			}
		}
	}

	bot.Discord.AddHandler(messageHandler)
}
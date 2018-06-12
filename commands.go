package drunkmbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type CommandCallable func(args string, s *discordgo.Session, m *discordgo.MessageCreate)

type Command struct {
	Trigger     string
	Callable    CommandCallable
	Permissions int
}

var commands []*Command

func (bot *Bot) AddCommand(trigger string, permissions int, callable CommandCallable) {
	trigger = bot.CommandPrefix + trigger

	commands = append(
		commands,
		&Command{
			trigger,
			callable,
			permissions,
		},
	)
}

func (bot *Bot) initCommands() {
	bot.Discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		for _, command := range commands {
			if isCommand(command.Trigger, m.Content) {
				runCommand(command, s, m)
			}
		}
	})
}

func isCommand(trigger string, message string) bool {
	return message == trigger || strings.HasPrefix(message, trigger+" ")
}

func runCommand(command *Command, s *discordgo.Session, m *discordgo.MessageCreate) {
	if command.Permissions != 0 {
		userPermissions, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			log.Fatalln("Could no execute admin command ", err)
			return
		}

		if userPermissions&command.Permissions != command.Permissions {
			return
		}
	}

	command.Callable(
		extractCommandArgs(command.Trigger, m.Content),
		s,
		m,
	)
}

func extractCommandArgs(trigger string, message string) string {
	return strings.Trim(message[len(trigger):], " ")
}

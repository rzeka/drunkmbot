package main

import (
	"github.com/rzeka/drunkmbot"
)

func LoadPlugin(bot *drunkmbot.Bot) {
	bot.AddCommand("rank", commandRank)
}

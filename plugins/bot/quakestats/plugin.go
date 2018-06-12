package main

import (
	"github.com/rzeka/drunkmbot"
)

func LoadPlugin(bot *drunkmbot.Bot) {
	bot.AddCommand("rank", 0, commandRank)
	bot.AddCommand("ranks", 0, commandRanks)
}

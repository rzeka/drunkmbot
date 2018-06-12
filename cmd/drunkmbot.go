package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rzeka/drunkmbot"
	"log"
)

var (
	token string
	commandPrefix string
)

func init() {
	godotenv.Load(".env") //don't care about errors here

	token = os.Getenv("BOT_TOKEN")

	prefix := os.Getenv("COMMAND_PREFIX")
	if prefix == "" {
		prefix = "!"
	}

	commandPrefix = prefix
}

func main() {
	discord := drunkmbot.New(token)
	discord.CommandPrefix = commandPrefix

	discord.LoadPlugins("plugins/bot")
	discord.Start()

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Discord.Close()
}

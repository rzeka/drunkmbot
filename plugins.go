package drunkmbot

import (
	"path/filepath"
	"log"
	"os"
	"plugin"
	"strings"
)

func (bot *Bot) LoadPlugins(path string) {
	path = strings.TrimRight(path, "/\\")
	files, err := filepath.Glob(path + "/**/*.so")

	if err != nil {
		log.Fatalln("Error while looking for plugins", err)
		os.Exit(1)
	}

	for _,file := range files {
		p, err := plugin.Open(file)

		if err != nil {
			log.Fatalln("Could not load plugin ", file)
			panic(err)
		}

		loadPluginSymbol, err := p.Lookup("LoadPlugin")
		if err != nil {
			log.Fatalln("Could not load plugin ", file)
			panic(err)
		}

		loadPluginSymbol.(func(bot *Bot))(bot)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	)

type Response struct {
	PlayerRatings map[string]PlayerRating `json:"playerRatings"`
}

type PlayerRating struct {
	Rating     int `json:"rating"`
	Deviation  int `json:"deviation"`
	LastChange int `json:"lastChange"`
}

type Rank struct {
	Name   string
	Rating int
}

var Ranks = []Rank{
	{"Bronze 1", 0},
	{"Bronze 2", 675},
	{"Bronze 3", 750},
	{"Bronze 4", 825},
	{"Bronze 5", 900},

	{"Silver 1", 975},
	{"Silver 2", 1050},
	{"Silver 3", 1125},
	{"Silver 4", 1200},
	{"Silver 5", 1275},

	{"Gold 1", 1350},
	{"Gold 2", 1425},
	{"Gold 3", 1500},
	{"Gold 4", 1575},
	{"Gold 5", 1650},

	{"Diamond 1", 1725},
	{"Diamond 2", 1800},
	{"Diamond 3", 1870},
	{"Diamond 4", 1950},
	{"Diamond 5", 2025},

	{"Elite", 2100},
}

func commandRank(args string, s *discordgo.Session, m *discordgo.MessageCreate) {
	name := strings.Trim(args, " ")

	if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
		name = name[1 : len(name)-1]
	}

	ratings, err := apiGetPlayerRatings(name)
	var message string
	if err != nil {
		message = "Hmm... Something went wrong :) Check name and retry"
	} else if len(ratings.PlayerRatings) == 0 {
		message = "No rank info for " + name
	} else {
		message = name + ""
		lastChangeSign := ""

		for name, rating := range ratings.PlayerRatings {
			if rating.LastChange > 0 {
				lastChangeSign = "+"
			}

			message += fmt.Sprintf(
				" # **%v** @ %v *[%v±%v %v%v]*",
				getRankNameFromRating(rating.Rating),
				name,
				rating.Rating,
				rating.Deviation,
				lastChangeSign,
				rating.LastChange,
			)
		}
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

func commandRanks(_ string, s *discordgo.Session, m *discordgo.MessageCreate) {
	message := ""

	for i, rank := range Ranks {
		if i%5 == 0 {
			message = message + "\n"
		}

		message = message + fmt.Sprintf("**%v**: %d\n", rank.Name, rank.Rating)
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

func apiGetPlayerRatings(name string) (Response, error) {
	var response = Response{}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://stats.quake.com/api/v2/Player/Stats?name="+name,
		nil,
	)
	if err != nil {
		return response, err
	}

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func getRankNameFromRating(rating int) string {
	var rankName = Ranks[0].Name

	for _, rank := range Ranks {
		if rating < rank.Rating {
			break
		}

		rankName = rank.Name
	}

	return rankName
}

package main

import (
	"fmt"
	"strings"

	"github.com/PioKozi/xkcdbot-discord/cmd/gosearch"

	. "github.com/PioKozi/xkcdbot-discord/pkg/common"

	"github.com/bwmarrin/discordgo"
)

// maps input to id of relevant xkcd
// some are relevant more often, so are added here
var Cache = map[string]string{
	"security": "538",
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, BotPrefix) {
		message := m.Content
		channel := m.ChannelID
		LogMessage(message)
		var response string

		/* COMMANDS USING IDS MUST BE FIRST */

		// Searching xkcd comics
		if strings.HasPrefix(message, BotPrefix+"xkcdid") { // immediately send by id
			message = CleanInput(message, BotPrefix+"xkcdid")
			if ValidID(message) {
				response = fmt.Sprintf("https://xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, BotPrefix+"xkcd") { // search by name via gosearc
			message = CleanInput(message, BotPrefix+"xkcd")
			if cached, exists := Cache[message]; exists { // check cache first, may save time
				response = fmt.Sprintf("https://xkcd.com/%s/", cached)
			} else { // wasn't in cache, search with GoogleScrape
				searchTerm := fmt.Sprintf("site:xkcd.com AND inurl:https://xkcd.com/ %s", message)
				result, err := gosearch.GoogleScrape(searchTerm)
				if err != nil {
					response = fmt.Sprintf("There was an error searching for results: %s", err)
				}
				if result == (gosearch.GoogleResult{}) { // check if there are no results
					response = "no good results for that search"
				} else {
					response = result.Url
				}
			}
		}

		// Searching what if
		if strings.HasPrefix(message, BotPrefix+"whatifid") { // immediately send by id
			message = CleanInput(message, BotPrefix+"whatifid")
			if ValidID(message) {
				response = fmt.Sprintf("https://what-if.xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, BotPrefix+"whatif") { // search by name via gosearch - no cache for whatif
			message = CleanInput(message, BotPrefix+"whatif")
			searchTerm := fmt.Sprintf("site:what-if.xkcd.com AND inurl:https://what-if.xkcd.com/ %s", message)
			result, err := gosearch.GoogleScrape(searchTerm)
			if err != nil {
				response = fmt.Sprintf("There was an error searching for results: %s", err)
			}
			if result == (gosearch.GoogleResult{}) {
				response = "no good results for search"
			} else {
				response = result.Url
			}
		}

		s.ChannelMessageSend(channel, response)
	}
}

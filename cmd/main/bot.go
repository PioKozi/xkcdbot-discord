package main

import (
	"fmt"
	"strings"

	"github.com/PioKozi/xkcdbot-discord/cmd/cache"
	"github.com/PioKozi/xkcdbot-discord/cmd/search"

	. "github.com/PioKozi/xkcdbot-discord/pkg/common"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		message := m.Content
		channel := m.ChannelID
		LogMessage(message)
		var response string

		/* COMMANDS USING IDS MUST BE FIRST */

		// Searching xkcd comics
		if strings.HasPrefix(message, botPrefix+"xkcdid") { // immediately send by id
			message = PrepareInput(message, botPrefix+"xkcdid")
			if ValidID(message) {
				response = fmt.Sprintf("https://xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, botPrefix+"xkcd") { // search by name via package search
			message = PrepareInput(message, botPrefix+"xkcd")
			if cached, exists := cache.Cache[message]; exists { // check cache first, may save time
				response = cached
			} else { // wasn't in cache, search with GoogleScrape
				searchTerm := fmt.Sprintf("site:xkcd.com AND inurl:https://xkcd.com/ %s", message)
				result, err := search.GoogleScrape(searchTerm)
				if err != nil {
					response = fmt.Sprintf("There was an error searching for results: %s", err)
				}
				if result == "" { // check if there are no results
					response = "no good results for that search"
				} else {
					response = result
					cache.UpdateLastSearches(message, result) // updates list of previous searches
				}
			}
		}

		// Searching what if
		if strings.HasPrefix(message, botPrefix+"whatifid") { // immediately send by id
			message = PrepareInput(message, botPrefix+"whatifid")
			if ValidID(message) {
				response = fmt.Sprintf("https://what-if.xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, botPrefix+"whatif") { // search by name via package search - no cache for whatif
			message = PrepareInput(message, botPrefix+"whatif")
			searchTerm := fmt.Sprintf("site:what-if.xkcd.com AND inurl:https://what-if.xkcd.com/ %s", message)
			result, err := search.GoogleScrape(searchTerm)
			if err != nil {
				response = fmt.Sprintf("There was an error searching for results: %s", err)
			}
			if result == "" {
				response = "no good results for search"
			} else {
				response = result
			}
		}

		s.ChannelMessageSend(channel, response)

		cache.UpdateCache()
	}
}

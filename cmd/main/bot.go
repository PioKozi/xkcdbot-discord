package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PioKozi/xkcdbot-discord/cmd/gosearch"

	. "github.com/PioKozi/xkcdbot-discord/pkg/common"

	"github.com/bwmarrin/discordgo"
)

// maps input to id of relevant xkcd
// some are relevant more often, so are added here
// TODO: make gocache package
// TODO: move this to gocache package
var cache = map[string]string{
	"security": "538",
}
var lastSearches = make(map[string]string)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		message := m.Content
		message = strings.ToLower(message)
		channel := m.ChannelID
		LogMessage(message)
		var response string

		/* COMMANDS USING IDS MUST BE FIRST */

		// Searching xkcd comics
		if strings.HasPrefix(message, botPrefix+"xkcdid") { // immediately send by id
			message = CleanInput(message, botPrefix+"xkcdid")
			if ValidID(message) {
				response = fmt.Sprintf("https://xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, botPrefix+"xkcd") { // search by name via gosearch
			message = CleanInput(message, botPrefix+"xkcd")
			if cached, exists := cache[message]; exists { // check cache first, may save time
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
					id := strings.Split(response, "/")[3] // following https format, this should always be the id
					lastSearches[message] = id            // add id to lastSearches
					log.Print(lastSearches)
				}
			}
		}

		// Searching what if
		if strings.HasPrefix(message, botPrefix+"whatifid") { // immediately send by id
			message = CleanInput(message, botPrefix+"whatifid")
			if ValidID(message) {
				response = fmt.Sprintf("https://what-if.xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, botPrefix+"whatif") { // search by name via gosearch - no cache for whatif
			message = CleanInput(message, botPrefix+"whatif")
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

		// adding to cache
		if len(lastSearches) == 5 { // every 5 .xkcd commands
			UpdateStringsMap(cache, lastSearches)
			log.Print(cache)
		}
	}
}

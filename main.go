package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/PioKozi/xkcdbot-discord/gosearch"

	"github.com/bwmarrin/discordgo"
)

var (
	// values for used for bot
	Token     = os.Getenv("xkcdbottoken")
	BotPrefix = "."

	// maps input to id of relevant xkcd
	// some are relevant more often, so are added here
	Cache = map[string]string{
		"security": "538",
	}
)

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, BotPrefix) {
		message := m.Content
		channel := m.ChannelID
		logMessage(message)
		var response string

		/* COMMANDS USING IDS MUST BE FIRST */

		// Searching xkcd comics
		if strings.HasPrefix(message, BotPrefix+"xkcdid") {
			message = cleanInput(message, BotPrefix+"xkcdid")
			if validID(message) {
				response = fmt.Sprintf("https://xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, BotPrefix+"xkcd") {
			message = cleanInput(message, BotPrefix+"xkcd")
			if cached, exists := Cache[message]; exists {
				response = fmt.Sprintf("https://xkcd.com/%s/", cached)
			} else {
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
		if strings.HasPrefix(message, BotPrefix+"whatifid") {
			message = cleanInput(message, BotPrefix+"whatifid")
			if validID(message) {
				response = fmt.Sprintf("https://what-if.xkcd.com/%s/", message)
			} else {
				response = "not a valid ID"
			}
		} else if strings.HasPrefix(message, BotPrefix+"whatif") {
			message = cleanInput(message, BotPrefix+"whatif")
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

// useful is it's going pretty bad or is malicious around some time, somehow
func logMessage(s string) {

	fmt.Println("------------------------")
	fmt.Println(time.Now())
	fmt.Println(s)
}

func cleanInput(message, prefix string) string {

	message = strings.TrimPrefix(message, prefix)
	message = strings.TrimSpace(message)
	return message
}

func validID(id string) bool {

	i, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	if i < 0 {
		return false
	}
	return true
}

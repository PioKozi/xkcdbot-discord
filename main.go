package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func flagsParse() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	flagsParse()

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

	if strings.HasPrefix(m.Content, ".") {
		message := m.Content
		channel := m.ChannelID
		logMessage(message)

		/* COMMANDS USING IDS MUST BE FIRST */

		// Searching xkcd comics
		if strings.HasPrefix(message, ".xkcdid") {
			message = cleanInput(message, ".xkcdid")
			if validID(message) {
				link := fmt.Sprintf("https://xkcd.com/%s", message)
				s.ChannelMessageSend(channel, link)
			} else {
				s.ChannelMessageSend(channel, "not an ID")
			}
		} else if strings.HasPrefix(message, ".xkcd") {
			message = cleanInput(message, ".xkcd")
			searchTerm := fmt.Sprintf("site:xkcd.com AND inurl:https://xkcd.com/ %s", message)
			result, err := GoogleScrape(searchTerm)
			if err != nil {
				fmt.Println("ERROR: ", err)
				return
			}
			if result == (GoogleResult{}) { // check if there are no results
				s.ChannelMessageSend(channel, "no good results for search")
			} else {
				s.ChannelMessageSend(channel, result.Url)
			}
		}

		// Searching what if
		if strings.HasPrefix(message, ".whatifid") {
			message = cleanInput(message, ".whatifid")
			if validID(message) {
				link := fmt.Sprintf("https://what-if.xkcd.com/%s", message)
				s.ChannelMessageSend(channel, link)
			} else {
				s.ChannelMessageSend(channel, "not an ID")
			}
		} else if strings.HasPrefix(message, ".whatif") {
			message = cleanInput(message, ".whatif")
			searchTerm := fmt.Sprintf("site:what-if.xkcd.com AND inurl:https://what-if.xkcd.com/ %s", message)
			result, err := GoogleScrape(searchTerm)
			if err != nil {
				fmt.Println("ERROR: ", err)
				return
			}
			if result == (GoogleResult{}) {
				s.ChannelMessageSend(channel, "no good results for search")
			} else {
				s.ChannelMessageSend(channel, result.Url)
			}
		}
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

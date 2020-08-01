package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	// values for used for bot
	token     = os.Getenv("XKCDBOTTOKEN")
	botPrefix = "."
)

func main() {

	// initialise the bot
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer bot.Close() // defer closing the session

	// add messageCreate() as a response to MessageCreate events
	bot.AddHandler(messageCreate)

	// able to receive message events
	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// open the connection
	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// wait for signal to kill bot
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

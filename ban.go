package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "OTc1OTM5NzIxNjY4NzM5MTgy.G3m6i9.Sg4lOEi2TUnTGnQ_aWdOcBNGNCUqCm4ylQYx-g")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.

	// dg.AddHandler(test)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("Bot is up!") })
	dg.AddHandler(messageCreate)
	dg.AddHandler(testing)
	discordgo.NewState()
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func testing(client *discordgo.Session, user *discordgo.GuildMembersChunk) {
	/*
		log.Println("Tracke!")
		for i := 0; i < len(user.Presences); i++ {
			for x := 0; x < len(user.Presences[i].Activities); x++ {
				fmt.Println(user.Presences[i].Activities[x].Name)
				if user.Presences[i].Activities[x].Name == "Visual Studio Code" {
					fmt.Println("BANNED")
					fmt.Println(user.Presences[i].User.ID)
					fmt.Println(user.GuildID)
					client.GuildBanCreate(user.GuildID, user.Presences[i].User.ID, 0)
				}
			}

		}
	*/
	playing := 0
	for tick := range time.Tick(15 * time.Minute) {
		fmt.Println(tick)
		log.Println("Tracke!")
		for i := 0; i < len(user.Presences); i++ {
			for x := 0; x < len(user.Presences[i].Activities); x++ {
				fmt.Println(user.Presences[i].Activities[x].Name)
				if user.Presences[i].Activities[x].Name == "Visual Studio Code" {
					if playing == 1 {
						fmt.Println("BANNED")
						// client.GuildBanCreate("919834369902915584", user.Presences[i].User.ID, 999)
						client.GuildBanCreate(user.GuildID, user.Presences[i].User.ID, 0)
					}

					playing = 1
				} else {
					playing = 0
				}
			}

		}
	}

}

// func test()
// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	test := s.RequestGuildMembers("919834369902915584", "", 0, "5325", true)
	fmt.Println(test)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

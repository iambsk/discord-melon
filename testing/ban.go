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
	"github.com/emirpasic/gods/maps/hashmap"
)

// Variables used for command line parameters
var (
	Token string
)
var dg *discordgo.Session

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	dg, _ = discordgo.New("Bot " + "OTc1OTM5NzIxNjY4NzM5MTgy.G3m6i9.Sg4lOEi2TUnTGnQ_aWdOcBNGNCUqCm4ylQYx-g")

}

func main() {

	// Create a new Discord session using the provided bot token.

	// Register the messageCreate func as a callback for MessageCreate events.

	// dg.AddHandler(test)
	// dg.Type()

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("Bot is up!") })
	dg.AddHandlerOnce(messageCreate)

	discordgo.NewState()
	// Open a websocket connection to Discord and begin listening.
	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go ban()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func ban() {

	fmt.Println("Ban function up and running")
	for tick := range time.Tick(15 * time.Minute) {
		fmt.Println(tick)
		m := hashmap.New()
		mem, _ := dg.GuildMembers("919834369902915584", "", 1000)
		for i := 0; i < len(mem); i++ {
			user, err2 := dg.State.Presence("919834369902915584", mem[i].User.ID)
			if err2 == nil {
				for x := 0; x < len(user.Activities); x++ {
					if user.Activities[x].Name == "Fortnite" {
						_, boo := m.Get(mem[i].User.ID)
						if boo {
							fmt.Println("BANNED")
							// client.GuildBanCreate("919834369902915584", user.Presences[i].User.ID, 999)
							// dg.GuildBanCreate("919834369902915584", mem[i].User.ID, 0)
						}
						m.Put(mem[i].User.ID, 1)
					} else {
						m.Remove(mem[i].User.ID)
					}

				}
			}
			err2 = nil
		}
	}
}

// func test()
// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	s.ChannelMessageSend(m.ChannelID, "I have been activated. Fear me.")
}

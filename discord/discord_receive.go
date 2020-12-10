package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/mvc"
)

var session **discordgo.Session

func init() {
	config.Discord_Session, _ = discordgo.New()
	session = &config.Discord_Session
	//s := config.Discord_Session
	//fmt.Printf("%v %p %v\n", &s, s, *s)
	//s = session
	//fmt.Printf("%v %p %v\n", &s, s, *s)

	//os.Exit(1)
	// Discord Authentication Token
}

func Main() {
	fmt.Println("discord package")

	(*session).Token = db.DiscordToken
	err := (*session).Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	(*session).AddHandler(messageHandler)

	// In this example, we only care about receiving message events.
	(*session).Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

}
func Close() {
	log.Println("Discord Shutting Down")
	(*session).Close()
	log.Println("Discord Shut Down")
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { // ignore itself //ignore all bots not just itself
		return
	}
	err := mvc.DiscordReceive(m.Author.Username, m.Content, m.ChannelID, m.Attachments)
	if err != nil {
		fmt.Println(err)
	}

}

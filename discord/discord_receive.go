package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
)

var Session, _ = discordgo.New()

func init() {

	// Discord Authentication Token
	Session.Token = config.DiscordToken
}

func Main() {
	fmt.Println("discord package")
	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	Session.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	log.Println(Session.Token)

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID { // ignore itself
		return
	}
	fmt.Println(m.ChannelID)
	_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Username+" sent this message")
	if err != nil {
		fmt.Println(err)
	}

}

func IssueWebhook(ci string) (string, string, error) {

	wh, err := Session.WebhookCreate(ci, "Groupme Sync", "")
	if err != nil {
		return "", "", err
	}
	return wh.ID, wh.Token, nil
}

func CallWebhook(id, token string, data *discordgo.WebhookParams) {

	Session.WebhookExecute(id, token, false, data)
}

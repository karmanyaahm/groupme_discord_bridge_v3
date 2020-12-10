package discord_utils

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
)

var session **discordgo.Session

func init() {
	//fmt.Println('a')
	//	s := session
	//	fmt.Printf("%v %p %v\n", &s, s, s)

	session = &config.Discord_Session
	//	s = session
	//	fmt.Printf("%v %p %v\n", &config.Discord_Session, config.Discord_Session, *config.Discord_Session)

}

func GetChannelName(ci string) (string, error) {
	log.Println("aaaa")
	c, e := (*session).Channel(ci)

	if e != nil {
		log.Println(e)
		return "", errors.New("Not Found")
	}

	return c.Name, nil

}

func CallWebhook(id, token string, data *discordgo.WebhookParams) {

	(*session).WebhookExecute(id, token, false, data)
}
func IssueWebhook(ci string) (string, string, error) {

	wh, err := (*session).WebhookCreate(ci, "Groupme Sync", "")
	if err != nil {
		return "", "", err
	}
	return wh.ID, wh.Token, nil
}

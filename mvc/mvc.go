package mvc

import (
	"fmt"
	"log"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
)

func GroupmeReceive(name, avatar, content, groupID string) error {
	wi, wt, err := db.WebhookFromGroupID(groupID)
	if err != nil {
		log.Print(groupID + ": ")
		log.Println(err)
		if err.Error() == "No Webhook" {
			ci, err := db.ChannelIDFromGroupID(groupID)
			if err != nil {
				return err
				log.Print("serious issue with " + groupID)
				log.Println(err)
			}
			wi, wt, err = discord.IssueWebhook(ci)
			if err != nil {
				return err
				log.Print("serious issue with " + groupID)
				log.Println(err)
			}

			err = db.AddWebhook(groupID, wi, wt)
			if err != nil {
				return err
				log.Print("serious issue with " + groupID)
				log.Println(err)
			}
		}
	}

	fmt.Println(wi, wt)
	return nil
}
func DiscordSend(name, content, avatar, webhookURL, webhookToken string) error { return nil }

func DiscordReceive() {}
func GroupmeSend()    {}

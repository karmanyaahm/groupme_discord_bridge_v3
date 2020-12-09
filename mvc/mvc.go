package mvc

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
)

func GroupmeReceive(name, avatar, content, groupID string, attachments interface{}) error {
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

	DiscordSend(name, content, avatar, wi, wt, attachments)
	return nil
}
func DiscordSend(name, content, avatar, webhookID, webhookToken string, attachments interface{}) error {
	data := discordgo.WebhookParams{Content: content, Username: name, AvatarURL: avatar}

	discord.CallWebhook(webhookID, webhookToken, &data)

	return nil
}

func DiscordReceive() {}
func GroupmeSend()    {}

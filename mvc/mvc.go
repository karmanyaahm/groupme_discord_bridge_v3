package mvc

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord_utils"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/groupme_send"
)

func GroupmeReceive(name, avatar, content, groupID string, attachments []map[string]interface{}) error {
	wi, wt, err := db.WebhookFromGroupID(groupID)
	if err != nil {
		log.Print(groupID + ": ")
		log.Println(err)
		if err.Error() == "No Webhook" {
			ci, err := db.ChannelIDFromGroupID(groupID)
			if err != nil {
				log.Print("serious issue with " + groupID)
				log.Println(err)
				return err
			}
			wi, wt, err = discord_utils.IssueWebhook(ci)
			if err != nil {
				log.Print("serious issue with " + groupID)
				log.Println(err)
				return err
			}

			err = db.AddWebhook(groupID, wi, wt)
			if err != nil {
				log.Print("serious issue with " + groupID)
				log.Println(err)
				return err
			}
		} else {
			return err
		}
	}

	DiscordSend(name, content, avatar, wi, wt, attachments)
	return nil
}
func DiscordSend(name, content, avatar, webhookID, webhookToken string, attachments []map[string]interface{}) error {
	data := discordgo.WebhookParams{Content: content, Username: name, AvatarURL: avatar}
	for _, i := range attachments {
		if i["type"] == "image" {
			data.Embeds = append(data.Embeds, &discordgo.MessageEmbed{
				URL: i["url"].(string),
				Image: &discordgo.MessageEmbedImage{
					URL: i["url"].(string),
				},
			})
		} else {
			log.Println(i)
		}
	}
	discord_utils.CallWebhook(webhookID, webhookToken, &data)

	return nil
}

func DiscordReceive(name, content, channelID string, attachments []*discordgo.MessageAttachment) error {
	channelName, e := db.ChannelNameFromChannelID(channelID)

	if e != nil {
		log.Print(channelID + ": ")
		log.Println(e)
		if e.Error() == "No Name" {
			channelName, err := discord_utils.GetChannelName(channelID)
			if err != nil {
				log.Println("serious issue with " + channelID)
				log.Println(err)
				return err
			}
			err = db.AddChannelName(channelID, channelName)
			if err != nil {
				log.Println("serious issue with " + channelID)
				log.Println(err)
				return err
			}
		} else {
			return e
		}
	}

	botID, err := db.BotIDFromChannelID(channelID)
	if err != nil {
		log.Println("Groupme Bot Not Found")
		log.Println(err)
		log.Println(channelID)
	}

	GroupmeSend(name, content, channelName, botID, attachments)
	return nil
}

func GroupmeSend(name, content, channelName, botID string, attachments []*discordgo.MessageAttachment) {
	//log.Println(name, content, channelName)

	groupme_send.Send(botID, fmt.Sprintf("%s: %s: %s", channelName, name, content))
	for _, attachment := range attachments {
		if m, e := regexp.Match(`(?i)\.(gif|jpe?g|tiff?|png|webp|bmp)$`, []byte(attachment.Filename)); m {
			//log.Println("image")
			if e != nil {
				log.Println(e)
			}
			groupme_send.SendWithImage(botID, "", attachment.URL)
		}
	}

}

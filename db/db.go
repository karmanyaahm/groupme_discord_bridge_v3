package db

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

var conf Config
var DiscordToken string
var Addr string

type Connection struct {
	BotID   string `mapstructure:"groupme_bot_id" toml:"groupme_bot_id"`
	GroupID string `mapstructure:"groupme_group_id" toml:"groupme_group_id"`

	ChannelID        map[string]string `mapstructure:"discord_channels" toml:"discord_channels"`
	PrimaryChannelID string            `mapstructure:"primary_channel" toml:"primary_channel"`
	WebhookID        string            `mapstructure:"webhook_id,omitempty" toml:"webhook_id"`
	WebhookToken     string            `mapstructure:"webhook_token,omitempty" toml:"webhook_token"`
}

type Config struct {
	Address         string
	DiscordBotToken string
	Connection      []Connection
}

func Parse() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	conf = Config{}
	err := viper.Unmarshal(&conf)

	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	//	fmt.Println(conf)
	if conf.DiscordBotToken[:4] != "Bot " {
		conf.DiscordBotToken = "Bot " + conf.DiscordBotToken
	}
	DiscordToken = conf.DiscordBotToken
	//log.Println(conf.DiscordBotToken, conf.DiscordBotToken[:4])

	Addr = conf.Address
}

func BotIDFromChannelID(s string) (string, error) {
	for _, i := range conf.Connection {
		for j := range i.ChannelID {
			if j == s {
				return i.BotID, nil
			}
		}
	}
	return "", errors.New("Not Found")
}

func ChannelNameFromChannelID(s string) (string, error) {
	for _, i := range conf.Connection {
		for j, k := range i.ChannelID {
			if j == s {
				if len(k) == 0 {
					return "", errors.New("No Name")

				}
				return k, nil
			}
		}
	}
	return "", errors.New("Not Found")
}

func ChannelIDFromGroupID(s string) (string, error) {
	for _, i := range conf.Connection {
		if i.GroupID == s {
			if len(i.ChannelID) == 0 {
				return "", errors.New("Not Found")
			}
			return i.PrimaryChannelID, nil
		}
	}
	return "", errors.New("Not Found")
}

func WebhookFromGroupID(s string) (string, string, error) {
	for _, i := range conf.Connection {
		if i.GroupID == s {
			if (len(i.WebhookID) == 0) || (len(i.WebhookToken) == 0) {
				return "", "", errors.New("No Webhook")
			}
			return i.WebhookID, i.WebhookToken, nil
		}
	}
	return "", "", errors.New("Not Found")
}
func updateDB() error {

	b, err := toml.Marshal(conf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(viper.ConfigFileUsed(), b, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func AddWebhook(gid, wi, wt string) error {
	var index int
	//fmt.Println(gid, wi, wt)
	for j, i := range conf.Connection {
		if i.GroupID == gid {
			index = j
			break
		}
	}
	conf.Connection[index].WebhookID = wi
	conf.Connection[index].WebhookToken = wt
	return updateDB()
	//	return nil
}

func AddChannelName(channelID, channelName string) error {
	var index int
	for k, i := range conf.Connection {
		for j := range i.ChannelID {
			if j == channelID {
				index = k
			}
		}
	}
	conf.Connection[index].ChannelID[channelID] = channelName
	return updateDB()

}

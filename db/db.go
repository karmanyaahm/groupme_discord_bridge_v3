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

type Connection struct {
	BotID   string `mapstructure:"groupme_bot_id" toml:"groupme_bot_id"`
	GroupID string `mapstructure:"groupme_group_id" toml:"groupme_group_id"`

	ChannelID []string `mapstructure:"discord_channels" toml:"discord_channels" comment:"First channel is where groupmes are sent back to"`

	WebhookID    string `mapstructure:"webhook_id,omitempty" toml:"webhook_id"`
	WebhookToken string `mapstructure:"webhook_token,omitempty" toml:"webhook_token"`
}

type Config struct {
	Connection []Connection
}

func Parse() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	conf = Config{}
	err := viper.Unmarshal(&conf)

	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	//	fmt.Println(conf)

}

func BotIDFromChannelID(s string) (string, error) {
	for _, i := range conf.Connection {
		for _, j := range i.ChannelID {
			if j == s {
				return i.BotID, nil
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
			return i.ChannelID[0], nil
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

func AddWebhook(gid, wi, wt string) error {
	var index int
	for j, i := range conf.Connection {
		if i.GroupID == gid {
			index = j
			break
		}
	}
	conf.Connection[index].WebhookID = wi
	conf.Connection[index].WebhookToken = wt

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

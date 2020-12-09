package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var conf Config

type Connection struct {
	BotID   string `mapstructure:"groupme_bot_id"`
	GroupID string `mapstructure:"groupme_group_id"`

	ChannelID []string `mapstructure:"discord_channels"`

	WebhookID string `mapstructure:"webhook_id,omitempty"`
	Token     string `mapstructure:"webhook_token,omitempty"`
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

func ChannelIDFromBotID(s string) (string, error) {
	for _, i := range conf.Connection {
		if i.BotID == s {
			if len(i.ChannelID) == 0 {
				return "", errors.New("Not Found")
			}
			return i.ChannelID[0], nil
		}
	}
	return "", errors.New("Not Found")
}

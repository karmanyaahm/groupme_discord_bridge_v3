package groupme

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adammohammed/groupmebot"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
)

func Listen() {
	logger := groupmebot.StdOutLogger{}
	bot := groupmebot.GroupMeBot{Logger: logger}
	bot.Hooks = make(map[string]func(groupmebot.InboundMessage) string)
	bot.AddHook(".*", handler)
	bot.TrackBotMessages = false
	//bot.HandleMessage = handler

	http.HandleFunc("/", bot.Handler())

	log.Fatal(http.ListenAndServe(config.Addr, nil))

}

func handler(m groupmebot.InboundMessage) string {

	fmt.Println(m.Name)
	fmt.Println(m.Group_id)
	fmt.Println(m.Avatar_url)
	fmt.Println(m.Text)
	fmt.Println(m.Attachments)

	return ""
}

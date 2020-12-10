package groupme

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adammohammed/groupmebot"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/mvc"
)

var serv *http.Server

func Listen() {
	logger := groupmebot.StdOutLogger{}
	bot := groupmebot.GroupMeBot{Logger: logger}
	bot.Hooks = make(map[string]func(groupmebot.InboundMessage) string)
	bot.AddHook(".*", handler)
	bot.TrackBotMessages = false
	//bot.HandleMessage = handler

	mux := http.NewServeMux()

	mux.HandleFunc("/", bot.Handler())

	serv = &http.Server{Addr: config.Addr, Handler: mux}
	log.Fatal(serv.ListenAndServe())

}

func Close() {
	log.Println("Http Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() { cancel() }()
	serv.Shutdown(ctx)
	log.Println("Http Shut Down")
}

func handler(m groupmebot.InboundMessage) string {
	mvc.GroupmeReceive(m.Name, m.Avatar_url, m.Text, m.Group_id, m.Attachments)

	return ""
}

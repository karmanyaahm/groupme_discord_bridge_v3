package groupme

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adammohammed/groupmebot"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/mvc"
)

var serv *http.Server

func Listen() {
	logger := groupmebot.CompositeLogger{}
	bot := groupmebot.GroupMeBot{Logger: logger}
	bot.Hooks = make(map[string]func(groupmebot.InboundMessage) string)
	bot.AddHook(".*", handler)
	bot.TrackBotMessages = false
	//bot.HandleMessage = handler

	mux := http.NewServeMux()

	mux.HandleFunc("/", bot.Handler())

	serv = &http.Server{Addr: db.Addr, Handler: mux}
	go func() {
		err := serv.ListenAndServe()
		if err.Error() != "http: Server closed" {
			log.Println(err)
		}
	}()

}

func Close() {
	log.Println("Http Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Print("shutdown error: ")
		log.Println(err)
	}

	log.Println("Http Shut Down")
}

var receiveFunc = mvc.GroupmeReceive

func handler(m groupmebot.InboundMessage) string {
	receiveFunc(m.Name, m.Avatar_url, m.Text, m.Group_id, m.Attachments)

	return ""
}

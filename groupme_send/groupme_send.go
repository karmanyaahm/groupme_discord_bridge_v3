package groupme_send

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseurl = "https://api.groupme.com/v3"

type responseWithImage struct {
	BotID      string `json:"bot_id"`
	Text       string `json:"text"`
	PictureURL string `json:"picture_url,omitempty"`
}

//Send groupme message without image
func Send(BotID, content string) error {
	return SendWithImage(BotID, content, "")
}

//SendWithImage send groupme stuff with image
func SendWithImage(BotID, content, img string) error {
	img = ProcImage(img)

	buf, err := json.Marshal(responseWithImage{
		BotID:      BotID,
		Text:       content,
		PictureURL: img,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, err := http.Post(baseurl+"/bots/post", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 202 {
		log.Println("aaaaa" + string(buf))
		log.Println(string(bytes))
		log.Println(resp.StatusCode)
		return errors.New("Groupme Server Not Accept")
	}
	return nil
}

//ProcImage in groupme's image handler
func ProcImage(url string) string {
	//TODO: use groupme processing stuffs
	return url
}

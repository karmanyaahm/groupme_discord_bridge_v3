package groupme

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/config"
)

func TestAaaa(t *testing.T) {
	//send(config.GroupmeBotId, "aaa")
	SendWithImage(config.GroupmeBotId, "aa", "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1280px-Image_created_with_a_mobile_phone.png")
}
func TestReceive(t *testing.T) {
	//send(config.GroupmeBotId, "aaa")
	config.Addr = "localhost:53048"
	go Listen()
	time.Sleep(1 * time.Second)
	http.Post("http://localhost:53048/", "application/json", bytes.NewBuffer([]byte(`
	{
		"attachments": [        {
          "type": "image",
          "url": "https://i.groupme.com/123456789"
        },
        {
          "type": "image",
          "url": "https://i.groupme.com/123456789"
        },
        {
          "type": "location",
          "lat": "40.738206",
          "lng": "-73.993285",
          "name": "GroupMe HQ"
        },
        {
          "type": "emoji",
          "placeholder": "☃",
          "charmap": [
            [
              1,
              42
            ],
            [
              2,
              34
            ]
	  ]}],
		"avatar_url": "https://i.groupme.com/123456789",
		"created_at": 1302623328,
		"group_id": "1234567890",
		"id": "1234567890",
		"name": "John",
		"sender_id": "12345",
		"sender_type": "user",
		"source_guid": "GUID",
		"system": false,
		"text": "Hello world ☃☃",
		"user_id": "1234567890"
	  }`)))
}

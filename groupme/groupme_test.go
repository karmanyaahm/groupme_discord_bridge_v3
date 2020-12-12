package groupme

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
)

//func TestAaaa(t *testing.T) {
//	SendWithImage(config.GroupmeBotId, "aa", "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1280px-Image_created_with_a_mobile_phone.png")
//}
func TestReceive(t *testing.T) {
	db.Parse()
	db.Addr = "localhost:5000"
	works := false
	receiveFunc = func(name, avatar_url, text, group_id string, attachments []map[string]interface{}) error {
		works = name == "John" && avatar_url == "https://i.groupme.com/123456789" && text == "Hello world ☃☃" && group_id == "1234567890"
		return nil
	}

	Listen()
	time.Sleep(10 * time.Millisecond)
	http.Post("http://localhost:5000/", "application/json", testInput)
	Close()
	if !works {
		t.Fail()
	}
}

var testInput = bytes.NewBuffer([]byte(`
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
	  }`))

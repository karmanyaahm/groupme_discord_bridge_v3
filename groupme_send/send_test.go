package groupme_send

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	code := 202
	var req []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		rr, err := ioutil.ReadAll(r.Body)
		req = rr
		r.Body.Close()
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}))
	defer ts.Close()

	t.Log(ts.URL)
	baseurl = ts.URL

	err := SendWithImage("myid", "content", "url")
	if err != nil {
		t.Fatal(err)
	}
	if string(req) != `{"bot_id":"myid","text":"content","picture_url":"url"}` {
		t.Fatal(string(req))
	}

	err = Send("myid", "content")
	if err != nil {
		t.Fatal(err)
	}
	if string(req) != `{"bot_id":"myid","text":"content"}` {
		t.Fatal(string(req))
	}

	code = 403

	log.SetOutput(ioutil.Discard)
	err = Send("myid", "content")
	log.SetOutput(os.Stderr)

	if err.Error() != "Groupme Server Not Accept" {
		t.Fatal(err)
	}

}

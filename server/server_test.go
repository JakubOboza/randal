package server

import (
	"net/http/httptest"
	"testing"

	"github.com/JakubOboza/randal/config"
)

var (
	TEST_FILE_CONTENTS = `
root_url:
endpoints:
  test:
    destinations:
      - http://youtube.com/wow
      - http://google.com/wow2
      - http://bing.com/wow3
  test2:
    destinations:
      - http://lambdacu.be 
`
)

func TestServerBasicFlow(t *testing.T) {

	conf, err := config.Load([]byte(TEST_FILE_CONTENTS))

	if err != nil {
		t.Errorf("didnt expect error but got '%v'", err)
	}

	app := New(1234, conf)
	err = app.Setup()
	if err != nil {
		t.Errorf("didnt expect error but got '%v'", err)
	}

	req := httptest.NewRequest("GET", "/test", nil)

	resp, _ := app.engine.Test(req, 1)

	if resp.StatusCode != 302 {
		t.Errorf("expected status 302 redirected but got '%v'", resp.StatusCode)
	}

	location, err := resp.Location()
	location1 := location.String()
	if err != nil {
		t.Errorf("didnt expect error but got '%v'", err)
	}
	if location1 != "http://youtube.com/wow" && location1 != "http://google.com/wow2" && location1 != "http://bing.com/wow3" {
		t.Errorf("redirect location should point to one of configured locations but got '%v'", location1)
	}

	req2 := httptest.NewRequest("GET", "/test2", nil)
	resp, _ = app.engine.Test(req2, 1)

	if resp.StatusCode != 302 {
		t.Errorf("expected status 302 redirected but got '%v'", resp.StatusCode)
	}

	location, _ = resp.Location()
	location2 := location.String()
	if err != nil {
		t.Errorf("didnt expect error but got '%v'", err)
	}
	if location2 != "http://lambdacu.be" {
		t.Errorf("redirect location should point to configured location but got '%v'", location2)
	}

}

package config

import (
	"fmt"
	"testing"
)

var (
	TEST_FILE_CONTENTS = `
root_url:
endpoints:
  one:
    destinations:
      - http://youtube.com/wow
      - http://google.com/wow2
      - http://bing.com/wow3
  two:
    destinations:
      - http://lambdacu.be 
`
)

func TestConfigLoad(t *testing.T) {
	_, err := Load([]byte{1})
	if err == nil {
		t.Error("expected error but didnt get it")
	}
	if err.Error() != "yaml: control characters are not allowed" {
		t.Errorf("expected error '%v' but got '%v'", "yaml: control characters are not allowed", err)
	}

	cfg, err := Load([]byte(TEST_FILE_CONTENTS))

	if err != nil {
		t.Errorf("didnt expect error but got '%v'", err)
	}

	if len(cfg.EndPoints["one"].Destinations) != 3 {
		t.Errorf("expected 'one' endpoint config to have 3 destination but got '%v'", cfg.EndPoints["one"].Destinations)
	}

}
func TestConfigNextIsMoreLessConsistent(t *testing.T) {

	cfg, _ := Load([]byte(TEST_FILE_CONTENTS))

	counters := map[string]int{}

	for i := 0; i < 100000; i++ {
		counters[cfg.EndPoints["one"].Next()] += 1
	}

	for k, v := range counters {
		fmt.Println("Rolls: ", k, " ", v)
	}

	for k, v := range counters {
		if v < 30000 {
			t.Errorf("looks like the randomizer might not be super consistent for '%v' got '%v'", k, v)
		}
	}
}

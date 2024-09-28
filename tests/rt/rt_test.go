package rt_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

var (
	config = make(map[string]interface{})
)

func TestRT(t *testing.T) {
	c := flag.String("config", "", "json file with config")

	go func() {
		for {
			data, err := ioutil.ReadFile(*c)
			if err != nil {
				log.Println(err)
			}
			json.Unmarshal(data, config)

			time.Sleep(1000000000) // 1 Second
		}
	}()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{}

		servers := config["servers"].([]string)
		for i := 0; i < len(servers); i++ {
			go func() {
				r, err := http.Get(servers[i])
				if err != nil {
					response[servers[i]] = err.Error()
				} else {
					response[servers[i]] = strconv.Itoa(r.StatusCode)
				}
			}()
		}

		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(config["listen"].(string), nil)
}

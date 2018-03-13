package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
	"fmt"
)
const URL = "https://api.opendota.com/api/proplayers"

var allPlayers []Player
var playerMap map[string]Player

func init() {
	playerMap = make(map[string]Player)
}


func main() {
	channel := make(chan []Player)

	playerClient := http.Client{
		Timeout: time.Second * 10,
	}

	 go func(channel chan []Player) {
		for {
			players := <- channel
			for _, item := range players {
				playerMap[item.Name] = item
			}
		}
	}(channel)

	go RefreshPLayersArray(&playerClient, channel)


	//	start the server since we have initialized all data
	router := mux.NewRouter()
	router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		//	read from a file

		key := request.Header.Get("name")

		if player, ok := playerMap[key]; ok {
			bytes, _ := json.Marshal(player)
			responseWriter.Write([]byte(bytes))
		}else{
			http.Error(responseWriter, "Not found", 505)
		}

	})

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}


func RefreshPLayersArray(client *http.Client, channel chan []Player) {

	var playersToReturn []Player

	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("player-name", "name")

	response, getErr := client.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &playersToReturn)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//write to the channel
	channel <-  playersToReturn
	fmt.Println("Refreshing data")

	time.Sleep(5 * time.Minute)
	go RefreshPLayersArray(client, channel)
}

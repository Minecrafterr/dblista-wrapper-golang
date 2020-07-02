package dblistawrapper

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"bytes"
)

func GetBotInfoFull(id string, info string) {
	resp, err := http.Get("https://api.dblista.pl/v1/bots/"+id)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	gjson.Get(string(body), "data."+info).String()
}

func GetBotInfo(id string) {
	resp, err := http.Get("https://api.dblista.pl/v1/bots/"+id)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Name: " + gjson.Get(string(body), "data.name").String() + "\nRating average: " + gjson.Get(string(body), "data.rating.average").String()+"\nVotes count: "+gjson.Get(string(body), "data.votes").String()+"\nYour bot is boosted by user with id "+gjson.Get(string(body), "data.premium.booster").String()+"\nBoost ends at (timestamp) "+gjson.Get(string(body), "data.premium.end").String())
	}


func UpdateStats(token string, users string, servers string) {
	client := &http.Client{}
    
        body := []byte("{\n  \"members\":"+users+",\n  \"servers\":"+servers+"\n}")
    
        req, _ := http.NewRequest("POST", "https://api.dblista.pl/v1/bots/stats", bytes.NewBuffer(body))
    
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Authorization", token)
    
        resp, err := client.Do(req)
    
        if err != nil {
            fmt.Println("Errored when sending request to the server")
            return
        }
    
        defer resp.Body.Close()
        resp_body, _ := ioutil.ReadAll(resp.Body)
    
        fmt.Println(resp.Status)
        fmt.Println(string(resp_body))
}

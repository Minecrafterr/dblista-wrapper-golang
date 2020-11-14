package dblistawrapper

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"strconv"
)
type VotelogModel struct {
    Timestamp int `json:"timestamp"`
    User string `json:"user"`
}
type RatingsModel struct {
    Author string `json:"author"`
    Rating int `json:"rating"`
    Details string `json:"details"`
}
type InfoModel struct {
    FullDescription string `json:"fullDescription"`
    Library string `json:"library"`
    Prefix string `json:"prefix"`
    ShortDescription string `json:"shortDescription"`
    Tags []string `json:"tags"`
}
type LinksModel struct {
    Discord string `json:"discordServer"`
    Git string `json:"gitRepository"`
    Www string `json:"website"`
}
type PremiumModel struct {
    Booster string `json:"booster"`
    CardColor string `json:"cardStyle"`
    End string `json:"end"`
    Start string `json:"start"`
}
type RatingModel struct {
    Avarage int `json:"average"`
    Order int `json:"order"`
    Ratings []RatingsModel `json:"ratings"`
}
type StatModel struct {
    Members int `json:"members"`
    MonthlyInvites int `json:"monthlyInvites"`
    Servers int `json:"servers"`
    Status string `json:"status"`
}
type StatusModel struct {
    Verification string `json:"verification"`
}
type UptimeModel struct {
    Max int `json:"max"`
    Online int `json:"online"`
}
type BotModel struct {
    Avatar string `json:"avatarURL"`
    ID string  `json:"id"`
    Name string  `json:"name"`
    Owner string  `json:"ownerID"`
    Votes uint16  `json:"votes"`
    Owners []string  `json:"owners"`
    InvitePermissions int `json:"invitePermissions"`
    Votelog []VotelogModel `json:"votelog"`
    Info InfoModel `json:"info"`
    Links LinksModel `json:"links`
    Ratings RatingModel `json:"rating"`
    Stats StatModel `json:"stats"`
    Status StatusModel `json:"status"`
    Uptime UptimeModel `json:"uptime"`
}
type BotRes struct {
    Data BotModel `json:"data"`
  }

type ServerVotelogModel struct {
        Timestamp int `json:"timestamp"`
        User string `json:"user"`
    }
    type ServerRatingsModel struct {
        Author string `json:"author"`
        Rating int `json:"rating"`
        Details string `json:"details"`
    }
    type ServerInfoModel struct {
        FullDescription string `json:"fullDescription"`
        ShortDescription string `json:"shortDescription"`
        Tags []string `json:"tags"`
    }
    type ServerLinksModel struct {
        Discord string `json:"discordServer"`
        Www string `json:"website"`
    }
    type ServerPremiumModel struct {
        Booster string `json:"booster"`
        CardColor string `json:"cardStyle"`
        End string `json:"end"`
        Start string `json:"start"`
    }
    type ServerRatingModel struct {
        Avarage int `json:"average"`
        Order int `json:"order"`
        Ratings []ServerRatingsModel `json:"ratings"`
    }
    type ServerStatModel struct {
        Max int `json:"max"`
        Online int `json:"online"`
    }
    type ServerModel struct {
        Avatar string `json:"avatarURL"`
        ID string  `json:"id"`
        Name string  `json:"name"`
        Owner string  `json:"ownerID"`
        Votes uint16  `json:"votes"`
        InvitePermissions int `json:"invitePermissions"`
        Votelog []ServerVotelogModel `json:"votelog"`
        Info ServerInfoModel `json:"info"`
        Links ServerLinksModel `json:"links`
        Ratings ServerRatingModel `json:"rating"`
        Stats ServerStatModel `json:"stats"`
    }
type ServerRes struct {
    Data ServerModel `json:"data"`
  }

type UserPremium struct {
        BoostedBot string `json:"boostedBot"`
        BoostedServer string `json:"boostedServer"`
        End int `json:"end"`
        Start int `json:"start"`
    }
    type UserModel struct {
        Avatar string `json:"avatar"`
        ID string  `json:"id"`
        LastRequestTime int `json:"lastRequestTime"`
        RecentRequests int `json:"recentRequests"`
        Username string  `json:"username"`
        Ban bool `json:"ban"`
        Money int `json:"invitePermissions"`
        Perm int `json:"perm"`
        Premium UserPremium `json:"premium"`
    }
    type UserRes struct {
        Data UserModel `json:"data"`
      }

func VoteServer(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/servers/"+id+"/vote", nil)

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}
func RateServer(id string, rate int, description string, token string) {
	client := &http.Client{}

	body := []byte("{\n  \"rating\": "+rate+",\n  \"details\":"+description+"\n}")

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/servers/"+id+"/rate", bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}
func BoostServer(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/servers/"+ID+"/boost", nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}
func RemoveServerBoost(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "https://api.dlist.top/v1/servers/"+ID+"/boost", nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}

func BoostBot(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/bots/"+ID+"/boost", nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}

func RemoveBotBoost(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("DELETE", "https://api.dlist.top/v1/bots/"+ID+"/boost", nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}

func RateBot(id string, rate int, description string, token string) {
	client := &http.Client{}

	body := []byte("{\n  \"rating\": "+rate+",\n  \"details\":"+description+"\n}")

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/bots/"+id+"/rate", bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}
func VoteBot(id string, token string) {
client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/bots/"+id+"/vote", nil)

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status+"\n"+string(resp_body)
}
func GetUserInfo(id string) UserRes {
	resp, err := http.Get("https://api.dlist.top/v1/users/"+id)
	if err != nil {
        // handle err
        return err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    
    var d UserRes
    json.Unmarshal(body, &d)
    return d
}

func GetBotInfo(id string) BotRes {
	resp, err := http.Get("https://api.dlist.top/v1/bots/"+id)
if err != nil {
    // handle err
    return err
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    return err
}

var d BotRes
json.Unmarshal(body, &d)
return d
	}
func GetServerInfo(id string) ServerRes {
resp, err := http.Get("https://api.dlist.top/v1/servers/"+id)
    if err != nil {
        // handle err
        return err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    
    var d ServerRes
    json.Unmarshal(body, &d)
    return d
}

func UpdateStats(token string, users int, servers int) {
	client := &http.Client{}
	u := strconv.Itoa(users)
	s := strconv.Itoa(servers)
        body := []byte("{\n  \"members\":"+u+",\n  \"servers\":"+s+"\n}")
    
        req, _ := http.NewRequest("POST", "https://api.dlist.top/v1/bots/stats", bytes.NewBuffer(body))
    
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Authorization", token)
    
        resp, err := client.Do(req)
    
        if err != nil {
            fmt.Println("Errored when sending request to the server")
            return
        }
    
        defer resp.Body.Close()
        resp_body, _ := ioutil.ReadAll(resp.Body)
    
        return resp.Status+"\n"+string(resp_body)
}

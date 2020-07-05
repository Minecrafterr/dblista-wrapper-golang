package dblistawrapper

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
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

func GetUserInfo(id string) UserRes {
	resp, err := http.Get("https://api.dblista.pl/v1/users/"+id)
	if err != nil {
        // handle err
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    var d UserRes
    json.Unmarshal(body, &d)
    return d
}

func GetBotInfo(id string) BotRes {
	resp, err := http.Get("https://api.dblista.pl/v1/bots/"+id)
if err != nil {
    // handle err
    fmt.Println(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Println(err)
}

var d BotRes
json.Unmarshal(body, &d)
return d
	}
func GetServerInfo(id string) ServerRes {
resp, err := http.Get("https://api.dblista.pl/v1/servers/"+id)
    if err != nil {
        // handle err
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    var d ServerRes
    json.Unmarshal(body, &d)
    return d
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

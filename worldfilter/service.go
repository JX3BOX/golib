package worldfilter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"gopkg.in/JX3BOX/golib.v1/utils"
)

var client = http.Client{
	Timeout: 3 * time.Second,
}

type Service struct {
	API       string
	AppId     string
	ChannelId string
	SecretKey string
	fnSwitch  ISwitch // 是否进行过滤开关
}

type Result struct {
	Code       int    `json:"code"`
	RequestId  string `json:"requestId"`
	ResultType int    `json:"resultType"`
	Content    string `json:"content"`
	Reason     string `json:"reason"`
}

type ISwitch func() bool

func (s *Service) SetSwitch(fn ISwitch) {
	s.fnSwitch = fn
}

func (s Service) Filter(originContent string) (result Result, err error) {
	if s.fnSwitch != nil && !s.fnSwitch() {
		return Result{
			Content:    originContent,
			ResultType: 0,
			Code:       1,
		}, nil
	}

	var data = map[string]string{
		"appID":     s.AppId,
		"channelID": s.ChannelId,
		"requestID": utils.RandString(8),
		"content":   originContent,
	}

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signRaw := []string{}
	for _, k := range keys {
		signRaw = append(signRaw, data[k])
	}
	signRaw = append(signRaw, s.SecretKey)
	sign := utils.MD5(strings.Join(signRaw, ""))
	data["signature"] = sign
	jBody, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", s.API, bytes.NewReader(jBody))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &result)
	return
}

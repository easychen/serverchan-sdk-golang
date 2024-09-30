package serverchan_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScSendOptions struct {
	Tags    string `json:"tags,omitempty"`
	Short   string `json:"short,omitempty"`
	Noip    int    `json:"noip,omitempty"`
	Channel string `json:"channel,omitempty"`
	Openid  string `json:"openid,omitempty"`
}

type ScSendResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ScSend(sendkey, title, desp string, options *ScSendOptions) (*ScSendResponse, error) {
	var url string
	if len(sendkey) >= 4 && sendkey[:4] == "sctp" {
		url = fmt.Sprintf("https://%s.push.ft07.com/send", sendkey)
	} else {
		url = fmt.Sprintf("https://sctapi.ftqq.com/%s.send", sendkey)
	}

	params := map[string]interface{}{
		"title": title,
		"desp":  desp,
	}
	if options != nil {
		if options.Tags != "" {
			params["tags"] = options.Tags
		}
		if options.Short != "" {
			params["short"] = options.Short
		}
		if options.Noip != 0 {
			params["noip"] = options.Noip
		}
		if options.Channel != "" {
			params["channel"] = options.Channel
		}
		if options.Openid != "" {
			params["openid"] = options.Openid
		}
	}

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ScSendResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

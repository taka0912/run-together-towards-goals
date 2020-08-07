package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Text string `json:"text"`
}

const adminUrl = "https://hooks.slack.com/services/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"

func NoticeForgotPass(text string) {
	log.Print(text)
	err := postSlack(request{Text: "Forgot Password : " + text})
	if err != "" {
		log.Println(err)
	}
}

func postSlack(requestBody interface{}) string {
	jsonModel, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	req, err := http.NewRequest("POST", adminUrl, bytes.NewBuffer(jsonModel))
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return ""
}

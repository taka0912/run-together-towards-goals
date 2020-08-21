package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Method string
	Url    string
	Body   interface{}
}


type responseJson struct {
	Code float64 `json:"code,string"`
	Msg  string  `json:"msg"`
	ID   float64 `json:"id,string"`
}

func tryTestRequest(reqInfo Request) interface{} {
	jsonModel, _ := json.Marshal(reqInfo.Body)
	req, err := http.NewRequest(reqInfo.Method, reqInfo.Url, bytes.NewBuffer(jsonModel))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
		return nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
		return nil
	}
	defer resp.Body.Close()

	var responseJson map[string]interface{}

	err = json.Unmarshal(respBody, &responseJson)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
		return nil
	}

	return responseJson
}

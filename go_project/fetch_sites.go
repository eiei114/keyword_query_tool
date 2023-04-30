package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseLinks struct {
	Links []string `json:"links"`
}

func FetchLinks(scriptID, apiKey, customSearchID string) []string {
	scriptURL := fmt.Sprintf("https://script.google.com/macros/s/%s/exec?KEY=%s&ID=%s", scriptID, apiKey, customSearchID)

	response, err := http.Get(scriptURL)
	if err != nil {
		fmt.Println("Error executing script:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	var resp responseLinks
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil
	}

	fmt.Println("Script response: Links:", resp.Links)
	return resp.Links
}

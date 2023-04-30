package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Response struct {
	Links []string `json:"links"`
}
type MediaData struct {
	URLs  []string `json:"urls"`
	Media []string `json:"media"`
}

func main() {
	print("Server is running at http://localhost:8080")
	scriptId, err := loadEnv("APP_SCRIPT_ID")
	if err != nil {
		panic(err)
	}
	customSearchKey, err := loadEnv("CUSTOM_SEARCH_API_KAY")
	if err != nil {
		panic(err)
	}
	searchEngineId, err := loadEnv("SEARCH_ENGINE_ID")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/get-links", func(w http.ResponseWriter, r *http.Request) {
		getLinks(w, r, scriptId, customSearchKey, searchEngineId)
	})

	http.HandleFunc("/post-media-class", func(w http.ResponseWriter, r *http.Request) {
		postMediaClass(w, r, scriptId)
	})

	http.ListenAndServe(":8080", nil)
}

func getLinks(w http.ResponseWriter, r *http.Request, scriptId string, customSearchKey string, searchEngineId string) {
	// リンク取得処理を実装
	links := FetchLinks(scriptId, customSearchKey, searchEngineId)

	response := Response{
		Links: links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func postMediaClass(w http.ResponseWriter, r *http.Request, scriptId string) {
	// GAS WebアプリURLを生成
	gasWebAppURL := "https://script.google.com/macros/s/" + scriptId + "/exec"
	fmt.Println("GAS Web App URL loaded:", gasWebAppURL)

	// リクエスト本文からJSONデータを読み取る
	var mediaData MediaData
	err := json.NewDecoder(r.Body).Decode(&mediaData)
	if err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}
	fmt.Println("URLs:", mediaData.URLs)
	fmt.Println("Media:", mediaData.Media)

	// GASへのリクエストを作成
	formData := url.Values{}
	urlsJson, err := json.Marshal(mediaData.URLs)
	if err != nil {
		fmt.Println("Error marshaling URLs:", err)
		return
	}
	mediaJson, err := json.Marshal(mediaData.Media)
	if err != nil {
		fmt.Println("Error marshaling Media:", err)
		return
	}
	formData.Add("urls", string(urlsJson))
	formData.Add("media", string(mediaJson))

	req, err := http.NewRequest("POST", gasWebAppURL, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Request created")

	// Content-Typeヘッダーを設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// リクエストを実行
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Request executed")
	defer resp.Body.Close()

	// 応答データを読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Response data read: " + string(body))

	// 応答データを返す
	w.Header().Set("Content-Type", "application/json")
	responseJson := `{"result": "success", "message": "Data processed."}`
	w.Write([]byte(responseJson))
}

// env読み込み関数
func loadEnv(variable string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv(variable), nil
}

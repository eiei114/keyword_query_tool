package go_project

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Links []string `json:"links"`
}

func getLinks(w http.ResponseWriter, r *http.Request) {
	// リンク取得処理を実装
	links := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
	}

	response := Response{
		Links: links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/get-links", getLinks)
	http.ListenAndServe(":8080", nil)
	print("Server is running at http://localhost:8080")
}

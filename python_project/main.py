import requests
import json
from media_classifier import MediaClassifier

GO_API_URL = "http://go_project:8080"


def main():
    print("Start scraping...")
    # Go APIサーバーからリンクを取得
    response = requests.get(f"{GO_API_URL}/get-links")
    links = response.json()["links"]

    # 取得したリンクをスクレイピング
    url_media_pairs = []
    for link in links:
        url, media_type = scrape_link(link)
        url_media_pairs.append((url, media_type))

    # URLと媒体の種類の配列をGo APIサーバーに送信
    urls = [pair[0] for pair in url_media_pairs]
    media = [pair[1] for pair in url_media_pairs]
    data = {"urls": urls, "media": media}

    headers = {"Content-Type": "application/json"}
    print("Sending URLs and media types to the Go API server...")
    response = requests.post(f"{GO_API_URL}/post-media-class", data=json.dumps(data), headers=headers)
    print("Response from Go API server:", response.json())


def scrape_link(url):
    print(f"Scraping {url}...")
    # MediaClassifier
    media_classifier = MediaClassifier(url)
    media_type = media_classifier.get_media_type()
    print(f"媒体: {media_type}")

    return url, media_type


if __name__ == "__main__":
    main()

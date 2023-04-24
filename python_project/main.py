import requests
from bs4 import BeautifulSoup

GO_API_URL = "http://go_project:8080/get-links"


def main():
    print("Start scraping...")
    # Go APIサーバーからリンクを取得
    response = requests.get(GO_API_URL)
    links = response.json()["links"]

    # 取得したリンクをスクレイピング
    for link in links:
        scrape_link(link)


def scrape_link(url):
    response = requests.get(url)

    print(f"Scraping {url}...")

    # スクレイピング処理を実装
    # ...


if __name__ == "__main__":
    main()

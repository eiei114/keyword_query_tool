### Premise
私はDockerとDocker Composeを使って同時に実行したいGoプロジェクトとPythonプロジェクトを持っています 。
GoプロジェクトはHTTPサーバーを提供し、PythonプロジェクトはGoプロジェクトのエンドポイントにリクエストを行えるようにしたい。
最終的に

現状のレポジトリ(https://github.com/eiei114/keyword_query_tool)

```tree
D:.
│  compose.yaml
│  README.md
│
├─go_project
│      Dockerfile
│      go.mod
│      main.go
│
└─python_project
        Dockerfile
        main.py
        requirements.txt

```

```go
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
```

```python
import requests
from bs4 import BeautifulSoup

GO_API_URL = "http://go_project:8080/get-links"


def main():
    print("Start scraping...")
    
    response = requests.get(GO_API_URL)
    links = response.json()["links"]

    for link in links:
        scrape_link(link)


def scrape_link(url):
    response = requests.get(url)

    print(f"Scraping {url}...")

if __name__ == "__main__":
    main()
```

### Dockerfileが正式に構築できていないのではないかと思う

```dockerfile
# go_scraper/Dockerfile
FROM golang:1.20

# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# Goモジュールの初期化
RUN go mod init example.com/go_project

COPY . /go/src/app

COPY . .
RUN go build -o main .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
```

```dockerfile
# ベースイメージの選択
FROM python:3.9
USER root

RUN apt-get update
RUN apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

COPY . .

# 必要なパッケージのインストール
RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "main.py"]
```

```yml
# docker-compose.yml
version: '3.8'

services:
  go_project:
    build: ./go_project
    container_name: go_project
    tty: true
    ports:
      - "8080:8080"

  python_project:
    build: ./python_project
    container_name: python_project
    working_dir: '/root/python_project'
    tty: true
    depends_on:
      - go_project
```

### 気になる部分
- DockerfileのCMDの部分
- docker compose exec go_project bashでコンテナに入ろうとすると`service "go_project" is not running container #1`というエラーが出るPythonも同様


### Problems / Error messages that are occurring

### Applicable source code

`` ```

`` ```

### What I tried


### Supplementary information (FW / tool version, etc.)
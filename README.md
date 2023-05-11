# keyword_query_tool

## Overview
- 一覧Sheetにキーワードを入力
- 一覧Sheetに入力したキーワードを元に、検索対象のシートからキーワードを含む行を抽出し、検索結果シートに出力する
 - どのような情報が分かるの？
 ![image](https://user-images.githubusercontent.com/60887155/235364639-681173ea-7cd2-441f-ae60-d82038fac168.png)
- Cron jobで一日に一回更新するようになっている

## 設計
![無題のプレゼンテーション (4)](https://user-images.githubusercontent.com/60887155/235364662-33e63241-d83a-4690-8a0d-ebbc51d224ac.png)

## フロー
1. Gitにアップロード
2. Actions起動
3. Docker Compose
4. Goの仮想APIサーバーを起動
5. スクレイピング処理を行う仮想Pythonサーバー起動
6. AppScriptにGetリクエストしリンクを収集
7. シートに書き込み
8. リンクをGoサーバーからPythonサーバーにリレー
9. スクレイピングし何の媒体の記事なのか判別しシートまで返す

# 開発環境構築(How to build development environment)
Go,ngrokを使用して「HelloWorld」を表示するまで。

1. Dockerfileからimage作成
`docker build -t go_alpine -f DockerfileForDev .`
-fオプションはDockerfileのファイル名がDockerfileでないときに必要なオプション

2. コンテナ起動
`docker run -d -e "PORT=3010" -p 3010:3010 -v /Users/moriyama/Desktop/study/go_wordwolf:/go/src/github.com/go-server-dev -t go_alpine`

-d：コンテナをバックグランドで実行
-e：環境変数設定
-p：ポートの設定
-v：ホストPCとのマウント設定。※コロンより左は自分がマウントしたい絶対パスを入力
-t：build時に「go_alpine」という名前でタグをつけたので、タグ指定で起動することが可能

3. コンテナのIDを確認
`docker ps`

4. IDを指定してコンテナの中に入る。
`docker exec -it 60c959e8d29d ash`
ash：alpineを操作する際に使用するコマンド

5. main.goファイルをビルド
`go build main.go`

6. ポート3010で起動
`./main 3010`

7. 別ターミナルで同じくコンテナに入りngrokを3010で起動
`ngrok http 3010`

8. ブラウザで表示されたURLに接続し、「HelloWorld!!!」が出れば成功
```
...
Forwarding       http://3fbffcc069d1.ngrok.io -> http://localhost:3010
Forwarding       https://3fbffcc069d1.ngrok.io -> http://localhost:3010
...
```

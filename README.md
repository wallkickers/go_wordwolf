# 開発環境構築(How to build development environment)

Go,ngrok を使用して「HelloWorld」を表示するまで。

**1. Dockerfile から image 作成**

```
docker build -t go_alpine -f DockerfileForDev .
```

-f オプションは Dockerfile のファイル名が Dockerfile でないときに必要なオプション

**2. コンテナ起動**

```
docker run -d -e "PORT=3010" -p 3010:3010 -v /Users/moriyama/Desktop/study/go_wordwolf:/go/src/github.com/go-server-dev -t go_alpine
```

-d：コンテナをバックグランドで実行  
-e：環境変数設定  
-p：ポートの設定  
-v：ホスト PC とのマウント設定。※コロンより左は自分がマウントしたい絶対パスを入力  
-t：build 時に「go_alpine」という名前でタグをつけたので、タグ指定で起動することが可能

**3. コンテナの ID を確認**

```
docker ps
```

**4. ID を指定してコンテナの中に入る。**

```
docker exec -it 60c959e8d29d ash
```

ash：alpine を操作する際に使用するコマンド

**5. main.go ファイルをビルド**

```
go build main.go
```

**6. ポート 3010 で起動**

```
./main 3010
```

**7. 別ターミナルで同じくコンテナに入り ngrok を 3010 で起動**

```
ngrok http 3010
```

**8. ブラウザで表示された URL に接続し、「HelloWorld!!!」が出れば成功**

```
...
Forwarding http://3fbffcc069d1.ngrok.io -> http://localhost:3010
Forwarding https://3fbffcc069d1.ngrok.io -> http://localhost:3010
...
```

# LINEBot 設定(LINEBot setting method)

**1. LINE Developers コンソールでチャネルを作成**

> 参考：LINE Developers コンソールでチャネルを作成する
> https://developers.line.biz/ja/docs/messaging-api/getting-started/#using-console

**2. チャネル内の「Channel secret」と「Channel access token」をメモ**

**3. main.go ファイルに下記の形で記載**

```
func lineHandler(w http.ResponseWriter, r \*http.Request) {
bot, err := linebot.New(
// TODO: 開発者の LINE アカウントごとに変更する必要あり
// "(自分のシークレットを入力)",
// "(自分のアクセストークンを入力)",
"<Channel secret>",
"<Channel access token>",
)
...
```

**4. コンテナ内で `ngrok http localhost:8000` を実行。**

下記のように表示されるので https の方をメモ

(例)

```
Forwarding http://2403c85009c4.ngrok.io -> http://localhost:8000
Forwarding https://2403c85009c4.ngrok.io -> http://localhost:8000
```

**5. LINE Developers 管理画面の「Webhook URL」に下記の形で記載。**

```
Webhook URL
<ngrok で作成した https の URL>/callback
```

**6. 友達登録後、メッセージを送信し、返信があれば成功。**

友達登録は LINE Developers 管理画面の QR コードを読み込むことで可能。

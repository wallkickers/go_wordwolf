# 開発環境構築(How to build development environment)

Go,ngrok を使用して「HelloWorld」を表示するまで。

**1. docker-compose.ymlを用いてコンテナ作成&起動**

```
cd env
docker-compose up -d
```
-d：コンテナをバックグランドで実行  

**2. コンテナ起動確認**
```
docker-compose ps
```
全てUpになっていれば成功。

**3. コンテナ名を指定してコンテナの中に入る。**

```
# botコンテナ(アプリケーションサーバ)の場合
docker-compose exec bot ash
```
ash：alpine を操作する際に使用するコマンド


**4. main.go ファイルをビルド**
```
go build main.go
```

**5. ポート 3010 で起動**

```
./main 3010
```

**6. 別ターミナルで同じくコンテナに入り ngrok を 3010 で起動**

```
ngrok http 3010
```

**7. ブラウザで表示された URL に接続し、「HelloWorld!!!」が出れば成功**

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

# 開発環境構築【Dockerfileからコンテナ作成】(How to build development environment)
※基本は上部の「開発環境構築」で良いが備忘録として残しておく。

Go,ngrok を使用して「HelloWorld」を表示するまで。

**1. Dockerfile から image 作成**

```
cd env
docker-compose up -d
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

# golang-migrateによるDBへのテーブル追加(How to create table using Migration by golang-migrate)
migrateファイルでテーブルを追加する。

**1. env\mysql\migrateに{連番}_{title}.up.sqlと{連番}_{title}.down.sql作成**
(例)初期のテーブル追加
作成用　　　　：1_initialize_schema.up.sql
失敗時に戻す用：1_initialize_schema.down.sql

**2. upにはcreate tableを、downにはdropする処理を記載**
```up.sql
create table user (
    id integer auto_increment primary key,
    name varchar(40)
)
```

```down.sql
drop tabel user
```

**3. botコンテナに入り、migrateコマンド実行**
```
# botコンテナに入る
docker-compose exec bot ash
# migrateコマンド実行
migrate -source file://env/mysql/migrate -database 'mysql://root:password@tcp(mysql:3306)/wordwolf' up 1
```

-source：migrateファイルのディレクトリ
-database：DBの指定
mysql：DBのコンテナ名
root,password：DBへの接続情報
mysql:3306：DBのコンテナ名とポート(compose.yml参照)
wordwolf：DB名
up：up.sqlを実行
1：連番指定

※downするときは「up」を「down」で実行。
```
# migrateコマンド実行
migrate -source file://env/mysql/migrate -database 'mysql://root:password@tcp(mysql:3306)/wordwolf' down 1
```

**4. migrate実行を確認**
```
# dbコンテナに入る
docker-compose exec mysql bash
# migrateコマンド実行
mysql -u root -p
# パスワード入力
password
# DB「wordwolf」指定
use wordwolf
# テーブル確認
show tables;
```
userテーブルが追加されていれば成功。

## 参考
【docker-compose系】  
・docker-compose.ymlでDockerfileを指定したい  
https://cloudpack.media/44104  

【DB系】  
・golang-migrate/migrate  
https://github.com/golang-migrate/migrate/tree/master/database/mysql  

・go modules環境でgolang-migrate/migrateを動かす  
https://yuzu441.hateblo.jp/entry/2019/06/11/150000  

Dockerでmysqlに接続できない(GO/gin/GORM)  
https://qiita.com/paragaki/items/9bba24c57e468400cb2c#%E8%A7%A3%E6%B1%BA%E7%AD%96  

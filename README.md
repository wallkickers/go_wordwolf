# go_wordwolf

# go実行環境構築_docker
# ルートディレクトリで下記を実行。
# 1. タグ「godocker」でimage作成
docker build -t godocker .
# 2. ポート3000でイメージからコンテナ実行
docker run -e "PORT=3000" -p 3000:3000 -t godocker
# 3. ウェブブラウザで「localhost:3000」で接続し、「Hello World」と表示されれば成功
# ４. コンテナに入る時
docker exec -it {コンテナID} sh
FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/yokoe/go-server-example
COPY . .
RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/yokoe/go-server-example /app

CMD /app/main $PORT
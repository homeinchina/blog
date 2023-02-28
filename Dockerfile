FROM golang:alpine

WORKDIR /app

COPY . /app

ENV GOPROXY https://goproxy.cn,direct

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o server main.go

# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

CMD ./server
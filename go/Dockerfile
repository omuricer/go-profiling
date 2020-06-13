# FROM golang:1.13-alpine
# alpineにはgitコマンドがなく、gitコマンドに依存しているgo getができないため、開発環境では使わない
# 本番用のマルチステージビルドのときは軽量化のためにalipineを使用する
FROM golang:1.13

# タイムゾーンとロケールの設定
RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apt update -y \
  && apt install -y locales \
  && locale-gen ja_JP.UTF-8 \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
ENV LANG ja_JP.UTF-8  
ENV LANGUAGE ja_JP:en  
ENV LC_ALL ja_JP.UTF-8

RUN go get -u github.com/gorilla/mux \
  github.com/oxequa/realize \
  github.com/go-sql-driver/mysql \
  github.com/jinzhu/gorm \
  github.com/gorilla/schema \
  github.com/yeka/zip \
  github.com/joho/godotenv \
  golang.org/x/text/encoding/japanese \
  golang.org/x/text/transform \
  github.com/alexmullins/zip \
  gopkg.in/gomail.v2 \
  github.com/go-playground/validator \
  github.com/google/wire/cmd/wire \
  github.com/fukata/golang-stats-api-handler

WORKDIR /go/src/github.com/omuricer/go-profiling

FROM ubuntu
RUN apt-get update \
  && apt-get install -y locales \ 
  golang-go \
  git \
  && locale-gen ja_JP.UTF-8 \
  && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc
WORKDIR /go/src
RUN go mod init example.com/m/v2 \
  && go mod tidy
  # && go get github.com/jinzhu/gorm \
  # && go get github.com/go-sql-driver/mysql \
  # && go get github.com/stretchr/testify/assert

# いつかちゃんとやる。いつかね
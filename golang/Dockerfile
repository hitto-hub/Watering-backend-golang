FROM ubuntu
RUN apt-get update \
  && apt-get install -y locales \ 
  golang-go \
  && locale-gen ja_JP.UTF-8 \
  && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc
# いつかちゃんとやる。いつかね
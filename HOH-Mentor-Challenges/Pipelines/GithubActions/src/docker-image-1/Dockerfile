FROM ubuntu:latest

RUN apt-get update

COPY ./src/index.html /usr/local/bin

ENV NAME="I am docker image 1"

CMD ["echo", "$NAME"]
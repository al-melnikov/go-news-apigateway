FROM golang:alpine
RUN mkdir /news
ADD . /news/
WORKDIR /news/cmd
RUN go build -o news .
CMD ["./news"]
FROM golang:alpine
RUN mkdir /rss-agg
ADD . /rss-agg/
WORKDIR /rss-agg/cmd
RUN go build -o rss-agg .
CMD ["./rss-agg"]
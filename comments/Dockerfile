FROM golang:alpine
RUN mkdir /comments
ADD . /comments/
WORKDIR /comments/cmd
RUN go build -o comments .
CMD ["./comments"]
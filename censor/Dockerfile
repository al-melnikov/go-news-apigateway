FROM golang:alpine
RUN mkdir /censor
ADD . /censor/
WORKDIR /censor/cmd
RUN go build -o censor .
CMD ["./censor"]
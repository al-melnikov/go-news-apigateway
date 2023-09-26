FROM golang:alpine
RUN mkdir /api-gateway
ADD . /api-gateway/
WORKDIR /api-gateway/cmd
RUN go build -o api-gateway .
CMD ["./api-gateway"]
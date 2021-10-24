FROM golang:1.16

ENV HOST="0.0.0.0"
ENV PORT="8080"
ENV GO111MODULE="on"

WORKDIR /usr/app
ADD . .

RUN go test ./...

WORKDIR /usr/app/cmd

RUN go build -o server main.go

EXPOSE 8080

CMD ["./server"]
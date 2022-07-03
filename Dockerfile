FROM golang:1.18.2

WORKDIR /src/

RUN mkdir ws

WORKDIR /ws

COPY . .

RUN go build -o bin/bin cmd/main.go

EXPOSE 8000

CMD ["./bin/bin"]
FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /gobinaries cmd/server/main.go

EXPOSE 9282

CMD [ "/gobinaries"]
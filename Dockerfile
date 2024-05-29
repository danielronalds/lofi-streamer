FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./

ADD cmd/ ./cmd
ADD templates/ ./templates
COPY streams.json .

RUN go mod tidy

RUN go build -o ./lofi-streamer ./cmd

EXPOSE 8080

CMD [ "./lofi-streamer" ]


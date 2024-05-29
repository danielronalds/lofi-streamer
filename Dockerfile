# The build stage
FROM golang:1.22.3-bookworm as builder

WORKDIR /app

COPY src/ ./src/
COPY go.mod ./

RUN go mod tidy

RUN go build -o ./lofi-streamer ./src

# The run stage
FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/lofi-streamer .
COPY streams.json .

EXPOSE 8080

CMD [ "./lofi-streamer" ]


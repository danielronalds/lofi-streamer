# TODO: Slim down this image https://medium.com/@minhaz1217/smallest-docker-image-for-go-api-project-b204b1f41d4e
FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./

ADD *.go .
ADD templates/ ./templates
COPY streams.json .

RUN go mod tidy

RUN go build -o ./lofi-streamer .

EXPOSE 8080

CMD [ "./lofi-streamer" ]


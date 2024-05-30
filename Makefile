build:
	docker build -t lofi-streamer .

run:
	docker stop lofi-streamer
	docker rm lofi-streamer
	docker run -d --name "lofi-streamer" -p 8080:8080 lofi-streamer

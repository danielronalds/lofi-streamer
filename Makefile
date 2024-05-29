docker-container:
	docker build -t lofi-streamer .

run-docker:
	docker run -p 8080:8080 lofi-streamer

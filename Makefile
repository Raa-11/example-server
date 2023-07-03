dev:
	go run -mod=vendor . -port="3000"

dobu:
	docker build -t http-server -f docker/Dockerfile .

dodev:
	docker run -it -p 3000:3000 http-server:latest
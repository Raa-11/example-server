osv:
	osv-scanner --lockfile=go.mod

lint:
	golangci-lint run

dev:
	go run -mod=vendor . -port="3000"

build:
	go build -mod=vendor -o server .

test:
	go test ./... -v -cover

dobu:
	docker build -t http-server -f docker/Dockerfile .

dodev:
	docker run -it -p 3000:3000 http-server:latest
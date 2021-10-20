GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

dev:
	go run main.go

test:
	go test -cover ./...

build:
	go build main.go

build-service:
	docker run -d --name qr-generator \
		-p 5002:5002 \
		--restart=always \
		qr-generator

start-service:
	docker start qr-generator

stop-service:
	docker stop qr-generator

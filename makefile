run: build
	clear && ./tmp/bin/tgbot

build:
	go build -o ./tmp/bin/tgbot ./cmd/bot/main.go  

docker-test:
	sudo docker run -d --name redis-test -p 6379:6379 redis

docker-test-stop:
	sudo docker stop redis-test

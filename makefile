default: build
	clear && ./tmp/bin/tgbot

build:
	go build -o ./tmp/bin/tgbot ./cmd/bot/main.go  

docker-run:
	sudo docker-compose up

extract-text:
	go generate ./internal/translations/translations.go

save-translations:
	bash ./save-translations.sh

default: build
	clear && ./tmp/bin/tgbot

build:
	go build -o ./tmp/bin/tgbot ./cmd/bot/main.go  

docker-run:
	sudo docker-compose up

redis-cli-language-sessions:
	sudo docker exec -it tgbot-lang-sessions redis-cli
redis-cli-chat-states:
	sudo docker exec -it tgbot-chat-states redis-cli

mariadb-cli:
	sudo docker exec -it tgbot-main-db mariadb -u root -padmin tgbot_db

extract-text:
	go generate ./internal/translations/translations.go

save-translations:
	bash ./save-translations.sh

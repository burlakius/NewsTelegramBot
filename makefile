default: build
	clear && ./tmp/bin/tgbot

build:
	go build -o ./tmp/bin/tgbot ./cmd/bot/main.go  

docker-run:
	sudo docker-compose up

redis-cli-language-sessions:
	sudo docker exec -it news_telegram_bot-redis-language-sessions-1 redis-cli
redis-cli-chat-states:
	sudo docker exec -it news_telegram_bot-redis-chat-states-1 redis-cli

mariadb-cli:
	sudo docker exec -it news_telegram_bot-tgbot-mariadb-1 mariadb -u root -padmin tgbot_db

extract-text:
	go generate ./internal/translations/translations.go

save-translations:
	bash ./save-translations.sh

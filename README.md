### Features

- Posting and reading the news
- Communication between users and admins by asking questions
- Multilingual interface(currently ukr and eng)

## Requirements
####For deployment
- docker
- docker-compose

####For editing
- Golang compiler and tools:
	- golang.org/x/text/cmd/gotext

- Redis
- Mariadb
- Make(optional)

### How does the bot work?
The bot makes a request for updates every interval (15 ms by default).

	Routing of updates is handled by Router. It receives and sends updates to Dispatchers one by one (currently 4 of them are created and each is running in a separate goroutine).
	Dispatcher processes updates by main type (message, callbackQuery, etc.). At the moment we have implemented processing of messages, edited messages and callbackQueries. When it matches the desired type, the dispatcher passes the content of the update to Handlers.
	Handler has a list of filters in it and a callback function that is called when the update matches the filters.


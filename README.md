### Features

- **News Management:** Easily post and read news updates within the system.
- **User-Admin Communication:** Seamless interaction between users and admins through a question-asking feature.
- **Multilingual Interface:** Currently supports Ukrainian (ukr) and English (eng) languages.

### Requirements

#### For Deployment

- Docker
- Docker-compose

#### For Editing

- **Golang Compiler and Tools:**
  - golang.org/x/text/cmd/gotext

- **Database and Caching:**
  - Redis
  - MariaDB

- **Optional:**
  - Make (for convenience)

### How does the bot work?

The bot operates on a scheduled interval (15 ms by default) for requesting updates.

#### Update Routing

- **Router:** Manages the routing of updates. It receives and sequentially sends updates to Dispatchers. Currently, four Dispatchers are created, each running in a separate goroutine.

#### Update Processing

- **Dispatcher:** Processes updates based on their main type (message, callbackQuery, etc.). Implemented handling for messages, edited messages, and callbackQueries. Upon matching the desired type, the dispatcher forwards the update content to Handlers.

#### Handling Updates

- **Handler:** Contains a list of filters and a callback function. The callback function is invoked when the update matches the specified filters. This allows for flexible and customizable handling of different types of updates.

package mariadb

// Returns admin chats ID slice from database
func getAdminsChatsFromDB() []int64 {
	rows, err := botStorage.mainDB.Query("SELECT chat_id FROM chat_admins")
	if err != nil {
		return []int64{}
	}
	defer rows.Close()

	result := make([]int64, 0, 5)
	var chat_id int64
	for rows.Next() {
		err = rows.Scan(&chat_id)
		if err != nil {
			panic(err)
		}

		result = append(result, chat_id)
	}

	return result
}

// Returns admin chats ID from bot storage
func GetAdminsChats() []int64 {
	return botStorage.adminsChats
}

// Add new chatID in database and bot storage
func AddNewAdminChat(chat_id int64) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO chat_admins (chat_id) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chat_id)

	botStorage.adminsChats = append(botStorage.adminsChats, chat_id)

	return err
}

// Check to see if there is a chat in bot storage
func IsAdminChat(chatID int64) bool {
	for _, chat := range botStorage.adminsChats {
		if chatID == chat {
			return true
		}
	}

	return false
}

type Question struct {
	UserID    int64
	FirstName string
	LastName  string
	Username  string
	ChatID    int64
	MessageID int
}

// Returns first question as Question struct from database
func GetQuestion() (*Question, error) {
	var question Question

	row := botStorage.mainDB.QueryRow("SELECT user_questions.user_id, users.first_name, users.last_name, users.username, user_questions.question_chat_id, user_questions.question_message_id FROM user_questions JOIN users ON user_questions.user_id = users.user_id")
	err := row.Scan(&question.UserID, &question.FirstName, &question.LastName, &question.Username, &question.ChatID, &question.MessageID)

	return &question, err
}

// Delete answer message from database
func DeleteAnswerMessage(chatID int64, messageID int) error {
	_, err := botStorage.mainDB.Exec("DELETE FROM admin_answers WHERE answer_chat_id = ? and answer_message_id = ?", chatID, messageID)

	return err
}

// Delete first question message from database
func DeleteQuestionFirstMessage() error {
	_, err := botStorage.mainDB.Exec("DELETE FROM user_questions LIMIT 1")

	return err
}

// Add answer message in database
func SaveAnswerMessage(answer_chat_id int64, answer_message_id int) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO admin_answers (answer_chat_id, answer_message_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(answer_chat_id, answer_message_id)

	return err
}

// Returns all answer messages(chatID, messageID)
func GetAllAnswerMessages() (map[int64]int, error) {
	result := make(map[int64]int)

	rows, err := botStorage.mainDB.Query("SELECT answer_chat_id, answer_message_id FROM admin_answers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		chatID    int64
		messageID int
	)
	for rows.Next() {
		err := rows.Scan(&chatID, &messageID)
		if err != nil {
			return result, err
		}

		result[chatID] = messageID
	}

	if err := rows.Err(); err != nil {
		return result, err
	}

	return result, nil
}

// Delete all answer messages from database
func DeleteAllAnswerMessages() error {
	_, err := botStorage.mainDB.Exec("DELETE FROM admin_answers")

	return err
}

// Add news message in database
func AddNewsMessage(chatID int64, messageID int) {

}

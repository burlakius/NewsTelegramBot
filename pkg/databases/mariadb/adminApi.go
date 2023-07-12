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

// Insert new chatID in database and bot storage
func SetAdminChat(chat_id int64) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO chat_admins (chat_id) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chat_id)
	if err != nil {
		panic(err)
	}

	botStorage.adminsChats = append(botStorage.adminsChats, chat_id)

	return nil
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

// Returns first question(chatID, messageID) from database
func GetQuestion() (int64, int, error) {
	var (
		chatID    int64
		messageID int
	)

	row := botStorage.mainDB.QueryRow("SELECT question_chat_id, question_message_id FROM user_questions LIMIT 1")
	err := row.Scan(&chatID, &messageID)
	if err != nil {
		return 0, 0, err
	}

	return chatID, messageID, nil
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

// Insert answer message in database
func InsertAnswerMessage(answer_chat_id int64, answer_message_id int) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO admin_answers (answer_chat_id, answer_message_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(answer_chat_id, answer_message_id)
	if err != nil {
		panic(err)
	}

	return nil
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

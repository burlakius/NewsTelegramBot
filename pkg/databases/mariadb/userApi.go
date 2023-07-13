package mariadb

// Insert new user data in database
func AddNewUser(userID int64, firstName, lastName, userName string) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO users (user_id, first_name, last_name, username) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, firstName, lastName, userName)

	return err
}

// Insert user question message data in database
func AddQuestionMessage(userID, questionChatID int64, questionMessageID int) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO user_questions (user_id, question_chat_id, question_message_id) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, questionChatID, questionMessageID)

	return err
}

// Delete user question message from database
func DeleteQuestionMessage(chatID int64, messageID int) error {
	_, err := botStorage.mainDB.Exec("DELETE FROM user_questions WHERE question_chat_id = ? and question_message_id = ?", chatID, messageID)

	return err
}

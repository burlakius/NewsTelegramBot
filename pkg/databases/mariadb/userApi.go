package mariadb

func SetQuestion(question_chat_id int64, question_message_id int) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT INTO user_questions (question_chat_id, question_message_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(question_chat_id, question_message_id)
	if err != nil {
		panic(err)
	}

	return nil
}

func DeleteQuestionMessage(chatID int64, messageID int) error {
	_, err := botStorage.mainDB.Exec("DELETE FROM user_questions WHERE question_chat_id = ? and question_message_id = ?", chatID, messageID)

	return err
}

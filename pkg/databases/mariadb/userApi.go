package mariadb

// Insert new user data in database
func AddNewUser(userID int64, firstName, lastName, userName string) error {
	botStorage.mu.Lock()
	defer botStorage.mu.Unlock()
	stmt, err := botStorage.mainDB.Prepare("INSERT IGNORE INTO users (user_id, first_name, last_name, username) VALUES (?, ?, ?, ?)")
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

// Returns all news for user from database
func GetUserNews(userID int64) ([]News, error) {
	result := make([]News, 0, 10)

	var news News
	rows, err := botStorage.mainDB.Query("SELECT news_id, news_chat_id, news_message_id, DATE(publication_date) FROM news WHERE news_id > (SELECT last_news_id FROM users WHERE user_id = ?)", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&news.NewsID, &news.ChatID, &news.MessageID, &news.PublicationDate)
		if err != nil {
			return nil, err
		}

		result = append(result, news)
	}

	return result, nil
}

func UpdateUserLastNews(lastNewsID int, userID int64) error {
	_, err := botStorage.mainDB.Exec("UPDATE users SET last_news_id = ? WHERE user_id = ?", lastNewsID, userID)
	return err
}

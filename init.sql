CREATE DATABASE IF NOT EXISTS tgbot_db;
USE tgbot_db;

-- News types table
CREATE TABLE IF NOT EXISTS news_types (
  news_type_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- News table
CREATE TABLE IF NOT EXISTS news (
  news_id INT AUTO_INCREMENT PRIMARY KEY,
  news_type_id INT NOT NULL,
  news_chat_id BIGINT NOT NULL,
  news_message_id INT NOT NULL,
  state ENUM('unhidden', 'hidden'),
  publication_date DATETIME DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_news_types
    FOREIGN KEY (news_type_id)
    REFERENCES news_types(news_type_id)
    ON DELETE CASCADE
);

-- Users table
CREATE TABLE IF NOT EXISTS users (
  user_id BIGINT PRIMARY KEY,
  news_type_id INT DEFAULT 1,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64) NOT NULL,
  username VARCHAR(32) NOT NULL,
  CONSTRAINT fk_users_news_types
    FOREIGN KEY (news_type_id)
    REFERENCES news_types(news_type_id)
    ON DELETE CASCADE
);

-- User questions table
CREATE TABLE IF NOT EXISTS user_questions (
  question_id INT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT,
  question_chat_id BIGINT NOT NULL,
  question_message_id INT NOT NULL,
  CONSTRAINT fk_user_questions_users
    FOREIGN KEY (user_id)
    REFERENCES users(user_id)
    ON DELETE CASCADE
);

-- Message table of answers to the question
CREATE TABLE IF NOT EXISTS admin_answers (
  answer_id INT AUTO_INCREMENT PRIMARY KEY,
  answer_chat_id BIGINT NOT NULL,
  answer_message_id INT NOT NULL
);

-- Admin chats table
CREATE TABLE IF NOT EXISTS chat_admins (
  admin_chat_id INT AUTO_INCREMENT PRIMARY KEY,
  chat_id BIGINT NOT NULL,
  UNIQUE (chat_id)
);

INSERT INTO news_types (name) VALUES ('regular'), ('important');

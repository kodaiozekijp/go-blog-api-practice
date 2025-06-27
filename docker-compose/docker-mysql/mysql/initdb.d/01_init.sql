DROP DATABASE IF EXISTS blog_api_db;
CREATE DATABASE blog_api_db;
USE blog_api_db;

# 記事データを格納するためのテーブル
CREATE TABLE IF NOT EXISTS articles (
  article_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  contents TEXT NOT NULL,
  username VARCHAR(100) NOT NULL,
  nice INTEGER NOT NULL,
  created_at DATETIME
);

# コメントデータを格納するためのテーブル
CREATE TABLE IF NOT EXISTS comments (
  comment_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  article_id INTEGER UNSIGNED NOT NULL,
  message TEXT NOT NULL,
  created_at DATETIME,
  FOREIGN KEY (article_id) REFERENCES articles(article_id)
);
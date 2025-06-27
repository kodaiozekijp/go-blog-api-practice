# 記事データを格納するためのテーブル
CREATE TABLE IF NOT EXISTS articles (
  article_id INTEGER AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  contents TEXT NOT NULL,
  username VARCHAR(100) NOT NULL,
  nice INTEGER NOT NULL,
  created_at DATETIME,
  CONSTRAINT article_id_check CHECK(article_id > 0)
);

# コメントデータを格納するためのテーブル
CREATE TABLE IF NOT EXISTS comments {
  comment_id INTEGER AUTO_INCREMENT PRIMARY KEY,
  article_id INTEGER NOT NULL,
  message TEXT NOT NULL,
  created_at DATETIME,
  FOREIGN KEY (article_id) REFERENCES articles(article_id)
}
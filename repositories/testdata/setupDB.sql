-- 記事データを格納するためのテーブル
CREATE TABLE IF NOT EXISTS articles (
  article_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  contents TEXT NOT NULL,
  author VARCHAR(100) NOT NULL,
  nice INTEGER NOT NULL,
  created_at DATETIME
);

-- コメントデータを格納するためのテーブル
CREATE TABLE IF NOT EXISTS comments (
  comment_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  article_id INTEGER UNSIGNED NOT NULL,
  message TEXT NOT NULL,
  created_at DATETIME,
  FOREIGN KEY (article_id) REFERENCES articles(article_id)
);

-- 記事データ2つ
INSERT INTO articles (title, contents, author, nice, created_at) VALUES 
  ('1st Post', 'This is my first blog', 'kodai', 1, now());

INSERT INTO articles (title, contents, author, nice) VALUES 
  ('2nd Post', 'Second blog post', 'kodai', 2);

-- コメントデータ2つ
INSERT INTO comments (article_id, message, created_at) VALUES 
  (1, '1st comment', now());

INSERT INTO comments (article_id, message) VALUES
  (1, '2nd comment');
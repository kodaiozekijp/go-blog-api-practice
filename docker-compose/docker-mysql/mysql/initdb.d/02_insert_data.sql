# 記事データ2つ
INSERT INTO articles (title, contents, author, nice, created_at) VALUES 
  ('1st Post', 'This is my first blog', 'kodai', 1, now());

INSERT INTO articles (title, contents, author, nice) VALUES 
  ('2nd Post', 'Second blog post', 'kodai', 2);

# コメントデータ2つ
INSERT INTO comments (article_id, message, created_at) VALUES 
  (1, '1st comment', now());

INSERT INTO comments (article_id, message) VALUES
  (1, '2nd comment');
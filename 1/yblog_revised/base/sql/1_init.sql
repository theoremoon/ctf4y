CREATE TABLE yblog.posts (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  private BOOLEAN NOT NULL DEFAULT 0,
  private_key TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
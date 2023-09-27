CREATE TABLE IF NOT EXISTS author (
  authorID      SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  realName      TINYTEXT NOT NULL,
  userName      TINYTEXT NOT NULL,
  createdDate   DATETIME NOT NULL
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS articles (
  articleID     SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title         TINYTEXT NOT NULL,
  contents      TEXT,
  createdDate   DATETIME NOT NULL,
  updatedDate   DATETIME,
  authorID      SMALLINT UNSIGNED NOT NULL,
  CONSTRAINT `fk_article_author`
    FOREIGN KEY (authorID) REFERENCES author (authorID)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT
) ENGINE = InnoDB;

INSERT INTO author
  (realName, userName, createdDate)
VALUES
  ('System', 'system', NOW());

INSERT INTO articles
  (title, contents, createdDate, authorID)
VALUES
  ('Welcome', 'Welcome to the Wiki!', NOW(), LAST_INSERT_ID());


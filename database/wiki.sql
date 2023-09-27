CREATE TABLE IF NOT EXISTS author (
  authorID      SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  realName      TINYTEXT NOT NULL,
  userName      TINYTEXT NOT NULL,
  createdDate   DATETIME NOT NULL
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS articles (
  articleID     SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title         TINYTEXT NOT NULL,
  body      TEXT,
  createdDate   DATETIME NOT NULL,
  updatedDate   DATETIME NOT NULL,
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
  (title, body, createdDate, updatedDate, authorID)
VALUES
  ('Lorem Ipsum', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent dictum, arcu eget consectetur interdum, risus magna viverra nibh, at placerat lacus est in dui. Aenean quis nisi id justo laoreet aliquam. Nunc a lectus interdum, rutrum massa quis, commodo nisl. Vestibulum ullamcorper, nunc id rutrum tempus, mi urna accumsan ex, ac faucibus elit lorem dictum nulla. Etiam dapibus, nulla non consectetur imperdiet, lacus lacus viverra lacus, dignissim gravida nulla odio eget libero. Suspendisse ac ex fermentum, vestibulum lorem eget, porta sapien. Vestibulum tempor enim vel elit egestas, eu hendrerit libero imperdiet. ', NOW(), NOW(), LAST_INSERT_ID());


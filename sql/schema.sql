CREATE TABLE users
(
    id VARCHAR(40) NOT NULL,
    username VARCHAR(255) NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
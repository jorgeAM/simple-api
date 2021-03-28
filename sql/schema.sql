CREATE TABLE users
(
    id VARCHAR(40) NOT NULL,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
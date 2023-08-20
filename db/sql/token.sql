CREATE TABLE IF NOT EXISTS user_token (
token varchar(255) NOT NULL PRIMARY KEY,
username varchar(32) NOT NULL UNIQUE,
user_id int(32) NOT NULL INDEX
);
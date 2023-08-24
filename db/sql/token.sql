CREATE TABLE IF NOT EXISTS `user_tokens` (
    `token` VARCHAR(255) NOT NULL,
    `username` VARCHAR(32) NOT NULL UNIQUE,
    `user_id` INT NOT NULL,
    PRIMARY KEY ( `token` ),
    INDEX `user_id_index` ( `user_id` )
);
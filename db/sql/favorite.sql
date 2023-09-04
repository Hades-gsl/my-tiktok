CREATE TABLE IF NOT EXISTS `favorites` (
    `id` INT AUTO_INCREMENT,
    `created_at` DATETIME,
    `updated_at` DATETIME,
    `deleted_at` DATETIME,
    `user_id` INT,
    `video_id` INT,
    PRIMARY KEY ( `id` )
);

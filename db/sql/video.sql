CREATE TABLE videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    user_id INT,
    title VARCHAR(32),
    file_addr VARCHAR(255),
    cover_addr VARCHAR(255),
    favorite_count INT DEFAULT 0,
    comment_count INT DEFAULT 0
);

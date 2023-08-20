CREATE TABLE IF NOT EXISTS videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    user_id INT,
    title VARCHAR(32),
    file_addr VARCHAR(255),
    cover_addr VARCHAR(255),
);

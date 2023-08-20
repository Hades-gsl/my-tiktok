CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    user_name VARCHAR(32),
    pass_word VARCHAR(32),
    avatar VARCHAR(255),
    background_image VARCHAR(255),
    signature VARCHAR(255),
);

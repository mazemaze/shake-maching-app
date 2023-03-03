CREATE TABLE IF NOT EXISTS users (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(16),
    password VARCHAR(40),
    created_at TIMESTAMP,
    deleted_at TIMESTAMP
)
CREATE TABLE IF NOT EXISTS chat_rooms (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    user_1 INTEGER,
    user_2 INTEGER,
    status INTEGER,
    created_at TIMESTAMP
)
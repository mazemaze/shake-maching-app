CREATE TABLE IF NOT EXISTS chat_rooms (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    room_id INTEGER,
    user_id INTEGER,
    score INTEGER,
    created_at TIMESTAMP
)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL
);

INSERT INTO users (username, password_hash) VALUES
('Sergey', '$2a$10$o5gF65.0CTWA4s/35uU3yeIMeCcXHKGkAL.vhoAhgA1mD.yCXJKiS'),  --12345
('Jessie', 'Diggins'),
('Gofer', 'golang');

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    userid INT REFERENCES users(id),
    note TEXT NOT NULL,
    date_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO notes (userid, note) VALUES
(1, 'First note'),
(2, 'Second note'),
(3, 'Third note');

CREATE INDEX IF NOT EXISTS idx_notes_userid ON notes(userid);



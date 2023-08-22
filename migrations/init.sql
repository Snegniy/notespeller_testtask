CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL
);

INSERT INTO users (username, password_hash) VALUES
('Sergey', '$2a$10$hemvLnUXcbRjX9wcNnqrZeah6QwWIYl2oQFKv2hzx8OHoSuftzbfa'),  --12345
('Jessie', '$2a$10$19YtWEKTW6Ior2ggM9.shuqOjnHY.6irrwsy4CdF7L0cNF52VCRpS'), --Diggins
('Gofer', '$2a$10$DTpHHK2u3CDkPrZxpBu/9uX3K/YG0RPxNTTb8d4vOAFlt2Gy4TioS'); --golang

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    userid INT REFERENCES users(id),
    note TEXT NOT NULL,
    date_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO notes (userid, note) VALUES
(1, 'First note'),
(1, 'I like it!'),
(2, 'Second note'),
(2, 'I LOVE SKI'),
(3, 'Third note'),
(3, 'Golang is cool!');

CREATE INDEX IF NOT EXISTS notes_userid_idx ON notes(userid);



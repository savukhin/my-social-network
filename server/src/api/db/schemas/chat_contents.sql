CREATE TABLE IF NOT EXISTS chat_contents
(
    id SERIAL NOT NULL,
    chat_id INTEGER NOT NULL,
    content_id INTEGER NOT NULL,

    CONSTRAINT chat_contents_PK PRIMARY KEY(id)
);

CREATE TABLE user_to_chat
(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    chat_id INTEGER NOT NULL,

    CONSTRAINT user_to_chat_PK PRIMARY KEY(id)
);

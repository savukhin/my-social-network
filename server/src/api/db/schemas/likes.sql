CREATE TABLE IF NOT EXISTS likes
(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    content_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT Like_PK PRIMARY KEY(id)
);

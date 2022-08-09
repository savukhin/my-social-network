CREATE TABLE likes
(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    content_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL SET DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT Like_PK PRIMARY KEY(id)
);

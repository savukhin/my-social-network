CREATE TABLE IF NOT EXISTS friendships
(
    id SERIAL NOT NULL,
    user1_id INTEGER NOT NULL,
    user2_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT Friendship_PK PRIMARY KEY(id)
)
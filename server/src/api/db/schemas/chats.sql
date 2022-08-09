CREATE TABLE chats 
(
    id SERIAL NOT NULL,
    title VARCHAR(200) NOT NULL,
    photo_id INTEGER,
    created_at TIMESTAMP NOT NULL SET DEFAULT now(),
    deleted_at TIMESTAMP,
    
    CONSTRAINT Chat_PK PRIMARY KEY(id)
);

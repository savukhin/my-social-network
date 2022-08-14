CREATE TABLE IF NOT EXISTS  chats 
(
    id SERIAL NOT NULL,
    title VARCHAR(200) NOT NULL,
    photo_id INTEGER,
    created_at TIMESTAMP DEFAULT now(),
    udpated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,
    
    CONSTRAINT Chat_PK PRIMARY KEY(id)
);

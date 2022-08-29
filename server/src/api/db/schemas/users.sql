CREATE TABLE IF NOT EXISTS users
(
    id SERIAL NOT NULL,
    username VARCHAR(200) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    status VARCHAR(100),
    city VARCHAR(100),
    birth_date DATE,
    avatar_id INTEGER,
    isOnline boolean DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT User_PK PRIMARY KEY(id),
    CONSTRAINT Email_UC UNIQUE(email),
    CONSTRAINT Username_UC UNIQUE(username)
);

CREATE TYPE content_types AS ENUM ('post', 'message', 'comment', 'photo');

CREATE TABLE contents (
    id SERIAL NOT NULL,
    filepath VARCHAR(200) NOT NULL,
    content_type content_types NOT NULL,
    parent_content_id INTEGER,
    user_id INTEGER,
    order INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL SET DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT Content_PK PRIMARY KEY(id)
);

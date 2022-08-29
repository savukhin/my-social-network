CREATE TABLE IF NOT EXISTS contents (
    id SERIAL NOT NULL,
    filepath VARCHAR(200) NOT NULL,
    content_type content_types NOT NULL,
    parent_id INTEGER,
    user_id INTEGER NOT NULL,
    attach_order INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT Content_PK PRIMARY KEY(id)
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'content_types') THEN
        CREATE TYPE content_types AS ENUM 
        (
            'post', 
            'message', 
            'comment', 
            'photo'
        );
    END IF;
END$$;

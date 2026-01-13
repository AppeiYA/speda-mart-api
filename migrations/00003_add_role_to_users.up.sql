DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'user_role'
    ) THEN
        CREATE TYPE user_role AS ENUM ('admin', 'user');
    END IF;
END $$;

ALTER TABLE IF EXISTS users
ADD COLUMN role user_role NOT NULL DEFAULT 'user';
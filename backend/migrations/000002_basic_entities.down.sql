DROP TABLE IF EXISTS post_hashtags;
DROP TABLE IF EXISTS hashtags;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS follows;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS posts;

DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;

ALTER TABLE users DROP CONSTRAINT IF EXISTS users_username_check;

ALTER TABLE users
    DROP COLUMN IF EXISTS username,
    DROP COLUMN IF EXISTS display_name,
    DROP COLUMN IF EXISTS status,
    DROP COLUMN IF EXISTS avatar_url,
    DROP COLUMN IF EXISTS bio,
    DROP COLUMN IF EXISTS email_verified,
    DROP COLUMN IF EXISTS password_reset_token,
    DROP COLUMN IF EXISTS password_reset_expires_at;

ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);

INSERT INTO users(id, email, role)
VALUES
    ('8794e589-0ddb-43ce-9f92-16faafcf4ee4', 'test.user@email.com',  'user'),
    ('249be7cf-d419-4c54-97f2-d04107806e36', 'test.admin@email.com', 'admin')
ON CONFLICT (id) DO NOTHING;

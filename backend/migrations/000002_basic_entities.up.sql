DELETE FROM users;

ALTER TABLE users
    ADD COLUMN username                  VARCHAR(15)  NOT NULL,
    ADD COLUMN display_name              VARCHAR(30),
    ADD COLUMN status                    VARCHAR(100),
    ADD COLUMN avatar_url                TEXT,
    ADD COLUMN bio                       TEXT,
    ADD COLUMN email_verified            BOOLEAN      NOT NULL DEFAULT FALSE,
    ADD COLUMN password_reset_token      VARCHAR(255),
    ADD COLUMN password_reset_expires_at TIMESTAMP;

ALTER TABLE users
    ADD CONSTRAINT users_username_check
        CHECK (char_length(username) BETWEEN 4 AND 15
               AND username ~ '^[a-zA-Z0-9_]+$');

ALTER TABLE users DROP CONSTRAINT IF EXISTS users_email_key;
CREATE UNIQUE INDEX idx_users_email    ON users(email);
CREATE UNIQUE INDEX idx_users_username ON users(username);

CREATE TABLE posts (
    id            VARCHAR(36)  PRIMARY KEY,
    user_id       VARCHAR(36)  NOT NULL,
    parent_id     VARCHAR(36),
    content       TEXT         NOT NULL CHECK (char_length(content) <= 280),
    image_url     TEXT,
    likes_count   INTEGER      NOT NULL DEFAULT 0,
    deleted_at    TIMESTAMP,
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    search_vector TSVECTOR,

    CONSTRAINT fk_posts_user   FOREIGN KEY (user_id)   REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_posts_parent FOREIGN KEY (parent_id) REFERENCES posts(id) ON DELETE SET NULL,

    CONSTRAINT posts_image_only_for_root
        CHECK (parent_id IS NULL OR image_url IS NULL)
);

CREATE INDEX idx_replies    ON posts(parent_id, created_at);
CREATE INDEX idx_feed       ON posts(created_at) WHERE parent_id IS NULL AND deleted_at IS NULL;
CREATE INDEX idx_user_posts ON posts(user_id, created_at);
CREATE INDEX idx_search     ON posts USING GIN (search_vector);

CREATE TABLE likes (
    user_id    VARCHAR(36) NOT NULL,
    post_id    VARCHAR(36) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_likes      PRIMARY KEY (user_id, post_id),
    CONSTRAINT fk_likes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_likes_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_likes_post ON likes(post_id);

CREATE TABLE follows (
    follower_id  VARCHAR(36) NOT NULL,
    following_id VARCHAR(36) NOT NULL,
    created_at   TIMESTAMP   NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_follows           PRIMARY KEY (follower_id, following_id),
    CONSTRAINT fk_follows_follower  FOREIGN KEY (follower_id)  REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_follows_following FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT follows_no_self      CHECK (follower_id <> following_id)
);

CREATE INDEX idx_follows_following ON follows(following_id);
CREATE INDEX idx_follows_follower  ON follows(follower_id);

CREATE TABLE notifications (
    id           VARCHAR(36) PRIMARY KEY,
    recipient_id VARCHAR(36) NOT NULL,
    actor_id     VARCHAR(36) NOT NULL,
    type         VARCHAR(20) NOT NULL CHECK (type IN ('follow', 'like', 'comment')),
    post_id      VARCHAR(36),
    comment_id   VARCHAR(36),
    is_read      BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP   NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_notif_recipient FOREIGN KEY (recipient_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_notif_actor     FOREIGN KEY (actor_id)     REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_notif_post      FOREIGN KEY (post_id)      REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_notif_comment   FOREIGN KEY (comment_id)   REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_notif_unread         ON notifications(recipient_id, is_read);
CREATE INDEX idx_notif_recipient_date ON notifications(recipient_id, created_at);

CREATE TABLE hashtags (
    id         VARCHAR(36) PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_hashtags_name ON hashtags(name);

CREATE TABLE post_hashtags (
    post_id    VARCHAR(36) NOT NULL,
    hashtag_id VARCHAR(36) NOT NULL,

    CONSTRAINT pk_post_hashtags      PRIMARY KEY (post_id, hashtag_id),
    CONSTRAINT fk_post_hashtags_post FOREIGN KEY (post_id)    REFERENCES posts(id)    ON DELETE CASCADE,
    CONSTRAINT fk_post_hashtags_tag  FOREIGN KEY (hashtag_id) REFERENCES hashtags(id) ON DELETE CASCADE
);

CREATE INDEX idx_post_hashtags_tag ON post_hashtags(hashtag_id);

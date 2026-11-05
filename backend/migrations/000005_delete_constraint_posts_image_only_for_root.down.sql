ALTER TABLE posts
    ADD CONSTRAINT posts_image_only_for_root
        CHECK (parent_id IS NULL OR image_url IS NULL);

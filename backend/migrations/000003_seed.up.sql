INSERT INTO users(id, email, password, role, username, display_name, email_verified) VALUES
    ('8794e589-0ddb-43ce-9f92-16faafcf4ee4', 'test.user@email.com',  '$2a$14$mnb3B/vx4JyPsht03kwCJ.bLRiAYFwBxX9C.R/HH8Z4YvV.OJwgIa', 'user',  'username', 'Test User', TRUE),
    ('249be7cf-d419-4c54-97f2-d04107806e36', 'test.admin@email.com', '$2a$14$mnb3B/vx4JyPsht03kwCJ.bLRiAYFwBxX9C.R/HH8Z4YvV.OJwgIa', 'admin', 'adminame', 'Admin',     TRUE);

INSERT INTO posts(id, user_id, parent_id, content) VALUES
    ('p1', '8794e589-0ddb-43ce-9f92-16faafcf4ee4', NULL, 'Первый пост'),
    ('p2', '249be7cf-d419-4c54-97f2-d04107806e36', NULL, 'Второй пост'),
    ('c1', '249be7cf-d419-4c54-97f2-d04107806e36', 'p1', 'Комментарий к первому посту'),
    ('c2', '8794e589-0ddb-43ce-9f92-16faafcf4ee4', 'c1', 'Ответ на комментарий');

INSERT INTO follows(follower_id, following_id) VALUES
    ('8794e589-0ddb-43ce-9f92-16faafcf4ee4', '249be7cf-d419-4c54-97f2-d04107806e36');

INSERT INTO hashtags(id, name) VALUES
    ('h1', 'postgresql'),
    ('h2', 'dbml');

INSERT INTO post_hashtags(post_id, hashtag_id) VALUES
    ('p1', 'h1'),
    ('p1', 'h2');

DELETE FROM post_hashtags
WHERE (post_id, hashtag_id) IN (
    ('p1', 'h1'),
    ('p1', 'h2')
);

DELETE FROM hashtags
WHERE id IN ('h1', 'h2');

DELETE FROM follows
WHERE follower_id  = '8794e589-0ddb-43ce-9f92-16faafcf4ee4'
  AND following_id = '249be7cf-d419-4c54-97f2-d04107806e36';

DELETE FROM posts
WHERE id IN ('c2', 'c1');

DELETE FROM posts
WHERE id IN ('p1', 'p2');

DELETE FROM users
WHERE id IN (
    '8794e589-0ddb-43ce-9f92-16faafcf4ee4',
    '249be7cf-d419-4c54-97f2-d04107806e36'
);

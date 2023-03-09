-- name: create
INSERT INTO users (id, username) VALUES ($1, $2);

-- name: update
UPDATE users SET
    name = :name,
    exp = :exp,
    level = :level,
    next_level_exp = :next_level_exp
WHERE id = :id;


-- name: get-by-id
SELECT id, name, exp, level, next_level_exp, created_at FROM users WHERE id = $1; 

-- name: delete
DELETE FROM users WHERE id = $1;

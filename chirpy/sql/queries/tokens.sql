-- name: CreateUserRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
  $1,
  NOW(),
  NOW(),
  $2,
  $3,
  NULL
) 
RETURNING *;
--

-- name: CheckToken :one
SELECT * FROM refresh_tokens
WHERE user_id = $1;
--

-- name: RevokeToken :exec
UPDATE refresh_tokens
SET updated_at = $1, revoked_at = $2
WHERE user_id = $3;
--

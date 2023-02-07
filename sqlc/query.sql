-- name: AddLink :exec
INSERT INTO links (link, shortened_link)
VALUES (?, ?);

-- name: GetFullLink :one
SELECT link
FROM links
WHERE shortened_link = (?)
ORDER BY link_id DESC
LIMIT 1;
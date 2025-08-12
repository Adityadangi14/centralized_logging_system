-- name: InsertParsedLog :one
INSERT INTO parsed_logs (
    timestamp,
    event_category,
    source_type,
    username,
    hostname,
    severity,
    raw_message
    is_blacklisted
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetParsedLogByID :one
SELECT * FROM parsed_logs
WHERE id = $1;

-- name: ListParsedLogs :many
SELECT * FROM parsed_logs
ORDER BY timestamp DESC
LIMIT $1 OFFSET $2;

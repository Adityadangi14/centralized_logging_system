-- name: ListLogsByService :many
SELECT * FROM parsed_logs
WHERE event_category = $1
ORDER BY timestamp DESC
LIMIT 100;

-- name: ListLogsBySeverity :many
SELECT * FROM parsed_logs
WHERE severity = $1
ORDER BY timestamp DESC
LIMIT 100;

-- name: ListLogsByServiceAndSeverity :many
SELECT * FROM parsed_logs
WHERE event_category = $1 AND severity = $2
ORDER BY timestamp DESC
LIMIT 100;

-- name: ListLogsByUsernameAndBlacklisted :many
SELECT * FROM parsed_logs
WHERE username = $1 AND is_blacklisted = $2
ORDER BY timestamp DESC
LIMIT 100;

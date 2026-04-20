-- name: CreateSubscription :one
INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSubscription :one
SELECT * FROM subscriptions
WHERE id = $1;

-- name: ListSubscriptions :many
SELECT * FROM subscriptions
ORDER BY id;

-- name: UpdateSubscription :one
UPDATE subscriptions
SET service_name = $2,
    price        = $3,
    start_date   = $4,
    end_date     = $5,
    updated_at   = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE id = $1;

-- name: SumSubscriptionCost :one
SELECT COALESCE(SUM(price), 0)::INTEGER AS total
FROM subscriptions
WHERE
    (sqlc.narg('user_id')::UUID IS NULL OR user_id = sqlc.narg('user_id'))
  AND (sqlc.narg('service_name')::TEXT IS NULL OR service_name = sqlc.narg('service_name'))
  AND start_date >= sqlc.arg('from_date')
  AND (end_date IS NULL OR end_date <= sqlc.arg('to_date'));
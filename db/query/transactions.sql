-- name: CreateTransaction :one
INSERT INTO transactions (
  id_user,
  total_value,
  category,
  description,
  is_expense
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateTransactionDetail :one
INSERT INTO transaction_detail (
  id_transaction,
  description,
  quantity,
  unit_value
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id_transaction = $1 LIMIT 1;

-- name: GetListTransactions :many
SELECT * FROM transactions
WHERE id_user = $1
ORDER BY transactions_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateTransaction :one
UPDATE transactions
SET total_value = $2,
  category = $3,
  description = $4,
  is_expense = $5
WHERE id_transaction = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id_transaction = $1;


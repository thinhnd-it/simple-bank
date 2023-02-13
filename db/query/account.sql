-- name: CreateAccount :one
insert into accounts (
  owner,
  balance,
  currency
) values (
  $1, $2, $3
) returning *;

-- name: GetAccount :one
select * from accounts
where id = $1 limit 1;

-- name: ListAccount :many
select * from accounts
order by id
limit $1
offset $2; 

-- name: UpdateAccountBalance :one
UPDATE accounts
SET balance = $2
where id = $1
returning *;

-- name: DeleteAuthor :exec
DELETE FROM accounts WHERE id = $1;
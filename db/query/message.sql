-- name: CreateCustomer :one
INSERT INTO customer (name, phone, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateProduct :one
INSERT INTO product (name, stock, price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateOrder :one
INSERT INTO "order" (customer_id, product_id, quantity,total_price)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateCustomer :many
UPDATE customer 
SET name =$2,
    phone= $3,
    email = $4
WHERE id = $1
RETURNING *;

-- name: UpdateProduct :many
UPDATE product 
SET name = $2,
    stock = $3,
    price = $4
WHERE id = $1
RETURNING *;

-- name: GetOrder :one
SELECT *
FROM "order"
WHERE id = $1;

-- name: GetCustomersByNameLike :many
SELECT *
FROM customer
WHERE name LIKE $1;

-- name: GetTotalRevenueForProduct :one
SELECT SUM(total_price)
FROM "order"
WHERE product_id = $1;

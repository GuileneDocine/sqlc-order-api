-- name: CreateCustomer :many
INSERT INTO "Customer" (name, phone, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateProduct :many
INSERT INTO "Product" (name, price, stock)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateOrder :one
INSERT INTO "Order" ( customer_id, product_id, price, quantity, order_status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM "Customer"
WHERE id = $1;

-- name: GetAllProducts :many
SELECT * FROM "Product"
ORDER BY name ASC;

-- name: DeleteCustomer :exec
DELETE FROM "Customer" WHERE id = $1;

-- name: UpdateByIdProduct :one
UPDATE "Product" 
SET price = $2, stock = $3
WHERE id = $1
RETURNING *;

-- name: GetOrderById :one
SELECT * FROM "Order"
WHERE id = $1;

-- name: GetProductById :many
 SELECT FROM "Product"
 WHERE id = $1;
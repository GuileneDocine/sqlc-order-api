



CREATE TABLE "customer" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "name" VARCHAR(36) NOT NULL,
  "phone" VARCHAR(100) NOT NULL,
  "email" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);

CREATE TABLE "product" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "name" VARCHAR(36) NOT NULL,
  "stock" VARCHAR(100) NOT NULL,
  "price" VARCHAR(36) NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);

CREATE TABLE "order" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "customer_id" VARCHAR(36) REFERENCES "customer"("id") ON DELETE CASCADE,
  "product_id" VARCHAR(100) REFERENCES "product"("id") ON DELETE CASCADE,
  "quantity" VARCHAR(100) NOT NULL,
  "total_price" DECIMAL(10, 2) NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);


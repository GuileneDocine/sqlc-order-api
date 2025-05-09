
CREATE TABLE IF NOT EXISTS "Customer"(
    "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
    "name" VARCHAR(36) NOT NULL,
    "phone" VARCHAR(36) NOT NULL UNIQUE,
    "email" VARCHAR(36) NOT NULL UNIQUE,
    "created_at" TIMESTAMP DEFAULT now()
);


CREATE TABLE IF NOT EXISTS "Product"(
    "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
    "name" VARCHAR(36) NOT NULL,
    "price" DECIMAL(10, 2) NOT NULL,
    "stock" INT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT now()
);



CREATE TABLE IF NOT EXISTS "Order"(
    "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
    "customer_id" VARCHAR(36) NOT NULL,
    "product_id" VARCHAR(36) NOT NULL,
    "price" DECIMAL(10, 2) NOT NULL,
    "quantity" INT NOT NULL CHECK (quantity > 0),
    "order_status" VARCHAR (20) DEFAULT 'PENDING',
    "order_DATE" TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_Customer FOREIGN KEY("customer_id") REFERENCES "Customer" ("id") ON DELETE CASCADE,
    CONSTRAINT fk_Product FOREIGN KEY("product_id") REFERENCES "Product" ("id") ON DELETE CASCADE
);

-- -- delete the thread column in the message table
-- ALTER TABLE "message" DROP COLUMN "thread";

-- -- add the thread id column to the message table
-- ALTER TABLE "message" ADD "thread_id" VARCHAR(36);

-- -- add foreign key constrain in the message table 
-- ALTER TABLE message ADD CONSTRAINT fk_thread FOREIGN KEY (thread_id)
-- REFERENCES thread (id);
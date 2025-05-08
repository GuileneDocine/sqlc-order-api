-- -- create the table
-- CREATE TABLE IF NOT EXISTS "thread"(
--     "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
--     "topic" VARCHAR(36),
--     "created_at" TIMESTAMP DEFAULT now()
-- );

-- -- delete the thread column in the message table
-- ALTER TABLE "message" DROP COLUMN "thread";

-- -- add the thread id column to the message table
-- ALTER TABLE "message" ADD "thread_id" VARCHAR(36);

-- -- add foreign key constrain in the message table 
-- ALTER TABLE message ADD CONSTRAINT fk_thread FOREIGN KEY (thread_id)
-- REFERENCES thread (id);
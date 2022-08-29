CREATE TABLE customer(
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "phone_number" bigint NOT NULL,
    "email_address" varchar NOT NULL,
    "password_hash" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE account(
    "id" bigserial PRIMARY KEY,
    "type" varchar NOT NULL,
    "customer_id" bigint,
    "balance" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);


CREATE TABLE txn(
    "id" bigserial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "transaction_type" varchar NOT NULL,
    "amount_type" varchar NOT NULL,
    "amount" bigint NOT NULL,
    "description" varchar,
    "corresponding_account_id" bigint,
    "current_balance" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

ALTER TABLE "account" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");
ALTER TABLE "txn" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
ALTER TABLE "txn" ADD FOREIGN KEY ("corresponding_account_id") REFERENCES "account" ("id");


CREATE INDEX idx_account_id ON "txn" ("account_id");
CREATE INDEX idx_corresponding_account_id ON "txn" ("corresponding_account_id");
CREATE INDEX idx_transaction_history ON "txn" ("account_id", "corresponding_account_id", "transaction_type", "created_at");

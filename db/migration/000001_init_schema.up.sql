CREATE TABLE "users" (
  "id_user" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id_transaction" bigserial PRIMARY KEY,
  "id_user" bigserial NOT NULL,
  "transactions_at" timestamp NOT NULL DEFAULT (now()),
  "total_value" bigint NOT NULL,
  "category" varchar NOT NULL,
  "description" varchar NOT NULL,
  "is_expense" bool NOT NULL
);

CREATE TABLE "transaction_detail" (
  "id_transaction_detail" bigserial PRIMARY KEY,
  "id_transaction" bigserial NOT NULL,
  "description" varchar,
  "quantity" integer,
  "unit_value" bigint,
  "transaction_detail_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "id_user" bigserial NOT NULL,
  "user_name" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "transaction_detail" ADD FOREIGN KEY ("id_transaction") REFERENCES "transactions" ("id_transaction");

ALTER TABLE "sessions" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");
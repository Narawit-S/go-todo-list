CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar NOT NULL,
  "encrypted_password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now ()',
  "updated_at" timestamp NOT NULL DEFAULT 'now ()'
);

CREATE TABLE "todos" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "user_id" int NOT NULL,
  "finished" boolean NOT NULL,
  "deadline" timestamp,
  "created_at" timestamp NOT NULL DEFAULT 'now ()',
  "updated_at" timestamp NOT NULL DEFAULT 'now ()'
);

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

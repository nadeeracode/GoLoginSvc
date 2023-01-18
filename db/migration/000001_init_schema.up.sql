CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "user_type" int NOT NULL DEFAULT(0),
  "active" int NOT NULL DEFAULT(0),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_accessed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

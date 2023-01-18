CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "hashed_password" varchar,
  "email" varchar,
  "user_type" int,
  "active" bool,
  "created_at" timestamp,
  "last_accessed_at" timestamp
);

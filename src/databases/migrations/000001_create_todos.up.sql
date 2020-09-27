CREATE TABLE IF NOT EXISTS "todos" (
  "id" uuid PRIMARY KEY DEFAULT  public.uuid_generate_v4(),
  "name" varchar(255) NOT NULL,
  "completed" BOOLEAN DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

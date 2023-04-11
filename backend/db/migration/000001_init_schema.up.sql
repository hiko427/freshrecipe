CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "cuisine" varchar NOT NULL,
  "num_people" bigint NOT NULL,
  "recipe_time" varchar NOT NULL,
  "unwanted_ingredients" text,
  "wanted_ingredients" text NOT NULL,
  "recipe" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "recipes" ("cuisine");

CREATE INDEX ON "recipes" ("recipe_time");

COMMENT ON COLUMN "recipes"."num_people" IS 'must be positive';
